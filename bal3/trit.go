package bal3

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Balanced_ternary

//------------------------------------------------------------------------------

type Trit int

var allTrits = [...]Trit{-1, 0, +1}

//------------------------------------------------------------------------------

var mapTritToChar = map[Trit]byte{
	-1: 'T',
	0:  '0',
	+1: '1',
}

var mapCharToTrit = map[byte]Trit{

	// Representations of -1:
	'T': -1,
	'N': -1,
	'-': -1,

	// 'Θ' - Greek letter theta
	// 'Θ': -1,

	// Representations of 0:
	'0': 0,
	'Z': 0,
	'|': 0,

	// Representations of +1:
	'1': 1,
	'P': 1,
	'+': 1,
}

func tritToChar(t Trit) (byte, error) {
	char, ok := mapTritToChar[t]
	if ok {
		return char, nil
	}
	return 0, fmt.Errorf("invalid trit value %d", t)
}

func charToTrit(char byte) (Trit, error) {
	t, ok := mapCharToTrit[char]
	if ok {
		return t, nil
	}
	return 0, fmt.Errorf("invalid trit char %q", char)
}

//------------------------------------------------------------------------------

// var (
// 	tritChars1 = [...]byte{'T', '0', '1'}
// 	tritChars2 = [...]byte{'N', '0', '1'}
// 	tritChars3 = [...]byte{'N', 'Z', 'P'}
// 	tritChars4 = [...]byte{'-', '0', '+'}
// 	tritChars5 = [...]byte{'-', '|', '+'}

// 	tritChars = tritChars1
// )

// func tritToChar(t Trit) (char byte, err error) {
// 	switch t {
// 	case -1:
// 		char = tritChars[0]
// 	case 0:
// 		char = tritChars[1]
// 	case +1:
// 		char = tritChars[2]
// 	default:
// 		return 0, fmt.Errorf("invalid trit value %d", t)
// 	}
// 	return char, nil
// }

// func charToTrit(char byte) (Trit, error) {
// 	var t Trit
// 	switch char {
// 	case tritChars[0]:
// 		t = -1
// 	case tritChars[1]:
// 		t = 0
// 	case tritChars[2]:
// 		t = +1
// 	default:
// 		return 0, fmt.Errorf("invalid trit char %q", char)
// 	}
// 	return t, nil
// }

//------------------------------------------------------------------------------

// Converting bits to trit

// trit bits:

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

const (
	bitsPerTrit = 2

	tbs_Mask = 0b_11
)

var tableBitsToTrit = [...]Trit{
	0: 0,  // 0 (00) ->  0
	1: -1, // 1 (01) -> -1
	2: +1, // 2 (10) -> +1
	3: 0,  // 3 (11) ->  0
}

func getTrit[T GenericTryte](x T, i int) Trit {

	offset := i * bitsPerTrit

	j := (x >> offset) & tbs_Mask

	return tableBitsToTrit[j]
}

var tableTritToBits = [...]byte{
	0: 0b_01, // ((-1 + 1) = 0) -> 01
	1: 0b_00, // (( 0 + 1) = 1) -> 00
	2: 0b_10, // ((+1 + 1) = 2) -> 10
}

func setTrit[T GenericTryte](x T, i int, t Trit) T {

	offset := i * bitsPerTrit

	x &^= T(tbs_Mask) << offset // reset trit bits

	index := tritToIndex(t)
	y := T(tableTritToBits[index])

	x |= y << offset

	return x
}

//------------------------------------------------------------------------------

// sum of 4 trits: [-4..4]

func splitTrits1(v int) (hi, lo Trit) {
	const (
		v0 = 0
	)
	switch v {
	case -4:
		return -1, -1
	case -3:
		return -1, v0
	case -2:
		return -1, +1
	case -1:
		return v0, -1
	case 0:
		return v0, v0
	case +1:
		return v0, +1
	case +2:
		return +1, -1
	case +3:
		return +1, v0
	case +4:
		return +1, +1
	default:
		panic(fmt.Errorf("splitTrits1: invalid value %d", v))
	}
}

func splitTrits2(v int) (hi, lo Trit) {
	const (
		vN = -1
		vZ = 0
		vP = +1
	)
	switch v {
	case -4:
		return vN, vN
	case -3:
		return vN, vZ
	case -2:
		return vN, vP
	case -1:
		return vZ, vN
	case 0:
		return vZ, vZ
	case +1:
		return vZ, vP
	case +2:
		return vP, vN
	case +3:
		return vP, vZ
	case +4:
		return vP, vP
	default:
		panic(fmt.Errorf("splitTrits2: invalid value %d", v))
	}
}

var (
	splitTrits = splitTrits1
	// splitTrits = splitTrits2
)
