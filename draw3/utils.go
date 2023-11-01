package draw3

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// subDigit
type sdNode struct {
	positive []Point2f
	negative []Point2f
}
