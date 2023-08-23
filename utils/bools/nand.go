package bools

func nand(a, b bool) bool {
	return !(a && b)
}

// Use only Nand
type NandCore struct{}

var _ Core = NandCore{}

func (NandCore) Not(a bool) bool {
	return nand(a, a)
}

func (NandCore) Or(a, b bool) bool {
	return nand(nand(a, a), nand(b, b))
}

func (NandCore) And(a, b bool) bool {
	return nand(nand(a, b), nand(a, b))
}

func (NandCore) Nor(a, b bool) bool {

	// return nand(nand(nand(a, a), nand(b, b)), nand(nand(a, a), nand(b, b)))

	c := nand(nand(a, a), nand(b, b)) // or
	return nand(c, c)                 // not
}

func (NandCore) Nand(a, b bool) bool {
	return nand(a, b)
}

func (NandCore) Xor(a, b bool) bool {
	return nand(nand(nand(a, a), b), nand(a, nand(b, b)))
}

func (NandCore) Xnor(a, b bool) bool {
	return nand(nand(a, b), nand(nand(a, a), nand(b, b)))
}
