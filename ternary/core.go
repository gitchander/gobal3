package ternary

type Core interface {
	Inverse(a Tri) Tri // bool: Not

	Min(a, b Tri) Tri // bool: And
	Max(a, b Tri) Tri // bool: Or

	// Amin(a, b int) int
	// Amax(a, b int) int

	Xmax(a, b Tri) Tri  // bool: Xor
	Xamax(a, b Tri) Tri // bool: Xnor
}
