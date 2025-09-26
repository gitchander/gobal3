package draw3

// import (
// 	"github.com/fogleman/gg"

// 	"github.com/gitchander/gobal3/geom"
// )

// var Pt2f = geom.Pt2f

// var (
// 	nodesV1 = []sdNode{
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 0), Pt2f(0, 1)},
// 			negative: []geom.Point2f{Pt2f(1, 4), Pt2f(2, 3)},
// 		},
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 2), Pt2f(0, 1)},
// 			negative: []geom.Point2f{Pt2f(1, 2), Pt2f(2, 3)},
// 		},
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 2), Pt2f(0, 3)},
// 			negative: []geom.Point2f{Pt2f(1, 2), Pt2f(2, 1)},
// 		},
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 4), Pt2f(0, 3)},
// 			negative: []geom.Point2f{Pt2f(1, 0), Pt2f(2, 1)},
// 		},
// 	}

// 	nodesV2 = []sdNode{
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 0), Pt2f(2, 1)},
// 			negative: []geom.Point2f{Pt2f(1, 0), Pt2f(0, 1)},
// 		},
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 2), Pt2f(2, 1)},
// 			negative: []geom.Point2f{Pt2f(1, 2), Pt2f(0, 1)},
// 		},
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 2), Pt2f(2, 3)},
// 			negative: []geom.Point2f{Pt2f(1, 2), Pt2f(0, 3)},
// 		},
// 		{
// 			positive: []geom.Point2f{Pt2f(1, 4), Pt2f(2, 3)},
// 			negative: []geom.Point2f{Pt2f(1, 4), Pt2f(0, 3)},
// 		},
// 	}
// )

// type Size = geom.Point2f

// type TritsDrawer interface {
// 	SetDigitSize(digitSize Size)
// 	GetDigitSize() Size

// 	// n - number of trits
// 	Bounds(n int) Size

// 	Draw(c *gg.Context, p geom.Point2f, digits []int)
// }

// type TritsDrawer1 struct {
// 	digitSize Size
// }

// var _ TritsDrawer = &TritsDrawer1{}

// func (p *TritsDrawer1) SetDigitSize(digitSize Size) {
// 	p.digitSize = digitSize
// }

// func (p *TritsDrawer1) GetDigitSize() Size {
// 	return p.digitSize
// }

// func (p *TritsDrawer1) Bounds(n int) Size {
// 	return Size{
// 		X: p.digitSize.X * float64(n),
// 		Y: p.digitSize.Y,
// 	}
// }

// func (p *TritsDrawer1) drawDigit(c *gg.Context, pos geom.Point2f, digit int) {

// 	c.Push()
// 	defer c.Pop()

// 	var (
// 		size = geom.MakePoint2i(2, 4)
// 		//size = geom.MakePoint2i(4, 8)
// 	)

// 	w := minFloat64(p.digitSize.X/float64(size.X), p.digitSize.Y/float64(size.Y))
// 	w *= 0.7

// 	lineWidth := w * 0.3

// 	var (
// 		dX2 = (p.digitSize.X - w*float64(size.X)) / 2
// 		dY2 = (p.digitSize.Y - w*float64(size.Y)) / 2
// 	)

// 	c.Translate(pos.X+dX2, pos.Y+dY2)

// 	c.Scale(w, w)

// 	DrawGreedGG(c, size.X, size.Y, 1)

// 	subDigits := calcSubDigits(digit)

// 	var (
// 		nodes = nodesV1
// 		//nodes = nodesV2
// 	)
// 	n := minInt(len(nodes), len(subDigits))

// 	c.MoveTo(1, 0)
// 	c.LineTo(1, 4)

// 	for i := 0; i < n; i++ {
// 		var (
// 			node     = nodes[i]
// 			subDigit = subDigits[i]
// 		)
// 		switch subDigit {
// 		case -1:
// 			drawPolyline(c, node.negative)
// 		case +1:
// 			drawPolyline(c, node.positive)
// 		}
// 	}

// 	c.SetLineWidth(lineWidth)
// 	c.Stroke()
// }

// func (p *TritsDrawer1) Draw(c *gg.Context, pos geom.Point2f, digits []int) {
// 	for _, digit := range digits {
// 		p.drawDigit(c, pos, digit)
// 		pos.X += p.digitSize.X
// 	}
// }

// func drawPolyline(c *gg.Context, ps []geom.Point2f) {
// 	n := len(ps)
// 	if n > 0 {
// 		p := ps[0]
// 		c.MoveTo(p.X, p.Y)
// 	}
// 	for i := 1; i < n; i++ {
// 		p := ps[i]
// 		c.LineTo(p.X, p.Y)
// 	}
// }
