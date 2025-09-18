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
	c.SetRGB(0, 0, 0)

	c.MoveTo(10, 50)
	c.QuadraticTo(60, 20, 60, 60)
	c.LineTo(60, 170)

	d := digit

	// Negative or Positive
	if (d == -1) || (d == +1) {
		c.MoveTo(60-14, 112)
		c.LineTo(60+14, 100)
	}

	c.Stroke()

	// Negative
	if d == -1 {
		c.DrawCircle(38, 59, 8)
		c.Fill()
	}
}
