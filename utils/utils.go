package utils

func Map[T, V comparable](v []T, f func(T) V) []V {
	vsm := make([]V, len(v))
	for i, v := range v {
		vsm[i] = f(v)
	}
	return vsm
}

type Stack []string

func (st *Stack) Push(s string) {
	*st = append(*st, s)
}

func (st *Stack) Pop() string {
	i := len(*st) - 1
	v := (*st)[i]
	*st = (*st)[:i]
	return v
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
