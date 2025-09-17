package draw3

import (
	"github.com/gitchander/gobal3/geom"
)

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
	positive []geom.Point2f
	negative []geom.Point2f
}
