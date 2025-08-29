package bal3

// add - addition

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

func tritsAddV1(a, b Trit, c0 Trit) (hi, lo Trit) {
	s, c1 := fullAdder(a, b, c0)

	hi = c1
	lo = s

	return hi, lo
}

func tritsAddV2(a, b Trit, c0 Trit) (hi, lo Trit) {
	return splitTrits(int(a + b + c0))
}

var (
	//tritsAdd = tritsAddV1
	tritsAdd = tritsAddV2
)

//------------------------------------------------------------------------------

// c0 - carryIn
// c1 - carryOut

func trytesAdd[T coreTryte](n int, x, y T, c0 Trit) (res T, c1 Trit) {
	var (
		carry, t Trit
	)
	carry = c0
	for i := 0; i < n; i++ {
		carry, t = tritsAdd(getTrit(x, i), getTrit(y, i), carry)
		res = setTrit(res, i, t)
	}
	return res, carry
}
