package bools

type BaseCore struct{}

var _ Core = BaseCore{}

func (BaseCore) Not(a bool) bool {
	return !a
}

func (BaseCore) Or(a, b bool) bool {
	return a || b
}

func (BaseCore) And(a, b bool) bool {
	return a && b
}

func (BaseCore) Nor(a, b bool) bool {
	return !(a || b)
}

func (BaseCore) Nand(a, b bool) bool {
	return !(a && b)
}

func (BaseCore) Xor(a, b bool) bool {
	return a != b
}

func (BaseCore) Xnor(a, b bool) bool {
	return a == b
}
