package draw3

import (
	"github.com/fogleman/gg"
)

var DrawGreedEnable = false

func DrawGreedGG(c *gg.Context, nx, ny int, greedWidth float64) {

	if !DrawGreedEnable {
		return
	}

	c.Push()
	defer c.Pop()

	for i := 0; i <= ny; i++ {
		y := float64(i)
		c.MoveTo(0, y)
		c.LineTo(float64(nx), y)
	}

	for i := 0; i <= nx; i++ {
		x := float64(i)
		c.MoveTo(x, 0)
		c.LineTo(x, float64(ny))
	}

	c.SetLineWidth(greedWidth)
	c.SetRGB(0, 0, 1)
	c.Stroke()
}
