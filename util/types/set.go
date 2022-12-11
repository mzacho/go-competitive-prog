package types

import "reflect"

type Set[T comparable] map[T]bool
type IntSet = Set[int]
type StringSet = Set[string]
type IntPairSet = Set[IntPair]

func NewSet[T comparable](xs ...T) Set[T] {
	a := make(Set[T])
	for _, x := range xs {
		a[x] = true
	}
	return a
}

func NewEmptySet[T comparable]() Set[T] {
	return make(Set[T])
}

func (as Set[T]) Contains(a T) bool {
	return as[a]
}

// Add element to set, return true if set was updated
func (as Set[T]) Add(a T) bool {
	p := as[a]
	as[a] = true
	return !p
}

// Remove element from set, returns true if set was updated
func (as Set[T]) Remove(a T) bool {
	p := as[a]
	delete(as, a)
	return p
}

func (as Set[T]) ToList() []T {
	keys := make([]T, 0, len(as))
	for k := range as {
		keys = append(keys, k)
	}
	return keys
}

func (as Set[T]) Union(bs Set[T]) Set[T] {
	cs := NewSet[T]()
	for a := range as {
		cs[a] = true
	}
	for b := range bs {
		cs[b] = true
	}
	return cs
}

func (as Set[T]) MutableUnion(bs Set[T]) {
	for b := range bs {
		as.Add(b)
	}
}

func (as Set[T]) Minus(bs Set[T]) Set[T] {
	cs := NewSet[T]()
	for a := range as {
		if !bs[a] {
			cs[a] = true
		}
	}
	return cs
}

func (as Set[T]) MutableMinus(bs Set[T]) {
	for b := range bs {
		delete(as, b)
	}
}

func (as Set[T]) Intersect(bs Set[T]) Set[T] {
	cs := NewSet[T]()
	for b := range bs {
		if as[b] {
			cs[b] = true
		}
	}
	return cs
}

func (as Set[T]) MutableIntersect(bs Set[T]) {
	for b := range bs {
		if !as[b] {
			delete(as, b)
		}
	}
}

func (as Set[T]) Equals(bs Set[T]) bool {
	return reflect.DeepEqual(as, bs)
}
