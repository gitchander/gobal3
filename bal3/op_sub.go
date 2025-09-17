package bal3

// sub - subtraction

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

func tritsSubV1(a, b Trit, c0 Trit) (hi, lo Trit) {
	b = tritNeg(b)
	return tritsAdd(a, b, c0)
}

func tritsSubV2(a, b Trit, c0 Trit) (hi, lo Trit) {
	return splitTrits(int(a - b + c0))
}

var (
	//tritsSub = tritsSubV1
	tritsSub = tritsSubV2
)

//------------------------------------------------------------------------------

// c0 - carryIn
// c1 - carryOut

// s = x - y

func trytesSub[Tryte GenericTryte](n int, x, y Tryte, c0 Trit) (s Tryte, c1 Trit) {
	var (
		carry, t Trit
	)
	carry = c0
	for i := 0; i < n; i++ {
		carry, t = tritsSub(getTrit(x, i), getTrit(y, i), carry)
		s = setTrit(s, i, t)
	}
	c1 = carry
	return s, c1
}
