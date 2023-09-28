package cow

func GroupBy[T any, K comparable](in Slice[T], groupOf func(x T) K) map[K]Slice[T] {
	out := map[K]Slice[T]{}
	for i := range in {
		group := groupOf(in[i])
		out[group] = append(out[group], in[i])
	}
	return out
}

func Map[T any, K any](in Slice[T], f func(T) K) Slice[K] {
	out := make(Slice[K], len(in))
	for i := range in {
		out[i] = f(in[i])
	}
	return out
}

func Reduce[TIN any, TOUT any](in Slice[TIN], init TOUT, f func(init TOUT, new TIN) TOUT) TOUT {
	for i := range in {
		init = f(init, in[i])
	}
	return init
}
