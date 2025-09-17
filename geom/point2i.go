package geom

type Point2i struct {
	X int
	Y int
}

func MakePoint2i(x, y int) Point2i {
	return Point2i{
		X: x,
		Y: y,
	}
}
