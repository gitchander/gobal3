package draw3

import (
	"github.com/gitchander/gobal3/geom"
	"github.com/gitchander/gobal3/utils/random"
)

type Coord = geom.Coord

type Digits2D = geom.Matrix2D[int]

func NewDigits2D(size Coord) *Digits2D {
	return geom.NewMatrix2D[int](size)
}

func RandDigits2D(r *random.Rand, nX, nY int) *Digits2D {
	var (
		size = geom.MakeCoord(nX, nY)
		m    = NewDigits2D(size)
	)
	m.Modify(
		func(coord Coord, value int) (int, bool) {
			return RandTrit(r), true
		})
	return m
}
