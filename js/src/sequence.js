// @flow

import type {ChunkStore} from './chunk_store.js';
import {invariant, notNull} from './assert.js';
import {Type} from './type.js';
import {ValueBase} from './value.js';

export class Sequence<T> extends ValueBase {
  items: Array<T>;
  isMeta: boolean;

  constructor(type: Type, items: Array<T>) {
    super(type);

    this.items = items;
    this.isMeta = false;
  }

  getChildSequence(cs: ChunkStore, idx: number): Promise<?Sequence> { // eslint-disable-line no-unused-vars
    return Promise.resolve(null);
  }
}

export class SequenceCursor<T, S:Sequence> {
  cs: ChunkStore;
  parent: ?SequenceCursor;
  sequence: S;
  idx: number;

  constructor(cs: ChunkStore, parent: ?SequenceCursor, sequence: S, idx: number) {
    this.cs = cs;
    this.parent = parent;
    this.sequence = sequence;
    this.idx = idx;
  }

  get length(): number {
    return this.sequence.items.length;
  }

  getItem(idx: number): T {
    return this.sequence.items[idx];
  }

  async sync(): Promise<void> {
    invariant(this.parent);
    this.sequence = notNull(await this.parent.getChildSequence());
  }

  getChildSequence(): Promise<?S> {
    return this.sequence.getChildSequence(this.cs, this.idx);
  }

  getCurrent(): T {
    invariant(this.valid);
    return this.getItem(this.idx);
  }

  get valid(): boolean {
    return this.idx >= 0 && this.idx < this.length;
  }

  get indexInChunk(): number {
    return this.idx;
  }

  advance(): Promise<boolean> {
    return this._advanceMaybeAllowPastEnd(true);
  }

  advanceLocal(): boolean {
    if (this.idx < this.length - 1) {
      this.idx++;
      return true;
    }

    return false;
  }

  async _advanceMaybeAllowPastEnd(allowPastEnd: boolean): Promise<boolean> {
    if (this.idx < this.length - 1) {
      this.idx++;
      return true;
    }

    if (this.idx === this.length) {
      return false;
    }

    if (this.parent && (await this.parent._advanceMaybeAllowPastEnd(false))) {
      await this.sync();
      this.idx = 0;
      return true;
    }
    if (allowPastEnd) {
      this.idx++;
    }

    return false;
  }

  retreat(): Promise<boolean> {
    return this._retreatMaybeAllowBeforeStart(true);
  }

  async _retreatMaybeAllowBeforeStart(allowBeforeStart: boolean): Promise<boolean> {
    if (this.idx > 0) {
      this.idx--;
      return true;
    }
    if (this.idx === -1) {
      return false;
    }
    invariant(this.idx === 0);
    if (this.parent && this.parent._retreatMaybeAllowBeforeStart(false)) {
      await this.sync();
      this.idx = this.length - 1;
      return true;
    }

    if (allowBeforeStart) {
      this.idx--;
    }

    return false;
  }

  async iter(cb: (v: T, i: number) => boolean): Promise<void> {
    let idx = 0;
    while (this.valid) {
      if (cb(this.getItem(this.idx), idx++)) {
        return;
      }
      this.advanceLocal() || await this.advance();
    }
  }
}

// Translated from golang source (https://golang.org/src/sort/search.go?s=2249:2289#L49)
export function search(n: number, f: (i: number) => boolean): number {
  // Define f(-1) == false and f(n) == true.
  // Invariant: f(i-1) == false, f(j) == true.
  let i = 0;
  let j = n;
  while (i < j) {
    let h = i + (((j - i) / 2) | 0); // avoid overflow when computing h
    // i ≤ h < j
    if (!f(h)) {
      i = h + 1; // preserves f(i-1) == false
    } else {
      j = h; // preserves f(j) == true
    }
  }

  // i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
  return i;
}