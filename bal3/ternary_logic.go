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

// Max table

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

// Min table

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
