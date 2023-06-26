package bal3

// Invert table:

// +---+---+
// | t |inv|
// +---+---+
// | T | 1 |
// +---+---+
// | 0 | 0 |
// +---+---+
// | 1 | T |
// +---+---+

func invertTrit(t int) int {
	switch t {
	case tv_T:
		return tv_1
	case tv_0:
		return tv_0
	case tv_1:
		return tv_T
	default:
		panic(errInvalidTrit(t))
	}
}
