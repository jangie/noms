// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// MapOfStringToSetOfRefOfRemotePhoto

type MapOfStringToSetOfRefOfRemotePhoto struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToSetOfRefOfRemotePhoto() MapOfStringToSetOfRefOfRemotePhoto {
	return MapOfStringToSetOfRefOfRemotePhoto{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToSetOfRefOfRemotePhotoDef map[string]SetOfRefOfRemotePhotoDef

func (def MapOfStringToSetOfRefOfRemotePhotoDef) New() MapOfStringToSetOfRefOfRemotePhoto {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New())
	}
	return MapOfStringToSetOfRefOfRemotePhoto{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Def() MapOfStringToSetOfRefOfRemotePhotoDef {
	def := make(map[string]SetOfRefOfRemotePhotoDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(SetOfRefOfRemotePhoto).Def()
		return false
	})
	return def
}

func MapOfStringToSetOfRefOfRemotePhotoFromVal(val types.Value) MapOfStringToSetOfRefOfRemotePhoto {
	// TODO: Do we still need FromVal?
	if val, ok := val.(MapOfStringToSetOfRefOfRemotePhoto); ok {
		return val
	}
	// TODO: Validate here
	return MapOfStringToSetOfRefOfRemotePhoto{val.(types.Map), &ref.Ref{}}
}

func (m MapOfStringToSetOfRefOfRemotePhoto) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToSetOfRefOfRemotePhoto); ok {
		return m.Ref() == other.Ref()
	}
	return false
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToSetOfRefOfRemotePhoto.
var __typeRefForMapOfStringToSetOfRefOfRemotePhoto types.TypeRef

func (m MapOfStringToSetOfRefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToSetOfRefOfRemotePhoto
}

func init() {
	__typeRefForMapOfStringToSetOfRefOfRemotePhoto = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef(types.SetKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-8829255f15d0403222188a23b8ced0afdbee2a97"), 1))))
	types.RegisterFromValFunction(__typeRefForMapOfStringToSetOfRefOfRemotePhoto, func(v types.Value) types.Value {
		return MapOfStringToSetOfRefOfRemotePhotoFromVal(v)
	})
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Get(p string) SetOfRefOfRemotePhoto {
	return m.m.Get(types.NewString(p)).(SetOfRefOfRemotePhoto)
}

func (m MapOfStringToSetOfRefOfRemotePhoto) MaybeGet(p string) (SetOfRefOfRemotePhoto, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewSetOfRefOfRemotePhoto(), false
	}
	return v.(SetOfRefOfRemotePhoto), ok
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Set(k string, v SetOfRefOfRemotePhoto) MapOfStringToSetOfRefOfRemotePhoto {
	return MapOfStringToSetOfRefOfRemotePhoto{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToSetOfRefOfRemotePhoto) Remove(p string) MapOfStringToSetOfRefOfRemotePhoto {
	return MapOfStringToSetOfRefOfRemotePhoto{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToSetOfRefOfRemotePhotoIterCallback func(k string, v SetOfRefOfRemotePhoto) (stop bool)

func (m MapOfStringToSetOfRefOfRemotePhoto) Iter(cb MapOfStringToSetOfRefOfRemotePhotoIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(SetOfRefOfRemotePhoto))
	})
}

type MapOfStringToSetOfRefOfRemotePhotoIterAllCallback func(k string, v SetOfRefOfRemotePhoto)

func (m MapOfStringToSetOfRefOfRemotePhoto) IterAll(cb MapOfStringToSetOfRefOfRemotePhotoIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(SetOfRefOfRemotePhoto))
	})
}

type MapOfStringToSetOfRefOfRemotePhotoFilterCallback func(k string, v SetOfRefOfRemotePhoto) (keep bool)

