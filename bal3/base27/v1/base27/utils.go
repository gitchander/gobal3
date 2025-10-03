package base27

func ceilDiv(a, b int) int {
	return (a + (b - 1)) / b
}

func hasDuplicate[T comparable](as []T) bool {
	m := make(map[T]struct{})
	for _, a := range as {
		if _, ok := m[a]; ok {
			return true
		}
		m[a] = struct{}{}
	}
	return false
}
