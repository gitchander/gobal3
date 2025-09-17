package geom

import (
	"math"
)

type Point2f struct {
	X float64
	Y float64
}

var ZP2f Point2f // Zero Point2f

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

func (a Point2f) Add(b Point2f) Point2f {
	return Point2f{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Point2f) Sub(b Point2f) Point2f {
	return Point2f{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Point2f) MulScalar(scalar float64) Point2f {
	return Point2f{
		X: a.X * scalar,
		Y: a.Y * scalar,
	}
}

func (a Point2f) DivScalar(scalar float64) Point2f {
	return Point2f{
		X: a.X / scalar,
		Y: a.Y / scalar,
	}
}

func (a Point2f) Invert() Point2f {
	return Point2f{
		X: -a.X,
		Y: -a.Y,
	}
}

func (a Point2f) InvertAxisX() Point2f {
	return Point2f{
		X: -a.X,
		Y: a.Y,
	}
}

func (a Point2f) InvertAxisY() Point2f {
	return Point2f{
		X: a.X,
		Y: -a.Y,
	}
}

//------------------------------------------------------------------------------

// Point2f lerp
func PtLerp(p0, p1 Point2f, t float64) Point2f {
	return Point2f{
		X: lerp(p0.X, p1.X, t),
		Y: lerp(p0.Y, p1.Y, t),
	}
}

func Distance(a, b Point2f) float64 {
	var (
		dx = a.X - b.X
		dy = a.Y - b.Y
	)
	return math.Sqrt((dx * dx) + (dy * dy))
}