func (m MapOfStringToSetOfRefOfRemotePhoto) Filter(cb MapOfStringToSetOfRefOfRemotePhotoFilterCallback) MapOfStringToSetOfRefOfRemotePhoto {
	nm := NewMapOfStringToSetOfRefOfRemotePhoto()
	m.IterAll(func(k string, v SetOfRefOfRemotePhoto) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// SetOfRefOfRemotePhoto

type SetOfRefOfRemotePhoto struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfRefOfRemotePhoto() SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{types.NewSet(), &ref.Ref{}}
}

type SetOfRefOfRemotePhotoDef map[ref.Ref]bool

func (def SetOfRefOfRemotePhotoDef) New() SetOfRefOfRemotePhoto {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = NewRefOfRemotePhoto(d)
		i++
	}
	return SetOfRefOfRemotePhoto{types.NewSet(l...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Def() SetOfRefOfRemotePhotoDef {
	def := make(map[ref.Ref]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(RefOfRemotePhoto).TargetRef()] = true
		return false
	})
	return def
}

func SetOfRefOfRemotePhotoFromVal(val types.Value) SetOfRefOfRemotePhoto {
	// TODO: Do we still need FromVal?
	if val, ok := val.(SetOfRefOfRemotePhoto); ok {
		return val
	}
	return SetOfRefOfRemotePhoto{val.(types.Set), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) InternalImplementation() types.Set {
	return s.s
}

func (s SetOfRefOfRemotePhoto) Equals(other types.Value) bool {
	if other, ok := other.(SetOfRefOfRemotePhoto); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s SetOfRefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfRefOfRemotePhoto) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfRefOfRemotePhoto.
var __typeRefForSetOfRefOfRemotePhoto types.TypeRef

func (m SetOfRefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForSetOfRefOfRemotePhoto
}

func init() {
	__typeRefForSetOfRefOfRemotePhoto = types.MakeCompoundTypeRef(types.SetKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-8829255f15d0403222188a23b8ced0afdbee2a97"), 1)))
	types.RegisterFromValFunction(__typeRefForSetOfRefOfRemotePhoto, func(v types.Value) types.Value {
		return SetOfRefOfRemotePhotoFromVal(v)
	})
}

func (s SetOfRefOfRemotePhoto) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRefOfRemotePhoto) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRefOfRemotePhoto) Has(p RefOfRemotePhoto) bool {
	return s.s.Has(p)
}

type SetOfRefOfRemotePhotoIterCallback func(p RefOfRemotePhoto) (stop bool)

func (s SetOfRefOfRemotePhoto) Iter(cb SetOfRefOfRemotePhotoIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoIterAllCallback func(p RefOfRemotePhoto)

func (s SetOfRefOfRemotePhoto) IterAll(cb SetOfRefOfRemotePhotoIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoFilterCallback func(p RefOfRemotePhoto) (keep bool)

func (s SetOfRefOfRemotePhoto) Filter(cb SetOfRefOfRemotePhotoFilterCallback) SetOfRefOfRemotePhoto {
	ns := NewSetOfRefOfRemotePhoto()
	s.IterAll(func(v RefOfRemotePhoto) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfRefOfRemotePhoto) Insert(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Remove(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Union(others ...SetOfRefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Subtract(others ...SetOfRefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Any() RefOfRemotePhoto {
	return s.s.Any().(RefOfRemotePhoto)
}

func (s SetOfRefOfRemotePhoto) fromStructSlice(p []SetOfRefOfRemotePhoto) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRefOfRemotePhoto) fromElemSlice(p []RefOfRemotePhoto) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// RefOfRemotePhoto

type RefOfRemotePhoto struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfRemotePhoto(target ref.Ref) RefOfRemotePhoto {
	return RefOfRemotePhoto{target, &ref.Ref{}}
}

func (r RefOfRemotePhoto) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfRemotePhoto) Equals(other types.Value) bool {
	if other, ok := other.(RefOfRemotePhoto); ok {
		return r.Ref() == other.Ref()
	}
	return false
}

func (r RefOfRemotePhoto) Chunks() []types.Future {
	return r.TypeRef().Chunks()
}

func RefOfRemotePhotoFromVal(val types.Value) RefOfRemotePhoto {
	// TODO: Do we still need FromVal?
	if val, ok := val.(RefOfRemotePhoto); ok {
		return val
	}
	return NewRefOfRemotePhoto(val.(types.Ref).TargetRef())
}

// A Noms Value that describes RefOfRemotePhoto.
var __typeRefForRefOfRemotePhoto types.TypeRef

func (m RefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForRefOfRemotePhoto
}

func init() {
	__typeRefForRefOfRemotePhoto = types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-8829255f15d0403222188a23b8ced0afdbee2a97"), 1))
	types.RegisterFromValFunction(__typeRefForRefOfRemotePhoto, func(v types.Value) types.Value {
		return RefOfRemotePhotoFromVal(v)
	})
}

func (r RefOfRemotePhoto) TargetValue(cs chunks.ChunkSource) RemotePhoto {
	return types.ReadValue(r.target, cs).(RemotePhoto)
}

func (r RefOfRemotePhoto) SetTargetValue(val RemotePhoto, cs chunks.ChunkSink) RefOfRemotePhoto {
	return NewRefOfRemotePhoto(types.WriteValue(val, cs))
}