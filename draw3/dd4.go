package draw3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer4 struct{}

var _ DigitDrawer = DigitDrawer4{}

func (DigitDrawer4) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

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

	d := digit

	// Negative or Positive
	if (d == -1) || (d == +1) {
		const (
			dx = 15.0
		)
		c.MoveTo(x0-dx, 110)
		c.LineTo(x0+dx, 100)
	}

	c.Stroke()

	// Negative
	if d == -1 {
		c.DrawCircle(x0-23, 68, 8)
		c.Fill()
	}
}
