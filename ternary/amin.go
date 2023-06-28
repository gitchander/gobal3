package ternary

// Negative Min, Nmin, Amin, AntiMin

// Amin table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 1 | 1 |
// +---+---+---+---+
// | 0 | 1 | 0 | 0 |
// +---+---+---+---+
// | 1 | 1 | 0 | T |
// +---+---+---+---+

func Amin(a, b int) int {
	return neg(min(a, b))
}

var _ BinaryFunc = Amin

// aminCore uses only the Amin func for all operations.
type aminCore struct{}

func (aminCore) Neg(a int) int {
	return Amin(a, a)
}

func (aminCore) Min(a, b int) int {
	return Amin(Amin(a, b), Amin(a, b))
}

func (aminCore) Max(a, b int) int {
	return Amin(Amin(a, a), Amin(b, b))
}

func (aminCore) Xor(a, b int) int {
	return Amin(Amin(Amin(a, a), b), Amin(a, Amin(b, b)))
}

func (aminCore) NotXor(a, b int) int {
	return Amin(Amin(a, b), Amin(Amin(a, a), Amin(b, b)))
}
