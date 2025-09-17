package geom

type Bounds struct {
	Min, Max Point2f
}

func MakeBounds(x1, y1, x2, y2 float64) Bounds {
	return Bounds{
		Min: MakePoint2f(x1, y1),
		Max: MakePoint2f(x2, y2),
	}
}

func Point2fToBounds(p Point2f) Bounds {
	return Bounds{
		Min: p,
		Max: p,
	}
}

func (b Bounds) Empty() bool {
	return !(b.notEmpty())
}

func (b Bounds) notEmpty() bool {
	return (b.Min.X < b.Max.X) && (b.Min.Y < b.Max.Y)
}

func (b Bounds) Dx() float64 {
	return b.Max.X - b.Min.X
}

func (b Bounds) Dy() float64 {
	return b.Max.Y - b.Min.Y
}

func (b Bounds) Size() Point2f {
	return Point2f{
		X: b.Max.X - b.Min.X,
		Y: b.Max.Y - b.Min.Y,
	}
}

func (b Bounds) Add(p Point2f) Bounds {
	return Bounds{
		Min: b.Min.Add(p),
		Max: b.Max.Add(p),
	}
}

func (b Bounds) Sub(p Point2f) Bounds {
	return Bounds{
		Min: b.Min.Sub(p),
		Max: b.Max.Sub(p),
	}
}

func (b Bounds) Center() Point2f {
	return (b.Min.Add(b.Max)).DivScalar(2)
}

func (b Bounds) Shrink(f Frame) Bounds {
	return Bounds{
		Min: Point2f{
			X: b.Min.X + f.Left,
			Y: b.Min.Y + f.Top,
		},
		Max: Point2f{
			X: b.Max.X - f.Right,
			Y: b.Max.Y - f.Bottom,
		},
	}
}

func (b Bounds) Grow(f Frame) Bounds {
	return Bounds{
		Min: Point2f{
			X: b.Min.X - f.Left,
			Y: b.Min.Y - f.Top,
		},
		Max: Point2f{
			X: b.Max.X + f.Right,
			Y: b.Max.Y + f.Bottom,
		},
	}
}

func (b Bounds) Vmin() float64 {
	return minFloat64(b.Dx(), b.Dy()) / 100.0
}

func (b Bounds) Vmax() float64 {
	return maxFloat64(b.Dx(), b.Dy()) / 100.0
}

//------------------------------------------------------------------------------

// aspectRatio = dy / dx
func BoundsAspect(b Bounds, aspectRatio float64) Bounds {

	var (
		dx = b.Dx()
		dy = b.Dy()
	)

	var (
		dx1 = dy / aspectRatio
		dy1 = dx * aspectRatio
	)

	if true {
		if dx1 < dx {
			dx = dx1
		} else {
			dy = dy1
		}
	} else { // or
		if dy1 < dy {
			dy = dy1
		} else {
			dx = dx1
		}
	}

	// dx = min(dx, dx1)
	// dy = min(dy, dy1)

	center := b.Center()
	f := MakeFrame2(dx, dy).DivScalar(2)
	b1 := Point2fToBounds(center).Grow(f)

	return b1
}
