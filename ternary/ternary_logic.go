package ternary

//------------------------------------------------------------------------------

func checkTrit(t int) {
	switch t {
	case -1, 0, 1:
	default:
		panic(errInvalidTrit(t))
	}
}

func checkTrits(ts ...int) {
	for _, t := range ts {
		checkTrit(t)
	}
}

//------------------------------------------------------------------------------

// Invert table:

// +---+---+
// | t |inv|
// +---+---+
// | T | 1 |
// +---+---+
// | 0 | 0 |
// +---+---+
// | 1 | T |
// +---+---+

func invertTrit(t int) int {
	switch t {
	case -1:
		return 1
	case 0:
		return 0
	case 1:
		return -1
	default:
		panic(errInvalidTrit(t))
	}
}

func neg(a int) int {
	return -a
}

// NEG - NEGATIVE
func Neg(a int) int {
	checkTrit(a)
	return neg(a)
}

var _ UnaryFunc = Neg

//------------------------------------------------------------------------------

// Trit inc table:

// +---+---+
// | t |inc|
// +---+---+
// | T | 0 |
// +---+---+
// | 0 | 1 |
// +---+---+
// | 1 | T |
// +---+---+

func tritInc(a int) int {
	switch a {
	case -1:
		return 0
	case 0:
		return 1
	case 1:
		return -1
	default:
		panic(errInvalidTrit(a))
	}
}

//------------------------------------------------------------------------------

// Trit dec table:

// +---+---+
// | t |dec|
// +---+---+
// | T | 1 |
// +---+---+
// | 0 | T |
// +---+---+
// | 1 | 0 |
// +---+---+

func tritDec(a int) int {
	switch a {
	case -1:
		return 1
	case 0:
		return -1
	case 1:
		return 0
	default:
		panic(errInvalidTrit(a))
	}
}

//------------------------------------------------------------------------------

// "Min" or "And"

// Min table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | T | T | T |
// +---+---+---+---+
// | 0 | T | 0 | 0 |
// +---+---+---+---+
// | 1 | T | 0 | 1 |
// +---+---+---+---+

func min(a, b int) int {
	return minInt(a, b)
}

func Min(a, b int) int {
	checkTrits(a, b)
	return min(a, b)
}

var _ BinaryFunc = min

//------------------------------------------------------------------------------

// "Max" or "Or"

// Max table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | T | 0 | 1 |
// +---+---+---+---+
// | 0 | 0 | 0 | 1 |
// +---+---+---+---+
// | 1 | 1 | 1 | 1 |
// +---+---+---+---+

func max(a, b int) int {
	return maxInt(a, b)
}

func Max(a, b int) int {
	checkTrits(a, b)
	return max(a, b)
}

var _ BinaryFunc = max

//------------------------------------------------------------------------------

// Decoders

func TritIs(a int, v int) int {
	if a == v {
		return 1 // true
	}
	return -1 // false
}

func TritIsNot(a int, v int) int {
	if a != v {
		return 1 // true
	}
	return -1 // false
}

//------------------------------------------------------------------------------

// Xor - Exclusive Or

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | T | 0 | 1 |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | 1 | 0 | T |
// +---+---+---+---+

// For binary logic: XOR = Nand(Nand(Nand(a, a), b), Nand(a, Nand(b, b)))

func Xor(a, b int) int {
	return max(min(a, neg(b)), min(neg(a), b))
}

//------------------------------------------------------------------------------

// Not Xor - Not Exclusive Or

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 0 | T |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | T | 0 | 1 |
// +---+---+---+---+

func NotXor(a, b int) int {
	return min(max(a, neg(b)), max(neg(a), b))
}

//------------------------------------------------------------------------------
