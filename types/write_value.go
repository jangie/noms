package types

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/d"
	"github.com/attic-labs/noms/enc"
	"github.com/attic-labs/noms/ref"
)

type primitive interface {
	ToPrimitive() interface{}
}

func WriteValue(v Value, cs chunks.ChunkSink) ref.Ref {
	d.Chk.NotNil(cs)
	return writeValueInternal(v, cs)
}

func writeChildValueInternal(v Value, cs chunks.ChunkSink) ref.Ref {
	if cs == nil {
		return v.Ref()
	}

	return writeValueInternal(v, cs)
}

func writeValueInternal(v Value, cs chunks.ChunkSink) ref.Ref {
	e := toEncodeable(v, cs)
	w := chunks.NewChunkWriter()
	enc.Encode(w, e)
	c := w.Chunk()
	if cs != nil {
		cs.Put(c)
	}
	return c.Ref()
}

func toEncodeable(v Value, cs chunks.ChunkSink) interface{} {
	switch v := v.(type) {
	case blobLeaf:
		return v.Reader()
	case compoundBlob:
		return encCompoundBlobFromCompoundBlob(v, cs)
	case List:
		processListChildren(v, cs)
	case Map:
		processMapChildren(v, cs)
	case Package:
		processPackageChildren(v, cs)
	case Set:
		processSetChildren(v, cs)
	case TypeRef:
		processTypeRefChildren(v, cs)
	}
	return encNomsValue(v, cs)
}

func encCompoundBlobFromCompoundBlob(cb compoundBlob, cs chunks.ChunkSink) interface{} {
	refs := make([]ref.Ref, len(cb.futures))
	for idx, f := range cb.futures {
		i := processChild(f, cs)
		// All children of compoundBlob must be Blobs, which get encoded and reffed by processChild.
		refs[idx] = i.(ref.Ref)
	}
	return enc.CompoundBlob{Offsets: cb.offsets, Blobs: refs}
}

func processListChildren(l List, cs chunks.ChunkSink) {
	l.IterAll(func(v Value, i uint64) {
		writeChildValueInternal(v, cs)
	})
}

func processMapChildren(m Map, cs chunks.ChunkSink) {
	for _, r := range m.m {
		processChild(r.key, cs)
		processChild(r.value, cs)
	}
}

func processSetChildren(s Set, cs chunks.ChunkSink) {
	for _, f := range s.m {
		processChild(f, cs)
	}
}

func processTypeRefChildren(t TypeRef, cs chunks.ChunkSink) {
	if t.HasPackageRef() {
		pkgRef := t.PackageRef()
		p := LookupPackage(pkgRef)
		if p != nil {
			writeChildValueInternal(*p, cs)
		}
	}
	if desc, ok := t.Desc.(CompoundDesc); ok {
		for _, t := range desc.ElemTypes {
			writeChildValueInternal(t, cs)
		}
	}
}

func processPackageChildren(p Package, cs chunks.ChunkSink) {
	for _, t := range p.types {
		writeChildValueInternal(t, cs)
	}
	for _, r := range p.dependencies {
		p := LookupPackage(r)
		if p != nil {
			writeChildValueInternal(*p, cs)
		}
	}
}

func processChild(f Future, cs chunks.ChunkSink) interface{} {
	if v, ok := f.(*unresolvedFuture); ok {
		return v.Ref()
	}

	v := f.Val()
	d.Exp.NotNil(v)
	switch v := v.(type) {
	// Blobs, lists, maps, and sets are always out-of-line
	case Blob, List, Map, Set:
		return writeChildValueInternal(v, cs)
	default:
		// Other types are always inline.
		return toEncodeable(v, cs)
	}
}
