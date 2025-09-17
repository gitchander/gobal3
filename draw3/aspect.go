package draw3

import (
	"math"
)

const AspectRatio = 2.0 // dy / dx

func CalcSizeY(sizeX int) (sizeY int) {
	sizeY = int(math.Ceil(float64(sizeX) * AspectRatio))
	return sizeY
}
