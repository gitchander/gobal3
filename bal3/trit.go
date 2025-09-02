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

func charToTrit(char byte) (Trit, error) {
	var t Trit
	switch char {
	case tc_T:
		t = tv_T
	case tc_0:
		t = tv_0
	case tc_1:
		t = tv_1
	default:
		return 0, fmt.Errorf("invalid trit char %q", char)
	}
	return t, nil
}

func mustTritToChar(t Trit) byte {
	c, ok := tritToChar(t)
	if !ok {
		panic(errInvalidTrit(t))
	}
	return c
}

// Converting bits to trit

// +-------+------+
// | bits: |      |
// +---+---+ trit |
// | 1 | 0 |      |
// +---+---+------+
// | 0 | 0 |   0  |
// | 0 | 1 |  -1  |
// | 1 | 0 |  +1  |
// | 1 | 1 |   0  |
// +---+---+------+

// +-------+-------+-------+-------+-------+
// | bits  | 7 | 6 | 5 | 4 | 3 | 2 | 1 | 0 |
// +-------+-------+-------+-------+-------+
// | trits |   3   |   2   |   1   |   0   |
// +-------+-------+-------+-------+-------+

func errInvalidTrit(t Trit) error {
	return fmt.Errorf("invalid trit value %d", t)
}

var tableBitsToTrit = [...]Trit{
	0: 0,  // 0 (00) ->  0
	1: -1, // 1 (01) -> -1
	2: +1, // 2 (10) -> +1
	3: 0,  // 3 (11) ->  0
}

func getTrit[T CoreTryte](x T, i int) Trit {

	offset := i * bitsPerTrit

	x = (x >> offset) & tbs_Mask

	return tableBitsToTrit[x]
}

var tableTritToBits = [...]byte{
	0: 0b_01, // ((-1 + 1) = 0) -> 01
	1: 0b_00, // (( 0 + 1) = 1) -> 00
	2: 0b_10, // ((+1 + 1) = 2) -> 10
}

func setTrit[T CoreTryte](x T, i int, t Trit) T {

	offset := i * bitsPerTrit

	x &^= T(tbs_Mask) << offset // reset trit bits

	y := T(tableTritToBits[t+1])

	x |= y << offset

	return x
}

//------------------------------------------------------------------------------

// sum 4 trits: [-4..4]
func splitTrits1(v int) (hi, lo Trit) {
	const (
		N = -1
		Z = 0
		P = +1
	)
	switch v {
	case -4:
		return N, N
	case -3:
		return N, Z
	case -2:
		return N, P
	case -1:
		return Z, N
	case 0:
		return Z, Z
	case +1:
		return Z, P
	case +2:
		return P, N
	case +3:
		return P, Z
	case +4:
		return P, P
	default:
		panic(fmt.Errorf("splitTrits1: invalid value %d", v))
	}
}

func splitTrits2(v int) (hi, lo Trit) {
	switch v {
	case -4:
		return tv_T, tv_T
	case -3:
		return tv_T, tv_0
	case -2:
		return tv_T, tv_1
	case -1:
		return tv_0, tv_T
	case 0:
		return tv_0, tv_0
	case +1:
		return tv_0, tv_1
	case +2:
		return tv_1, tv_T
	case +3:
		return tv_1, tv_0
	case +4:
		return tv_1, tv_1
	default:
		panic(fmt.Errorf("splitTrits2: invalid value %d", v))
	}
}

var (
	splitTrits = splitTrits1
	// splitTrits = splitTrits2
)
