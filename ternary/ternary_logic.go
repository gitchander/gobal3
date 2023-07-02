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

// func invertTrit(t int) int {
// 	switch t {
// 	case -1:
// 		return 1
// 	case 0:
// 		return 0
// 	case 1:
// 		return -1
// 	default:
// 		panic(errInvalidTrit(t))
// 	}
// }

// NEG - NEGATIVE
// Anti

func Neg(a int) int {
	// checkTrit(a)
	return -a
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

func Inc(a int) int {

	// switch a {
	// case -1:
	// 	return 0
	// case 0:
	// 	return 1
	// case 1:
	// 	return -1
	// default:
	// 	panic(errInvalidTrit(a))
	// }

	a++
	if a > 1 {
		a -= 3
	}
	return a
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

func Dec(a int) int {

	// switch a {
	// case -1:
	// 	return 1
	// case 0:
	// 	return -1
	// case 1:
	// 	return 0
	// default:
	// 	panic(errInvalidTrit(a))
	// }

	a--
	if a < -1 {
		a += 3
	}
	return a
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

func Min(a, b int) int {
	// checkTrits(a, b)
	return minInt(a, b)
}

var _ BinaryFunc = Min

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

func Max(a, b int) int {
	// checkTrits(a, b)
	return maxInt(a, b)
}

var _ BinaryFunc = Max

//------------------------------------------------------------------------------

// Decoders

func Is(a int, v int) int {
	if a == v {
		return 1 // true
	}
	return -1 // false
}

func NegIs(a int, v int) int {
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
	return Max(Min(a, Neg(b)), Min(Neg(a), b))
}

//------------------------------------------------------------------------------

// Neg Xor - Not Exclusive Or

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 0 | T |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | T | 0 | 1 |
// +---+---+---+---+

func NegXor(a, b int) int {
	return Min(Max(a, Neg(b)), Max(Neg(a), b))
}

//------------------------------------------------------------------------------
