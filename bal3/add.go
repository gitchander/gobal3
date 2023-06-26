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

func tritsOpAdd(a, b int) (hi, lo int) {
	return tritsSumSplit(a + b)
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

func tritsOpSub(a, b int) (hi, lo int) {
	return tritsSumSplit(a - b)
}

//------------------------------------------------------------------------------

func tritsAdd_v1(a, b int, carry int) (hi, lo int) {
	return tritsSumSplit((a + b) + carry)
}

func tritsSub_v1(a, b int, carry int) (hi, lo int) {
	return tritsSumSplit((a - b) + carry)
}

//------------------------------------------------------------------------------

func tritsAdd_v2(a, b int, carry int) (hi, lo int) {

	var (
		hi1, lo1 = tritsOpAdd(a, b)
		hi2, lo2 = tritsOpAdd(lo1, carry)
		hi3, lo3 = tritsOpAdd(hi1, hi2)
	)

	_ = hi3

	return lo3, lo2
}

func tritsSub_v2(a, b int, carry int) (hi, lo int) {

	var (
		hi1, lo1 = tritsOpSub(a, b)
		hi2, lo2 = tritsOpAdd(lo1, carry)
		hi3, lo3 = tritsOpAdd(hi1, hi2)
	)

	_ = hi3

	return lo3, lo2
}

//------------------------------------------------------------------------------

// tritsAdd: (a + b) + carry
// tritsSub: (a - b) + carry
var (
	tritsAdd = tritsAdd_v1
	tritsSub = tritsSub_v1

	// tritsAdd = tritsAdd_v2
	// tritsSub = tritsSub_v2
)
