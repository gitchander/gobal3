package bal3

import (
	"github.com/gitchander/gobal3/ternary"
)

// Mul table

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | 0 | T |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | T | 0 | 1 |
// +---+---+---+---+

func tritsMul(a, b int) int {
	// return a * b
	return ternary.NotXor(a, b)
}
