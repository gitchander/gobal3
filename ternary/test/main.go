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
)

// Decoders:

func tritIs(a, v int) int {
	if a == v {
		return 1
	}
	return -1
}

func tritIsNot(a, v int) int {
	if a != v {
		return 1
	}
	return -1
}

func testLogicTable() {

	f := func(a, b int) int {
		return max(min(0, max(a, b)), min(a, b))
		//return max(min(a, b), max(min(neg(tritIs(a, -1)), 0), min(neg(tritIs(b, -1)), 0)))
	}

	s := ternary.PrintableLogicTable("// ", f)
	fmt.Println(s)
}
