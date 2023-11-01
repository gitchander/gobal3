package main

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/bal3"
	"github.com/gitchander/gobal3/draw3"
	"github.com/gitchander/gobal3/utils/random"
)

func main() {

	var td draw3.TritsDrawer = &draw3.TritsDrawer1{}

	var (
		dh = 50.0
	)
	td.SetDigitSize(draw3.MakeSize(dh/2, dh))

	// r := random.NewRandNow()
	// digits := make([]int, 10)
	// for i := range digits {
	// 	digits[i] = randDigit(r)
	// }
	digits := makeIntsMinMax(-40, 41)

	b := td.Bounds(len(digits))

	var (
		w = int(math.Ceil(b.X))
		h = int(math.Ceil(b.Y))
	)

	c := gg.NewContext(w, h)
	c.SetColor(color.White)
	c.Clear()

	c.SetColor(color.Black)
	td.Draw(c, draw3.ZP, digits)

	c.SavePNG("result.png")
}

var allTrits = []bal3.Trit{-1, 0, 1}

func randDigit(r *random.Rand) int {
	const (
		min = -40
		max = +40
	)
	return random.RandIntMinMax(r, min, max+1)
}

func randTrit(r *random.Rand) bal3.Trit {
	return random.RandByCorpus(r, allTrits)
}

func makeIntsMinMax(min, max int) []int {
	n := max - min
	as := make([]int, n)
	for i := range as {
		as[i] = min + i
	}
	return as
}
