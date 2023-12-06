package utils

func PtrToSlice[T comparable](s *[]*T) []T {
	//var actualSlice []T =
	size := len(*s)
	actualSlice := make([]T, size)
	for i := 0; i < size; i++ {
		actualSlice[i] = *(*s)[i]
	}

	//for _, ptr := range *s {
	//	actualSlice = append(actualSlice, *ptr)
	//}

	return actualSlice
}
