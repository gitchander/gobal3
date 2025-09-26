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

func matrixSize[T any](aa [][]T) geom.Point2i {
	var (
		x = 0
		y = len(aa)
	)
	for _, a := range aa {
		x = max(x, len(a))
	}
	return geom.MakePoint2i(x, y)
}
