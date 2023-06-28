package bal3

import (
	"fmt"

	"github.com/gitchander/gobal3/ternary"
)

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
	c, s = tritsSumSplit(a + b)
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
	c, s = tritsSumSplit(a - b)
	return
}

//------------------------------------------------------------------------------

func tritsAddV1(a, b int, c0 int) (s, c1 int) {
	c1, s = tritsSumSplit((a + b) + c0)
	return
}

func tritsSubV1(a, b int, c0 int) (s, c1 int) {
	c1, s = tritsSumSplit((a - b) + c0)
	return
}

//------------------------------------------------------------------------------

func tritsAddV2(a, b int, c0 int) (s, c1 int) {

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

func tritsSubV2(a, b int, c0 int) (s, c1 int) {

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

func tritsAddV3(a, b int, c0 int) (s, c1 int) {
	return ternary.FullAdder(a, b, c0)
}

func tritsSubV3(a, b int, c0 int) (s, c1 int) {
	b = ternary.Neg(b)
	return ternary.FullAdder(a, b, c0)
}

//------------------------------------------------------------------------------

// tritsAdd: (a + b) + carry
// tritsSub: (a - b) + carry
var (
	// tritsAdd = tritsAddV1
	// tritsSub = tritsSubV1

	// tritsAdd = tritsAddV2
	// tritsSub = tritsSubV2

	tritsAdd = tritsAddV3
	tritsSub = tritsSubV3
)

//------------------------------------------------------------------------------
