package iter

func Map[T, V any](ts []T, fn func(T) V) []V {
	res := make([]V, len(ts))

	for i, t := range ts {
		res[i] = fn(t)
	}

	return res
}
