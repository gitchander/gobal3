package geom

// Indent
// Thickness
// Indents
// Border, Frame
type Frame struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (a Frame) Add(b Frame) Frame {
	return Frame{
		Left:   a.Left + b.Left,
		Right:  a.Right + b.Right,
		Top:    a.Top + b.Top,
		Bottom: a.Bottom + b.Bottom,
	}
}

func (a Frame) Sub(b Frame) Frame {
	return Frame{
		Left:   a.Left - b.Left,
		Right:  a.Right - b.Right,
		Top:    a.Top - b.Top,
		Bottom: a.Bottom - b.Bottom,
	}
}

func (a Frame) MulScalar(scalar float64) Frame {
	return Frame{
		Left:   a.Left * scalar,
		Right:  a.Right * scalar,
		Top:    a.Top * scalar,
		Bottom: a.Bottom * scalar,
	}
}

func (a Frame) DivScalar(scalar float64) Frame {
	return Frame{
		Left:   a.Left / scalar,
		Right:  a.Right / scalar,
		Top:    a.Top / scalar,
		Bottom: a.Bottom / scalar,
	}
}

// Uniform
func MakeFrameUniform(width float64) Frame {
	return MakeFrame1(width)
}

func MakeFrame1(width float64) Frame {
	return Frame{
		Left:   width,
		Right:  width,
		Top:    width,
		Bottom: width,
	}
}

func MakeFrame2(leftRight, topBottom float64) Frame {
	return Frame{
		Left:   leftRight,
		Right:  leftRight,
		Top:    topBottom,
		Bottom: topBottom,
	}
}

func MakeFrame4(left, right, top, bottom float64) Frame {
	return Frame{
		Left:   left,
		Right:  right,
		Top:    top,
		Bottom: bottom,
	}
}
