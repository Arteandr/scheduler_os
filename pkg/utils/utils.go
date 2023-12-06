package utils

func PtrToSlice[T comparable](s *[]*T) []T {
	size := len(*s)
	actualSlice := make([]T, size)
	for i := 0; i < size; i++ {
		actualSlice[i] = *(*s)[i]
	}

	return actualSlice
}
