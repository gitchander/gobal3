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

// NEG - NEGATIVE
// Anti

func Neg(a int) int {
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
	a--
	if a < -1 {
		a += 3
	}
	return a
}

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/AND_gate

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
	return minInt(a, b)
}

var _ BinaryFunc = Min

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/OR_gate

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
	return maxInt(a, b)
}

var _ BinaryFunc = Max

//------------------------------------------------------------------------------

// Decoders

// Is tables:

// +---+---+---+---+
// | is| T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | T | T |
// +---+---+---+---+
// | 0 | T | 1 | T |
// +---+---+---+---+
// | 1 | T | T | 1 |
// +---+---+---+---+

func Is(a int, v int) int {
	if a == v {
		return 1 // true
	}
	return -1 // false
}

// (a == -1)
func IsNegative(a int) int {
	return Is(a, -1)
}

// (a == 0)
func IsZero(a int) int {
	return Is(a, 0)
}

// (a == +1)
func IsPositive(a int) int {
	return Is(a, +1)
}

var (
	_ UnaryFunc = IsNegative
	_ UnaryFunc = IsZero
	_ UnaryFunc = IsPositive
)

//------------------------------------------------------------------------------

func NegIs(a int, v int) int {
	if a != v {
		return 1
	}
	return -1
}

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/XOR_gate

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
