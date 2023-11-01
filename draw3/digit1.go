package draw3

import (
	"image"

	"github.com/fogleman/gg"
)

type TritsDrawer interface {
	SetDigitSize(digitSize Size)
	GetDigitSize() Size

	// n - number of trits
	Bounds(n int) Size

	Draw(c *gg.Context, p Point2f, digits []int)
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

func (p *TritsDrawer1) drawDigit(c *gg.Context, pos Point2f, digit int) {

	c.Push()
	defer c.Pop()

	var (
		size = image.Pt(2, 4)
		//size = image.Pt(4, 8)
	)

	w := minFloat64(p.digitSize.X/float64(size.X), p.digitSize.Y/float64(size.Y))
	w *= 0.8

	lineWidth := w * 0.3

	var (
		dX2 = (p.digitSize.X - w*float64(size.X)) / 2
		dY2 = (p.digitSize.Y - w*float64(size.Y)) / 2
	)

	c.Translate(pos.X+dX2, pos.Y+dY2)

	c.Scale(w, w)

	DrawGreedGG(c, size.X, size.Y, 1)

	nodes := []sdNode{
		{
			positive: []Point2f{Pt2f(1, 0), Pt2f(0, 1)},
			negative: []Point2f{Pt2f(1, 4), Pt2f(2, 3)},
		},
		{
			positive: []Point2f{Pt2f(1, 2), Pt2f(0, 1)},
			negative: []Point2f{Pt2f(1, 2), Pt2f(2, 3)},
		},
		{
			positive: []Point2f{Pt2f(1, 2), Pt2f(0, 3)},
			negative: []Point2f{Pt2f(1, 2), Pt2f(2, 1)},
		},
		{
			positive: []Point2f{Pt2f(1, 4), Pt2f(0, 3)},
			negative: []Point2f{Pt2f(1, 0), Pt2f(2, 1)},
		},
	}

	// nodes := []sdNode{
	// 	{
	// 		positive: []Point2f{Pt2f(1, 0), Pt2f(2, 1)},
	// 		negative: []Point2f{Pt2f(1, 0), Pt2f(0, 1)},
	// 	},
	// 	{
	// 		positive: []Point2f{Pt2f(1, 2), Pt2f(2, 1)},
	// 		negative: []Point2f{Pt2f(1, 2), Pt2f(0, 1)},
	// 	},
	// 	{
	// 		positive: []Point2f{Pt2f(1, 2), Pt2f(2, 3)},
	// 		negative: []Point2f{Pt2f(1, 2), Pt2f(0, 3)},
	// 	},
	// 	{
	// 		positive: []Point2f{Pt2f(1, 4), Pt2f(2, 3)},
	// 		negative: []Point2f{Pt2f(1, 4), Pt2f(0, 3)},
	// 	},
	// }

	subDigits := calcSubDigits(digit)

	n := minInt(len(nodes), len(subDigits))

	c.MoveTo(1, 0)
	c.LineTo(1, 4)

	for i := 0; i < n; i++ {
		var (
			node     = nodes[i]
			subDigit = subDigits[i]
		)
		switch subDigit {
		case -1:
			drawPolyline(c, node.negative)
		case +1:
			drawPolyline(c, node.positive)
		}
	}

	c.SetLineWidth(lineWidth)
	c.Stroke()
}

func (p *TritsDrawer1) Draw(c *gg.Context, pos Point2f, digits []int) {
	for _, digit := range digits {
		p.drawDigit(c, pos, digit)
		pos.X += p.digitSize.X
	}
}

func drawPolyline(c *gg.Context, ps []Point2f) {
	n := len(ps)
	if n > 0 {
		p := ps[0]
		c.MoveTo(p.X, p.Y)
	}
	for i := 1; i < n; i++ {
		p := ps[i]
		c.LineTo(p.X, p.Y)
	}
}
