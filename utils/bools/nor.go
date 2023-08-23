package bools

func nor(a, b bool) bool {
	return !(a || b)
}

// Use only Nor
type NorCore struct{}

var _ Core = NorCore{}

func (NorCore) Not(a bool) bool {
	return nor(a, a)
}

func (NorCore) Or(a, b bool) bool {
	return nor(nor(a, b), nor(a, b))
}

func (NorCore) And(a, b bool) bool {
	return nor(nor(a, a), nor(b, b))
}

func (NorCore) Nor(a, b bool) bool {
	return nor(a, b)
}

func (NorCore) Nand(a, b bool) bool {

	// return nor(nor(nor(a, a), nor(b, b)), nor(nor(a, a), nor(b, b)))

	c := nor(nor(a, a), nor(b, b)) // and
	return nor(c, c)               // not
}

func (NorCore) Xor(a, b bool) bool {
	return nor(nor(a, b), nor(nor(a, a), nor(b, b)))
}

func (NorCore) Xnor(a, b bool) bool {
	return nor(nor(nor(a, a), b), nor(a, nor(b, b)))
}
