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

	r := random.NewRandNow()

	//ts := []bal3.Trit{0, 1, -1, 0, 1, -1}
	ts := make([]bal3.Trit, 20)
	for i := range ts {
		ts[i] = randTrit(r)
	}

	b := td.Bounds(len(ts))

	var (
		w = int(math.Ceil(b.X))
		h = int(math.Ceil(b.Y))
	)

	c := gg.NewContext(w, h)
	c.SetColor(color.White)
	c.Clear()

	c.SetColor(color.Black)
	td.Draw(c, draw3.ZP, ts)

	c.SavePNG("result.png")
}

var allTrits = []bal3.Trit{-1, 0, 1}

func randTrit(r *random.Rand) bal3.Trit {
	return random.RandByCorpus(r, allTrits)
}
