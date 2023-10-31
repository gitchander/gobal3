package draw3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/bal3"
)

type TritsDrawer interface {
	SetDigitSize(digitSize Size)
	GetDigitSize() Size

	// n - number of trits
	Bounds(n int) Size

	Draw(*gg.Context, Point2f, []bal3.Trit)
}

type TritsDrawer1 struct {
	digitSize Size
}

var _ TritsDrawer = &TritsDrawer1{}

func (p *TritsDrawer1) SetDigitSize(digitSize Size) {
	p.digitSize = digitSize
}

func (p *TritsDrawer1) GetDigitSize() Size {
	return p.digitSize
}

func (p *TritsDrawer1) Bounds(n int) Size {
	return Size{
		X: p.digitSize.X * float64(n),
		Y: p.digitSize.Y,
	}
}

func (p *TritsDrawer1) drawDigit(c *gg.Context, pos Point2f, t bal3.Trit) {

	c.Push()

	w := minFloat64(p.digitSize.X, p.digitSize.Y/2) / 2
	w *= 0.8

	lineWidth := w * 0.3

	var (
		dX2 = p.digitSize.X / 2
		dY2 = p.digitSize.Y / 2
	)

	c.Translate(pos.X+dX2-w, pos.Y+dY2-w*2)

	c.Scale(w, w)

	if false {
		c.SetHexColor("f00")
		c.DrawRectangle(0, 0, 2, 4)
		c.Fill()
		c.SetHexColor("000")
	}

	switch t {
	case -1:
		c.MoveTo(1, 0)
		c.LineTo(1, 4)
		c.LineTo(2, 3)
	case 0:
		if true {
			c.MoveTo(1, 0)
			c.LineTo(1, 4)
		} else {
			c.MoveTo(0, 1)
			c.LineTo(1, 0)
			c.LineTo(1, 4)
			c.LineTo(2, 3)
		}
	case +1:
		c.MoveTo(0, 1)
		c.LineTo(1, 0)
		c.LineTo(1, 4)
	}

	c.SetLineWidth(lineWidth)
	c.Stroke()

	c.Pop()
}

func (p *TritsDrawer1) Draw(c *gg.Context, pos Point2f, ts []bal3.Trit) {
	for _, t := range ts {
		p.drawDigit(c, pos, t)
		pos.X += p.digitSize.X
	}
}
