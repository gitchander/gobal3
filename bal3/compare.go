package bal3

// func Compare(a, b string) int {
// 	if a == b {
// 		return 0
// 	}
// 	if a < b {
// 		return -1
// 	}
// 	return +1
// }

// a, b - trits
func tritsCompare(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
