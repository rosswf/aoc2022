package utils

func Map[T, V comparable](v []T, f func(T) V) []V {
	vsm := make([]V, len(v))
	for i, v := range v {
		vsm[i] = f(v)
	}
	return vsm
}
