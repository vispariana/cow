package cow

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
