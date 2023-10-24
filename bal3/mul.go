package bal3

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

func tritsMulV1(a, b Trit) Trit {
	return tritByTable(tableMul, a, b)
}

func tritsMulV2(a, b Trit) Trit {
	return a * b
}

func tritsMulV3(a, b Trit) Trit {
	return terNegXor(a, b)
}

var (
	// tritsMul = tritsMulV1
	tritsMul = tritsMulV2
	//tritsMul = tritsMulV3
)
