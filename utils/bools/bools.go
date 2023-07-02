package bools

func Nand(a, b bool) bool {
	return !(a && b)
}

// Use only Nand:
func Not(a bool) bool {
	return Nand(a, a)
}

func Or(a, b bool) bool {
	return Nand(Nand(a, a), Nand(b, b))
}

func And(a, b bool) bool {
	return Nand(Nand(a, b), Nand(a, b))
}

func Xor(a, b bool) bool {
	return Nand(Nand(Nand(a, a), b), Nand(a, Nand(b, b)))
}

//------------------------------------------------------------------------------

// +---+---+---+
// |AND| F | T |
// +---+---+---+
// | F | F | F |
// +---+---+---+
// | T | F | T |
// +---+---+---+

// +---+---+---+
// |OR | F | T |
// +---+---+---+
// | F | F | T |
// +---+---+---+
// | T | T | T |
// +---+---+---+

// +---+---+---+
// |NOR| F | T |
// +---+---+---+
// | F | T | F |
// +---+---+---+
// | T | F | F |
// +---+---+---+

func Nor(a, b bool) bool {
	return !(a || b)
}
