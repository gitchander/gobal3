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

func tritsAddV1(a, b Trit, c0 Trit) (s, c1 Trit) {
	return fullAdder(a, b, c0)
}

func tritsAddV2(a, b Trit, c0 Trit) (s, c1 Trit) {
	t1, t0 := splitTrits(int(a + b + c0))
	s = t0
	c1 = t1
	return
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
		s     Trit
		carry Trit = c0
	)
	for i := 0; i < n; i++ {
		s, carry = tritsAdd(getTrit(x, i), getTrit(y, i), carry)
		res = setTrit(res, i, s)
	}
	return res, carry
}
