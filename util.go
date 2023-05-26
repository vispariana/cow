package cow

func GroupBy[T comparable, K comparable](in Slice[T], groupOf func(x T) K) map[K]Slice[T] {
	out := map[K]Slice[T]{}
	for i := range in {
		group := groupOf(in[i])
		out[group] = append(out[group], in[i])
	}
	return out
}

func Map[T comparable, K comparable](in Slice[T], f func(T) K) Slice[K] {
	out := make(Slice[K], len(in))
	for i := range in {
		out[i] = f(in[i])
	}
	return out
}
