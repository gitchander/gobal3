package draw3

import (
	"fmt"
	"strconv"

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

//------------------------------------------------------------------------------

func formatInt(a int) string {
	return strconv.Itoa(a)
}

func formatDigit(digit int) string {
	if digit == 0 {
		return formatInt(digit) // no sign
	}
	return fmt.Sprintf("%+d", digit) // with sign
}
