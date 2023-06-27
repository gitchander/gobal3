package bal3

// https://en.wikipedia.org/wiki/Three-valued_logic#Kleene_logic
// http://homepage.divms.uiowa.edu/%7Ejones/ternary/logic.shtml

//------------------------------------------------------------------------------

// Unary operation
// https://en.wikipedia.org/wiki/Unary_operation

type UnaryFunc func(int) int

// Binary operation
// https://en.wikipedia.org/wiki/Binary_operation

type BinaryFunc func(int, int) int

//------------------------------------------------------------------------------

func checkTrit(t int) {

	// switch t {
	// case tv_T, tv_0, tv_1:
	// default:
	// 	panic(errInvalidTrit(t))
	// }

	if (t < tv_T) || (tv_1 < t) {
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
	case tv_T:
		return tv_1
	case tv_0:
		return tv_0
	case tv_1:
		return tv_T
	default:
		panic(errInvalidTrit(t))
	}
}

// NEG - NEGATIVE
func Neg(a int) int {
	return invertTrit(a)
}

var _ UnaryFunc = Neg

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
	checkTrits(a, b)
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
	checkTrits(a, b)
	return maxInt(a, b)
}

var _ BinaryFunc = Max

//------------------------------------------------------------------------------

// AntiMin table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 1 | 1 |
// +---+---+---+---+
// | 0 | 1 | 0 | 0 |
// +---+---+---+---+
// | 1 | 1 | 0 | T |
// +---+---+---+---+

func AntiMin(a, b int) int {
	return Neg(Min(a, b))
}

var _ BinaryFunc = AntiMin

//------------------------------------------------------------------------------

// AntiMax table

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 0 | T |
// +---+---+---+---+
// | 0 | 0 | 0 | T |
// +---+---+---+---+
// | 1 | T | T | T |
// +---+---+---+---+

func AntiMax(a, b int) int {
	return Neg(Max(a, b))
}

var _ BinaryFunc = AntiMax

//------------------------------------------------------------------------------

// Exclusive Or

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
	amin := AntiMin
	return amin(amin(amin(a, a), b), amin(a, amin(b, b)))
}

//------------------------------------------------------------------------------

// Implication for Kleene logic

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 1 | 1 |
// +---+---+---+---+
// | 0 | 0 | 0 | 1 |
// +---+---+---+---+
// | 1 | T | 0 | 1 |
// +---+---+---+---+

func Imp(a, b int) int {
	return Max(Neg(a), b)
}

//------------------------------------------------------------------------------

type AminCore struct{}

var amin = AntiMin

func (AminCore) Neg(a int) int {
	return amin(a, a)
}

func (AminCore) Min(a, b int) int {
	return amin(amin(a, b), amin(a, b))
}

func (AminCore) Max(a, b int) int {
	return amin(amin(a, a), amin(b, b))
}

func (AminCore) Xor(a, b int) int {
	return amin(amin(amin(a, a), b), amin(a, amin(b, b)))
}

//------------------------------------------------------------------------------

type AmaxCore struct{}

var amax = AntiMax

func (AmaxCore) Neg(a int) int {
	return amax(a, a)
}

func (AmaxCore) Min(a, b int) int {
	return amax(amax(a, a), amax(b, b))
}

func (AmaxCore) Max(a, b int) int {
	return amax(amax(a, b), amax(a, b))
}

func (AmaxCore) Xor(a, b int) int {
	return amax(amax(a, b), amax(amax(a, a), amax(b, b)))
}
