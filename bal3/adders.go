package bal3

// http://homepage.divms.uiowa.edu/%7Ejones/ternary/arith.shtml

//------------------------------------------------------------------------------

// Add table:

// +----+----+----+----+
// |  + |  T |  0 |  1 |
// +----+----+----+----+
// |  T | T1 | 0T | 00 |
// +----+----+----+----+
// |  0 | 0T | 00 | 01 |
// +----+----+----+----+
// |  1 | 00 | 01 | 1T |
// +----+----+----+----+

//------------------------------------------------------------------------------

// +---+---+----+---+---+
// | A | B |A+B |con|sum|
// +---+---+----+---+---+
// | T | T | T1 | T | 1 |
// | T | 0 | 0T | 0 | T |
// | T | 1 | 00 | 0 | 0 |
// | 0 | T | 0T | 0 | T |
// | 0 | 0 | 00 | 0 | 0 |
// | 0 | 1 | 01 | 0 | 1 |
// | 1 | T | 00 | 0 | 0 |
// | 1 | 0 | 01 | 0 | 1 |
// | 1 | 1 | 1T | 1 | T |
// +---+---+----+---+---+

//------------------------------------------------------------------------------

// Sum table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | 1 | T | 0 |
// +---+---+---+---+
// | 0 | T | 0 | 1 |
// +---+---+---+---+
// | 1 | 0 | 1 | T |
// +---+---+---+---+

var tableAddSum = mustParseTable(
	"1T0",
	"T01",
	"01T",
)

func tritsAddSumV1(a, b Trit) Trit {
	return tritByTable(tableAddSum, a, b)
}

func tritsAddSumV2(a, b Trit) Trit {

	// (a + b) = ((a = -1) ∧ (b - 1)) ∨ ((a = 0) ∧ (b)) ∨ ((a = +1) ∧ (b + 1))
	// where:
	// ∧ - min
	// ∨ - max
	// (b - 1) - dec(b)
	// (b + 1) - inc(b)

	var (
		v1 = trico.Min(trico.Is(a, -1), trico.Dec(b))
		v2 = trico.Min(trico.Is(a, 0), b)
		v3 = trico.Min(trico.Is(a, +1), trico.Inc(b))
	)

	return trico.Max(v1, trico.Max(v2, v3))
}

var (
	//tritsAddSum = tritsAddSumV1
	tritsAddSum = tritsAddSumV2
)

//------------------------------------------------------------------------------

// Consider table:

// +---+---+---+---+
// |   | T | 0 | 1 |
// +---+---+---+---+
// | T | T | 0 | 0 |
// +---+---+---+---+
// | 0 | 0 | 0 | 0 |
// +---+---+---+---+
// | 1 | 0 | 0 | 1 |
// +---+---+---+---+

var tableAddCons = mustParseTable(
	"T00",
	"000",
	"001",
)

func tritsAddConsV1(a, b Trit) Trit {
	return tritByTable(tableAddCons, a, b)
}

func tritsAddConsV2(a, b Trit) Trit {
	if v := Trit(-1); (a == v) && (b == v) {
		return v
	}
	if v := Trit(+1); (a == v) && (b == v) {
		return v
	}
	return 0
}

func tritsAddConsV3(a, b Trit) Trit {
	var (
		v1 = trico.Min(a, b)
		v2 = trico.Min(trico.Inverse(trico.Is(a, -1)), 0)
		v3 = trico.Min(trico.Inverse(trico.Is(b, -1)), 0)
	)
	return trico.Max(v1, trico.Max(v2, v3))
}

func tritsAddConsV4(a, b Trit) Trit {
	return trico.Max(trico.Min(a, b), trico.Min(0, trico.Max(a, b)))
}

var (
	//tritsAddCons = tritsAddConsV1
	//tritsAddCons = tritsAddConsV2
	//tritsAddCons = tritsAddConsV3
	tritsAddCons = tritsAddConsV4
)

//------------------------------------------------------------------------------

// A balanced ternary "Half Adder"

// halfAdder:

// c - carry out

func tritsAdd2(a, b Trit) (sum, c Trit) {
	sum = tritsAddSum(a, b)
	c = tritsAddCons(a, b)
	return
}

//------------------------------------------------------------------------------

// A balanced ternary "Full Adder"

// fullAdder:

// c0 - carryIn (input carry trit)
// c1 - carryOut (output carry trit)

func tritsAdd3(a, b, c0 Trit) (sum, c1 Trit) {
	var (
		s1, x1 = tritsAdd2(a, b)
		s2, x2 = tritsAdd2(s1, c0)
		s3     = tritsAddSum(x1, x2)
	)
	sum = s2
	c1 = s3
	return sum, c1
}

//------------------------------------------------------------------------------
