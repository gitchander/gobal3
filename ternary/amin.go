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

func amin(a, b Tri) Tri {
	return Neg(Min(a, b))
}

var _ BinaryFunc = amin

// AminCore uses only the amin func for all operations.
type AminCore struct{}

var _ Core = AminCore{}

func (AminCore) Neg(a Tri) Tri {
	return amin(a, a)
}

func (AminCore) Min(a, b Tri) Tri {
	return amin(amin(a, b), amin(a, b))
}

func (AminCore) Max(a, b Tri) Tri {
	return amin(amin(a, a), amin(b, b))
}

func (AminCore) Xmax(a, b Tri) Tri {
	return amin(amin(amin(a, a), b), amin(a, amin(b, b)))
}

func (AminCore) Xamax(a, b Tri) Tri {
	return amin(amin(a, b), amin(amin(a, a), amin(b, b)))
}
