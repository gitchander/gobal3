package ternary

// Negative Max, Nmax, Amax, AntiMax

// Amax table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 0 | T |
// +---+---+---+---+
// | 0 | 0 | 0 | T |
// +---+---+---+---+
// | 1 | T | T | T |
// +---+---+---+---+

func Amax(a, b int) int {
	return Neg(Max(a, b))
}

var _ BinaryFunc = Amax

// amaxCore uses only the Amax func for all operations.
type amaxCore struct{}

func (amaxCore) Neg(a int) int {
	return Amax(a, a)
}

func (amaxCore) Min(a, b int) int {
	return Amax(Amax(a, a), Amax(b, b))
}

func (amaxCore) Max(a, b int) int {
	return Amax(Amax(a, b), Amax(a, b))
}

func (amaxCore) Xor(a, b int) int {
	return Amax(Amax(a, b), Amax(Amax(a, a), Amax(b, b)))
}

func (amaxCore) NegXor(a, b int) int {
	return Amax(Amax(Amax(a, a), b), Amax(a, Amax(b, b)))
}
