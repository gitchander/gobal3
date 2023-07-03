package bal3

import (
	. "github.com/gitchander/gobal3/ternary"
)

// http://homepage.divms.uiowa.edu/%7Ejones/ternary/arith.shtml

//------------------------------------------------------------------------------

// Add table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T |T1 | T | 0 |
// +---+---+---+---+
// | 0 | T | 0 | 1 |
// +---+---+---+---+
// | 1 | 0 | 1 |1T |
// +---+---+---+---+

//------------------------------------------------------------------------------

// Sub table:

// +---+---+---+---+
// | - | T | 0 | 1 |
// +---+---+---+---+
// | T | 0 | T |T1 |
// +---+---+---+---+
// | 0 | 1 | 0 | T |
// +---+---+---+---+
// | 1 |1T | 1 | 0 |
// +---+---+---+---+

//------------------------------------------------------------------------------

// +---+---+---+---+---+
// | A | B |A+B|con|sum|
// +---+---+---+---+---+
// | T | T |T1 | T | 1 |
// | T | 0 | T | 0 | T |
// | T | 1 | 0 | 0 | 0 |
// | 0 | T | T | 0 | T |
// | 0 | 0 | 0 | 0 | 0 |
// | 0 | 1 | 1 | 0 | 1 |
// | 1 | T | 0 | 0 | 0 |
// | 1 | 0 | 1 | 0 | 1 |
// | 1 | 1 |1T | 1 | T |
// +---+---+---+---+---+

//------------------------------------------------------------------------------

// Sum table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | T | 0 |
// +---+---+---+---+
// | 0 | T | 0 | 1 |
// +---+---+---+---+
// | 1 | 0 | 1 | T |
// +---+---+---+---+

var tableAddSum = mustParseTable(
	"1T0",
	"T01",
	"01T",
)

func addSumV1(a, b int) int {
	return tritByTable(tableAddSum, a, b)
}

func addSumV2(a, b int) int {
	_, t0 := splitTrits2(a + b)
	return t0
}

func addSumV3(a, b int) int {
	c := a + b
	if c < -1 {
		c += 3
	}
	if c > +1 {
		c -= 3
	}
	return c
}

func addSumV4(a, b int) int {

	// (a + b) = ((a = -1) ∧ (b - 1)) ∨ ((a = 0) ∧ (b)) ∨ ((a = +1) ∧ (b + 1))
	// where:
	// ∧ - min
	// ∨ - max
	// (b - 1) - dec(b)
	// (b + 1) - inc(b)

	var (
		v1 = Min(Is(a, -1), Dec(b))
		v2 = Min(Is(a, 0), b)
		v3 = Min(Is(a, +1), Inc(b))
	)

	return Max(v1, Max(v2, v3))
}

var (
	//addSum = addSumV1
	addSum = addSumV2
	//addSum = addSumV3
	//addSum = addSumV4
)

//------------------------------------------------------------------------------

// Consider table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | T | 0 | 0 |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | 0 | 0 | 1 |
// +---+---+---+---+

var tableAddCons = mustParseTable(
	"T00",
	"000",
	"001",
)

func addConsV1(a, b int) int {
	return tritByTable(tableAddCons, a, b)
}

func addConsV2(a, b int) int {
	t1, _ := splitTrits2(a + b)
	return t1
}

func addConsV3(a, b int) int {
	if v := -1; (a == v) && (b == v) {
		return v
	}
	if v := 1; (a == v) && (b == v) {
		return v
	}
	return 0
}

func addConsV4(a, b int) int {
	var (
		v1 = Min(a, b)
		v2 = Min(Neg(Is(a, -1)), 0)
		v3 = Min(Neg(Is(b, -1)), 0)
	)
	return Max(v1, Max(v2, v3))
}

func addConsV5(a, b int) int {
	return Max(Min(a, b), Min(0, Max(a, b)))
}

var (
	//addCons = addConsV1
	addCons = addConsV2
	//addCons = addConsV3
	//addCons = addConsV4
	//addCons = addConsV5
)

//------------------------------------------------------------------------------

// A balanced ternary half adder

// c - carryOut

func halfAdderV1(a, b int) (s, c int) {
	t1, t0 := splitTrits2(a + b)
	s = t0
	c = t1
	return
}

func halfAdderV2(a, b int) (s, c int) {
	s = addSum(a, b)
	c = addCons(a, b)
	return
}

var (
	halfAdder = halfAdderV1
	//halfAdder = halfAdderV2
)

//------------------------------------------------------------------------------

// Balanced Full Adder

// c0 - carryIn
// c1 - carryOut

func fullAdder(a, b int, c0 int) (s, c1 int) {
	var (
		s1, x1 = halfAdder(a, b)
		s2, x2 = halfAdder(s1, c0)
		s3     = addSum(x1, x2)
	)
	s = s2
	c1 = s3
	return
}

//------------------------------------------------------------------------------

func tritsAddV1(a, b int, c0 int) (s, c1 int) {
	return fullAdder(a, b, c0)
}

func tritsAddV2(a, b int, c0 int) (s, c1 int) {
	t1, t0 := splitTrits2(a + b + c0)
	s = t0
	c1 = t1
	return
}

var (
	//tritsAdd = tritsAddV1
	tritsAdd = tritsAddV2
)

//------------------------------------------------------------------------------

func tritsSubV1(a, b int, c0 int) (s, c1 int) {
	b = Neg(b)
	return tritsAdd(a, b, c0)
}

func tritsSubV2(a, b int, c0 int) (s, c1 int) {
	t1, t0 := splitTrits2(a - b + c0)
	s = t0
	c1 = t1
	return
}

var (
	//tritsSub = tritsSubV1
	tritsSub = tritsSubV2
)

//------------------------------------------------------------------------------
