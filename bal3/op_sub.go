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

func tritsSubV1(a, b Trit, c0 Trit) (s, c1 Trit) {
	b = trico.Neg(b)
	return tritsAdd(a, b, c0)
}

func tritsSubV2(a, b Trit, c0 Trit) (s, c1 Trit) {
	t1, t0 := splitTrits(int(a - b + c0))
	s = t0
	c1 = t1
	return
}

var (
	//tritsSub = tritsSubV1
	tritsSub = tritsSubV2
)

//------------------------------------------------------------------------------

func trytesSub[T coreTryte](n int, x, y T, c0 Trit) (res T, c1 Trit) {
	var (
		s     Trit
		carry Trit = c0
	)
	for i := 0; i < n; i++ {
		s, carry = tritsSub(getTrit(x, i), getTrit(y, i), carry)
		res = setTrit(res, i, s)
	}
	return res, carry
}
