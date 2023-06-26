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

func errInvalidTrit(t int) error {
	return fmt.Errorf("invalid trit value %d", t)
}

func bitsToTrit[T Unsigned](x T) int {
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

func tritToBits[T Unsigned](t int) T {
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

func setTrit[T Unsigned](x T, i int, t int) T {

	offset := i * bitsPerTrit

	x &^= T(tbs_Mask) << offset // reset trit bits

	y := tritToBits[T](t)

	x |= y << offset

	return x
}

func getTrit[T Unsigned](x T, i int) int {

	offset := i * bitsPerTrit

	x = (x >> offset) & tbs_Mask

	return bitsToTrit(x)
}
