package draw3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer6 struct{}

var _ DigitDrawer = DigitDrawer6{}

func (DigitDrawer6) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

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

	var dx, dy float64
	dx, dy = 20, 20

	var (
		drawT = func(x, y float64) {
			c.MoveTo(x, y)
			c.LineTo(x+dx, y+dy)
			c.CubicTo(x-dx, y-dy, x-3*dx, y+dy, x-dx, y+3*dy)
			c.LineTo(x, y+4*dy)
		}

		draw1 = func(x, y float64) {
			c.MoveTo(x, y)
			c.CubicTo(x+2*dx, y+2*dy, x, y+4*dy, x-2*dx, y+2*dy)
			c.LineTo(x, y+4*dy)
		}

		draw0 = func(x, y float64) {
			drawT(x, y)
			draw1(x, y)
		}
	)

	digits := []int{digit}

	var (
		x0 = 50.0 + 0.5*dx
		y0 = 100.0 - 1.5*dy
	)
	c.MoveTo(x0-2*dx, y0-2*dy)
	c.LineTo(x0, y0)
	for _, digit := range digits {
		switch digit {
		case -1:
			drawT(x0, y0)
		case 0:
			draw0(x0, y0)
		case 1:
			draw1(x0, y0)
		}
		y0 += 4 * dy
	}
	c.LineTo(x0+1*dx, y0+1*dy)

	c.Stroke()
}
