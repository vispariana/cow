package cow

import (
	"fmt"
	"strings"
)

type Set[K comparable] map[K]bool

func NewSet[K comparable](in ...K) Set[K] {
	out := Set[K]{}
	for i := range in {
		out.Insert(in[i])
	}
	return out
}

func (s Set[K]) Contains(k K) bool {
	return s[k]
}

func (s Set[K]) Insert(k K) bool {
	if s[k] {
		return false
	}
	s[k] = true
	return true
}

func (s Set[K]) Union(other Set[K]) Set[K] {
	out := Set[K]{}
	for k := range s {
		out[k] = true
	}
	for k := range other {
		out[k] = true
	}
	return out
}

func (s Set[K]) Intersect(other Set[K]) Set[K] {
	out := Set[K]{}
	for k := range s {
		if other[k] {
			out[k] = true
		}
	}
	return out
}

func (s Set[K]) ToSlice() (out Slice[K]) {
	for k := range s {
		out = append(out, k)
	}
	return out
}

func (x Set[K]) String() string {
	b := strings.Builder{}
	b.WriteString("Set{")
	first := true
	for k := range x {
		if !first {
			b.WriteString(", ")
		} else {
			first = false
		}
		b.WriteString(fmt.Sprintf("%+v", k))
	}
	b.WriteRune('}')
	return b.String()
}

func (x Set[T]) All(f func(T) bool) bool {
	for i := range x {
		if !f(i) {
			return false
		}
	}
	return true
}

func (x Set[T]) Any(f func(T) bool) bool {
	for i := range x {
		if f(i) {
			return true
		}
	}
	return false
}

func (x Set[T]) Filter(f func(T) bool) Set[T] {
	out := Set[T]{}
	for i := range x {
		if f(i) {
			out.Insert(i)
		}
	}
	return out
}

func (x Set[T]) CountWhere(f func(T) bool) (out int) {
	for i := range x {
		if f(i) {
			out++
		}
	}
	return
}
