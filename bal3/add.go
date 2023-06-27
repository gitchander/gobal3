package bal3

import "fmt"

// split sum of 4 trits

// min value:
// dec:  (-1)+(-1)+(-1)+(-1) = -4
// bal3: T+T+T = TT

// max value:
// dec:  1+1+1+1 = 4
// bal3: 1+1+1+1 = 11

func tritsSumSplit(v int) (hi, lo int) {
	switch v {
	case -4:
		hi, lo = tv_T, tv_T // TT
	case -3:
		hi, lo = tv_T, tv_0 // T0
	case -2:
		hi, lo = tv_T, tv_1 // T1
	case -1:
		hi, lo = tv_0, tv_T // 0T
	case 0:
		hi, lo = tv_0, tv_0 // 00
	case 1:
		hi, lo = tv_0, tv_1 // 01
	case 2:
		hi, lo = tv_1, tv_T // 1T
	case 3:
		hi, lo = tv_1, tv_0 // 10
	case 4:
		hi, lo = tv_1, tv_1 // 11
	default:
		panic(fmt.Errorf("there is invalid value %d", v))
	}
	return hi, lo
}

//------------------------------------------------------------------------------

// Add table

// +----+----+----+----+
// |    |  T |  0 |  1 |
// +----+----+----+----+
// |  T | T1 |  T |  0 |
// +----+----+----+----+
// |  0 |  T |  0 |  1 |
// +----+----+----+----+
// |  1 |  0 |  1 | 1T |
// +----+----+----+----+

func halfAdd(a, b int) (s, c int) {

	hi, lo := tritsSumSplit(a + b)

	s = lo
	c = hi

	return
}

//------------------------------------------------------------------------------

// Sub table:

// +----+----+----+----+
// |  - |  T |  0 |  1 |
// +----+----+----+----+
// |  T |  0 |  T | T1 |
// +----+----+----+----+
// |  0 |  1 |  0 |  T |
// +----+----+----+----+
// |  1 | 1T |  1 |  0 |
// +----+----+----+----+

func halfSub(a, b int) (s, c int) {

	hi, lo := tritsSumSplit(a - b)

	s = lo
	c = hi

	return
}

//------------------------------------------------------------------------------

func tritsAdd_v1(a, b int, c0 int) (s, c1 int) {
	hi, lo := tritsSumSplit((a + b) + c0)
	s = lo
	c1 = hi
	return
}

func tritsSub_v1(a, b int, c0 int) (s, c1 int) {
	hi, lo := tritsSumSplit((a - b) + c0)
	s = lo
	c1 = hi
	return
}

//------------------------------------------------------------------------------

func tritsAdd_v2(a, b int, c0 int) (s, c1 int) {

	var (
		s1, x1 = halfAdd(a, b)
		s2, x2 = halfAdd(s1, c0)
		s3, x3 = halfAdd(x1, x2)
	)

	_ = x3

	s = s2
	c1 = s3

	return
}

func tritsSub_v2(a, b int, c0 int) (s, c1 int) {

	var (
		s1, x1 = halfSub(a, b)
		s2, x2 = halfAdd(s1, c0)
		s3, x3 = halfAdd(x1, x2)
	)

	_ = x3

	s = s2
	c1 = s3

	return
}

//------------------------------------------------------------------------------

// tritsAdd: (a + b) + carry
// tritsSub: (a - b) + carry
var (
	// tritsAdd_ = tritsAdd_v1
	// tritsSub_ = tritsSub_v1

	tritsAdd = tritsAdd_v2
	tritsSub = tritsSub_v2
)

//------------------------------------------------------------------------------

// http://homepage.divms.uiowa.edu/%7Ejones/ternary/arith.shtml

// Add table

// +----+----+----+----+
// |    |  T |  0 |  1 |
// +----+----+----+----+
// |  T | T1 |  T |  0 |
// +----+----+----+----+
// |  0 |  T |  0 |  1 |
// +----+----+----+----+
// |  1 |  0 |  1 | 1T |
// +----+----+----+----+

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

func sum(a, b int) int {
	return 0
}

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

func cons(a, b int) int {
	return 0
}

//------------------------------------------------------------------------------

// A balanced ternary half adder

// c0 - carryIn
// c1 - carryOut

func halfAdder(a, c0 int) (s, c1 int) {
	s = sum(a, c0)
	c1 = cons(a, c0)
	return
}

//------------------------------------------------------------------------------

// Balanced Full Adder

// c0 - carryIn
// c1 - carryOut

func fullAdder(a, b, c0 int) (s, c1 int) {

	var (
		s1, x1 = halfAdder(a, b)
		s2, x2 = halfAdder(s1, c0)
		s3, x3 = halfAdder(x1, x2)
	)

	_ = x3

	s = s2
	c1 = s3

	return
}
