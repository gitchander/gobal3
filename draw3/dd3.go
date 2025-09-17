package draw3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer3 struct{}

func (DigitDrawer3) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	lw := 10.0
	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)
	c.SetRGB(0, 0, 0)

	c.MoveTo(10, 40)
	c.QuadraticTo(50, 10, 50, 50)
	c.LineTo(50, 180)

	d := digit

	// Negative or Positive
	if (d == -1) || (d == +1) {
		c.MoveTo(30, 100)
		c.LineTo(70, 100)
	}

	c.Stroke()

	// Negative
	if d == -1 {
		c.DrawCircle(28, 54, 8)
		c.Fill()
	}
}
