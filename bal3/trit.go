package bal3

import (
	"fmt"
)

// trit bits:
const (
	bitsPerTrit = 2

	tbs_Mask = 0b_11

	tbs_T = 0b_01 // -1
	tbs_0 = 0b_00 //  0
	tbs_1 = 0b_10 // +1
)

type Trit int

func tritsCompare(a, b Trit) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

var tritValues = [...]Trit{
	tv_T,
	tv_0,
	tv_1,
}

func tritToChar(t Trit) (c byte, ok bool) {
	switch t {
	case tv_T:
		c = tc_T
	case tv_0:
		c = tc_0
	case tv_1:
		c = tc_1
	default:
		return 0, false
	}
	return c, true
}

func charToTrit(c byte) (t Trit, ok bool) {
	switch c {
	case tc_T:
		t = tv_T
	case tc_0:
		t = tv_0
	case tc_1:
		t = tv_1
	default:
		return 0, false
	}
	return t, true
}

func mustTritToChar(t Trit) byte {
	c, ok := tritToChar(t)
	if !ok {
		panic(errInvalidTrit(t))
	}
	return c
}

// Converting bits to trit

// | bits: |
// | 1   0 | trit |
// +---+---+------+
// | 0 | 0 |   0  |
// | 0 | 1 |  -1  |
// | 1 | 0 |   1  |
// | 1 | 1 |   0  |

// +-------+-------+-------+-------+-------+
// | bits  | 7 | 6 | 5 | 4 | 3 | 2 | 1 | 0 |
// +-------+-------+-------+-------+-------+
// | trits |   3   |   2   |   1   |   0   |
// +-------+-------+-------+-------+-------+

func errInvalidTrit(t Trit) error {
	return fmt.Errorf("invalid trit value %d", t)
}

func bitsToTrit[T Unsigned](x T) Trit {
	x &= tbs_Mask
	switch x {
	case tbs_T:
		return tv_T
	case tbs_1:
		return tv_1
	default:
		return tv_0
	}
}

func tritToBits[T Unsigned](t Trit) T {
	switch t {
	case tv_T:
		return tbs_T
	case tv_0:
		return tbs_0
	case tv_1:
		return tbs_1
	default:
		panic(errInvalidTrit(t))
	}
}

func setTrit[T Unsigned](x T, i int, t Trit) T {

	offset := i * bitsPerTrit

	x &^= T(tbs_Mask) << offset // reset trit bits

	y := tritToBits[T](t)

	x |= y << offset

	return x
}

func getTrit[T Unsigned](x T, i int) Trit {

	offset := i * bitsPerTrit

	x = (x >> offset) & tbs_Mask

	return bitsToTrit(x)
}

func setTritsN[T Unsigned](n int, t Trit) T {
	var y T
	for i := 0; i < n; i++ {
		y = (y << bitsPerTrit) | (tritToBits[T](t))
	}
	return y
}

//------------------------------------------------------------------------------

// Mul table

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 0 | T |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | T | 0 | 1 |
// +---+---+---+---+

var tableMul = mustParseTable(
	"10T",
	"000",
	"T01",
)

func tritsMulV1(a, b Trit) Trit {
	return tritByTable(tableMul, a, b)
}

func tritsMulV2(a, b Trit) Trit {
	return a * b
}

func tritsMulV3(a, b Trit) Trit {
	return terNegXor(a, b)
}

var (
	// tritsMul = tritsMulV1
	tritsMul = tritsMulV2
	//tritsMul = tritsMulV3
)
