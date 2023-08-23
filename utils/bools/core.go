package bools

// XOR - https://en.wikipedia.org/wiki/XOR_gate
// XNOR - https://en.wikipedia.org/wiki/XNOR_gate

type Core interface {
	Not(a bool) bool

	Or(a, b bool) bool
	And(a, b bool) bool

	Nor(a, b bool) bool
	Nand(a, b bool) bool

	Xor(a, b bool) bool
	Xnor(a, b bool) bool
}
