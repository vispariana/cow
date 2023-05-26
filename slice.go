package slice

import (
	"golang.org/x/exp/slices"
)

type Slice[T comparable] []T

func New[T comparable](in ...T) Slice[T] {
	return in
}

func (x Slice[T]) Any(f func(T) bool) bool {
	return slices.ContainsFunc(x, f)
}

func (x Slice[T]) All(f func(T) bool) bool {
	return !slices.ContainsFunc(x, func(t T) bool {
		return !f(t)
	})
}

func (x Slice[T]) Contains(t T) bool {
	return slices.Contains(x, t)
}

func (x Slice[T]) Filter(f func(T) bool) (res Slice[T]) {
	for i := range x {
		if f(x[i]) {
			res = append(res, x[i])
		}
	}
	return
}

func (x Slice[T]) Reduce(init T, f func(init T, new T) T) T {
	for i := range x {
		init = f(init, x[i])
	}
	return init
}

func (x Slice[T]) SortFunc(cmp func(x, y T) bool) Slice[T] {
	dst := make(Slice[T], len(x))
	copy(dst, x)
	slices.SortFunc(dst, cmp)
	return dst
}

func (x Slice[T]) Unique() Slice[T] {
	dst := make(Slice[T], len(x))
	copy(dst, x)
	return slices.Compact(dst)
}

func (x Slice[T]) UniqueFunc(cmp func(x, y T) bool) Slice[T] {
	dst := make(Slice[T], len(x))
	copy(dst, x)
	return slices.CompactFunc(dst, cmp)
}
