package main

import (
	"fmt"

	"github.com/gitchander/gobal3/ternary"
)

func main() {

	var (
		neg = ternary.Neg
		min = ternary.Min
		max = ternary.Max
		xor = ternary.Xor
		inc = ternary.Inc
		dec = ternary.Dec
	)

	_, _, _, _, _, _ = neg, min, max, xor, inc, dec

	// +---+---+---+---+
	// |   | T | 0 | 1 |
	// +---+---+---+---+
	// | T | 1 | T | 0 |
	// +---+---+---+---+
	// | 0 | T | 0 | 1 |
	// +---+---+---+---+
	// | 1 | 0 | 1 | T |
	// +---+---+---+---+

	f := func(a, b int) int {
		return neg(dec(xor(a, b)))
	}

	s := ternary.PrintableBinaryTable("\t", f)
	fmt.Print(s)
}
