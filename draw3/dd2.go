package draw3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer2 struct{}

var _ DigitDrawer = DigitDrawer2{}

func (DigitDrawer2) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	lw := 10.0
	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)

	c.MoveTo(10, 60)
	c.LineTo(50, 20)
	c.LineTo(50, 180)
	c.LineTo(90, 140)

	d := digit

	// Positive
	if d == +1 {
		c.MoveTo(10, 60)
		c.LineTo(50, 100)
	}

	// Negative
	if d == -1 {
		c.MoveTo(50, 100)
		c.LineTo(90, 140)
	}

	c.Stroke()
}
