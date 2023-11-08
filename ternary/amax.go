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

func amax(a, b Tri) Tri {
	return Neg(Max(a, b))
}

var _ BinaryFunc = amax

// AmaxCore uses only the amax func for all operations.
type AmaxCore struct{}

var _ Core = AmaxCore{}

func (AmaxCore) Neg(a Tri) Tri {
	return amax(a, a)
}

func (AmaxCore) Min(a, b Tri) Tri {
	return amax(amax(a, a), amax(b, b))
}

func (AmaxCore) Max(a, b Tri) Tri {
	return amax(amax(a, b), amax(a, b))
}

func (AmaxCore) Xmax(a, b Tri) Tri {
	return amax(amax(a, b), amax(amax(a, a), amax(b, b)))
}

func (AmaxCore) Xamax(a, b Tri) Tri {
	return amax(amax(amax(a, a), b), amax(a, amax(b, b)))
}
