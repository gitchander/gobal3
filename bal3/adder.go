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

// Sub table:

// +----+----+----+----+
// |  - |  T |  0 |  1 |
// +----+----+----+----+
// |  T | 00 | 0T | T1 |
// +----+----+----+----+
// |  0 | 01 | 00 | 0T |
// +----+----+----+----+
// |  1 | 1T | 01 | 00 |
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

func addSumV1(a, b Trit) Trit {
	return tritByTable(tableAddSum, a, b)
}

func addSumV2(a, b Trit) Trit {
	_, t0 := splitTrits(int(a + b))
	return t0
}

func addSumV3(a, b Trit) Trit {
	c := a + b
	if c < -1 {
		c += 3
	}
	if c > +1 {
		c -= 3
	}
	return c
}

func addSumV4(a, b Trit) Trit {

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
	//addSum = addSumV1
	addSum = addSumV2
	//addSum = addSumV3
	//addSum = addSumV4
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

func addConsV1(a, b Trit) Trit {
	return tritByTable(tableAddCons, a, b)
}

func addConsV2(a, b Trit) Trit {
	t1, _ := splitTrits(int(a + b))
	return t1
}

func addConsV3(a, b Trit) Trit {
	if v := Trit(tv_T); (a == v) && (b == v) {
		return v
	}
	if v := Trit(tv_1); (a == v) && (b == v) {
		return v
	}
	return 0
}

func addConsV4(a, b Trit) Trit {
	var (
		v1 = trico.Min(a, b)
		v2 = trico.Min(trico.Neg(trico.Is(a, -1)), 0)
		v3 = trico.Min(trico.Neg(trico.Is(b, -1)), 0)
	)
	return trico.Max(v1, trico.Max(v2, v3))
}

func addConsV5(a, b Trit) Trit {
	return trico.Max(trico.Min(a, b), trico.Min(0, trico.Max(a, b)))
}

var (
	//addCons = addConsV1
	addCons = addConsV2
	//addCons = addConsV3
	//addCons = addConsV4
	//addCons = addConsV5
)

//------------------------------------------------------------------------------

// A balanced ternary half adder

// c - carryOut

func halfAdderV1(a, b Trit) (s, c Trit) {
	t1, t0 := splitTrits(int(a + b))
	s = t0
	c = t1
	return
}

func halfAdderV2(a, b Trit) (s, c Trit) {
	s = addSum(a, b)
	c = addCons(a, b)
	return
}

var (
	halfAdder = halfAdderV1
	//halfAdder = halfAdderV2
)

//------------------------------------------------------------------------------

// Balanced Full Adder

// c0 - carryIn
// c1 - carryOut

func fullAdder(a, b Trit, c0 Trit) (s, c1 Trit) {
	var (
		s1, x1 = halfAdder(a, b)
		s2, x2 = halfAdder(s1, c0)
		s3     = addSum(x1, x2)
	)
	s = s2
	c1 = s3
	return
}

//------------------------------------------------------------------------------
