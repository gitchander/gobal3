package ternary

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
