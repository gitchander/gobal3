package draw3

import (
	"fmt"

	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

type DigitDrawer interface {
	DrawDigit(c *gg.Context, b geom.Bounds, digit int)
}

func DrawDigits(c *gg.Context, dd DigitDrawer, digitSize geom.Point2f, digits []int) {
	var (
		y = 0

		y0 = float64(y+0) * digitSize.Y
		y1 = float64(y+1) * digitSize.Y
	)
	for x, digit := range digits {
		var (
			x0 = float64(x+0) * digitSize.X
			x1 = float64(x+1) * digitSize.X
		)
		b := geom.MakeBounds(x0, y0, x1, y1)
		dd.DrawDigit(c, b, digit)
	}
}

func DrawMatrix(c *gg.Context, matrixSize geom.Point2i,
	dd DigitDrawer, digitSize geom.Point2f, digits []int) {

	for y := 0; y < matrixSize.Y; y++ {
		var (
			y0 = float64(y+0) * digitSize.Y
			y1 = float64(y+1) * digitSize.Y
		)
		for x := 0; x < matrixSize.X; x++ {
			var (
				x0 = float64(x+0) * digitSize.X
				x1 = float64(x+1) * digitSize.X
			)
			b := geom.MakeBounds(x0, y0, x1, y1)
			if true {
				if ((x + y) % 2) == 0 {
					c.SetRGB(0.7, 0.9, 1.0)
				} else {
					c.SetRGB(1, 1, 1)
				}
				drawBounds(c, b)
				c.Fill()
			}
			if len(digits) > 0 {
				digit := digits[0]
				digits = digits[1:]

				c.SetRGB(0, 0, 0)
				c.DrawString(fmt.Sprintf("%d", digit), b.Min.X, b.Min.Y+c.FontHeight())

				c.SetRGB(0, 0, 0)
				dd.DrawDigit(c, b, digit)
			}
		}
	}
}

func drawBounds(c *gg.Context, b geom.Bounds) {
	c.DrawRectangle(b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
}
