package helpers

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func Remove[T any](s []T, i int) []T {
	tempS := make([]T, len(s))
	copy(tempS, s)
	result := append(tempS[:i], tempS[i+1:]...)
	return result
}

func IndexOf(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
