package draw3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer5 struct{}

var _ DigitDrawer = DigitDrawer5{}

func (DigitDrawer5) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	lw := 10.0
	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)

	x0 := 65.0

	c.MoveTo(x0-50, 60)
	c.QuadraticTo(x0, 20, x0, 70)
	c.LineTo(x0, 160)

	// Horizontal line
	{
		c.MoveTo(x0-30, 105)
		c.LineTo(x0+20, 105)
	}

	c.Stroke()

	d := digit

	var (
		y0 = 105.0

		dx = 20.0
		dy = 25.0
		// dy = 25.0
	)

	// Negative
	if d == -1 {
		c.DrawCircle(x0-dx, y0+dy, 8)
		c.Fill()
	}

	// Positive
	if d == +1 {
		c.DrawCircle(x0-dx, y0-dy, 8)
		c.Fill()
	}
}
