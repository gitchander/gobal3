package draw3

type Point2f struct {
	X, Y float64
}

var ZP Point2f // zero point

func MakePoint2f(x, y float64) Point2f {
	return Point2f{
		X: x,
		Y: y,
	}
}

func Pt2f(x, y float64) Point2f {
	return Point2f{
		X: x,
		Y: y,
	}
}

type Size = Point2f

func MakeSize(x, y float64) Size {
	return Size{
		X: x,
		Y: y,
	}
}
