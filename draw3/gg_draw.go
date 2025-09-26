package draw3

import (
	"fmt"

	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer interface {
	DrawDigit(c *gg.Context, b geom.Bounds, digit int)
}

func DrawDigits(c *gg.Context, dd DigitDrawer,
	digitSize, stride geom.Point2f, digits []int) {

	y := 0
	var (
		y0 = float64(y) * stride.Y
		y1 = y0 + digitSize.Y
	)
	for x, digit := range digits {
		var (
			x0 = float64(x) * stride.X
			x1 = x0 + digitSize.X
		)
		b := geom.MakeBounds(x0, y0, x1, y1)
		dd.DrawDigit(c, b, digit)
	}
}

//------------------------------------------------------------------------------

func DrawDigits2D(c *gg.Context, dd DigitDrawer,
	digitSize, stride geom.Point2f, d2d *Digits2D) {

	matrixSize := d2d.Size()

	for y := 0; y < matrixSize.Y; y++ {
		var (
			y0 = float64(y) * stride.Y
			y1 = y0 + digitSize.Y
		)
		for x := 0; x < matrixSize.X; x++ {
			var (
				x0 = float64(x) * stride.X
				x1 = x0 + digitSize.X
			)

			b := geom.MakeBounds(x0, y0, x1, y1)

			digit, _ := d2d.GetValueXY(x, y)

			if true {
				if ((x + y) % 2) == 0 {
					c.SetRGB(0.7, 0.9, 1.0)
				} else {
					c.SetRGB(1, 1, 1)
				}
				drawBounds(c, b)
				c.Fill()
			}

			// draw text
			if true {
				c.SetRGB(0, 0, 0)
				c.DrawString(fmt.Sprintf("%d", digit), b.Min.X, b.Min.Y+c.FontHeight())
			}

			c.SetRGB(0, 0, 0)
			dd.DrawDigit(c, b, digit)

		}
	}
}

func drawBounds(c *gg.Context, b geom.Bounds) {
	c.DrawRectangle(b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
}
