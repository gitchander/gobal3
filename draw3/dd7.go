package draw3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer7 struct{}

var _ DigitDrawer = DigitDrawer7{}

func (DigitDrawer7) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	lw := 10.0
	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)

	const (
		dx = 16
		dy = 16
	)

	var (
		dx1 = dx * 3.5
		dy1 = dy * 3.5
	)

	var (
		drawT = func(x, y float64) {
			c.MoveTo(x, y)
			c.LineTo(x+2*dx, y+2*dy)
			c.CubicTo((x+2*dx)-dx1, (y+2*dy)-dy1, x-dx1, (y+4*dy)-dy1, x, y+4*dy)
		}

		draw1 = func(x, y float64) {
			c.MoveTo(x, y)
			c.CubicTo(x+dx1, y+dy1, (x-2*dx)+dx1, (y+2*dy)+dy1, x-2*dx, y+2*dy)
			c.LineTo(x, y+4*dy)
		}

		draw0 = func(x, y float64) {
			drawT(x, y)
			draw1(x, y)
		}
	)

	digits := []int{digit}

	var (
		x0 = 50.0
		y0 = 100.0 - 2*dy
	)

	var (
		drawTail = true

		// tx = 1.5 * dx
		// ty = 1.5 * dy

		tx = 2.0 * dx
		ty = 2.0 * dy
	)
	if drawTail {
		c.MoveTo(x0-tx, y0-ty)
		c.LineTo(x0, y0)
	}
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
	if drawTail {
		c.MoveTo(x0, y0)
		c.LineTo(x0+tx, y0+ty)
	}

	c.Stroke()
}
