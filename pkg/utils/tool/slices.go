package tool

func Find[T ~[]E, E any](data T, f func(item E, index int) bool) *E {
	for i, v := range data {
		if flag := f(v, i); flag {
			return &v
		}
	}
	return nil
}

func FindIndex[T any](data []T, f func(item T, index int) bool) int {
	for i, v := range data {
		if flag := f(v, i); flag {
			return i
		}
	}
	return -1
}
