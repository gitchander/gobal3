package ternary

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

func Neg(a Tri) Tri {
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

func Inc(a Tri) Tri {
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

func Dec(a Tri) Tri {
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

func Min(a, b Tri) Tri {
	return Min2(a, b)
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

func Max(a, b Tri) Tri {
	return Max2(a, b)
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

// Equal, Eq

func Is(a Tri, v Tri) Tri {
	if a == v {
		return 1 // true
	}
	return -1 // false
}

//------------------------------------------------------------------------------

func NegIs(a Tri, v Tri) Tri {
	if a != v {
		return 1
	}
	return -1
}

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/XOR_gate

// Xmax - Exclusive max (bool: xor)

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

func Xmax(a, b Tri) Tri {
	return Max(Min(a, Neg(b)), Min(Neg(a), b))
}

//------------------------------------------------------------------------------

// Xamax - Anti Exclusive max (bool: xnor)

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 0 | T |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | T | 0 | 1 |
// +---+---+---+---+

func Xamax(a, b Tri) Tri {
	return Min(Max(a, Neg(b)), Max(Neg(a), b))
}

//------------------------------------------------------------------------------
