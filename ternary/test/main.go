package main

import (
	"fmt"

	"github.com/gitchander/gobal3/ternary"
)

func main() {
	testLogicTable()
}

var (
	neg  = ternary.Neg
	min  = ternary.Min
	max  = ternary.Max
	amin = ternary.Amin
	amax = ternary.Amax
	xor  = ternary.Xor
	nxor = ternary.NotXor
)

// Decoders:

func TritIs(a, v int) int {
	if a == v {
		return 1
	}
	return -1
}

func TritIsNot(a, v int) int {
	if a != v {
		return 1
	}
	return -1
}

func tritDec(a int) int {
	switch a {
	case -1:
		return 1
	case 0:
		return -1
	case 1:
		return 0
	default:
		return 0
	}
}

func testLogicTable() {

	// +---+---+---+---+
	// |   | T | 0 | 1 |
	// +---+---+---+---+
	// | T | 1 | 0 | T |
	// +---+---+---+---+
	// | 0 | 0 | 0 | 0 |
	// +---+---+---+---+
	// | 1 | T | 0 | 1 |
	// +---+---+---+---+

	f := func(a, b int) int {
		return min(xor(a, b), 1)
	}

	s := ternary.PrintableLogicTable("// ", f)
	fmt.Println(s)
}
