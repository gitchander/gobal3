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

var tableMul = mustParseTable(
	"10T",
	"000",
	"T01",
)

func tritsMulV1(a, b int) int {
	return tritByTable(tableMul, a, b)
}

func tritsMulV2(a, b int) int {
	return a * b
}

func tritsMulV3(a, b int) int {
	return ternary.NegXor(a, b)
}

var (
	// tritsMul = tritsMulV1
	tritsMul = tritsMulV2
	//tritsMul = tritsMulV3
)
