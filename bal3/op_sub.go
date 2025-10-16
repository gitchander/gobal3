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
	b = tritInverse(b)
	return tritsAdd3(a, b, c0)
}

var (
	tritsSub = tritsSubV1
)

//------------------------------------------------------------------------------

// c0 - carryIn (input carry trit)
// c1 - carryOut (output carry trit)

// z = (x - y) + c0

func trytesSub[Tryte GenericTryte](n int, x, y Tryte, c0 Trit) (z Tryte, c1 Trit) {
	var (
		carry, t Trit
	)
	carry = c0
	for i := 0; i < n; i++ {
		carry, t = tritsSub(getTrit(x, i), getTrit(y, i), carry)
		z = setTrit(z, i, t)
	}
	c1 = carry
	return z, c1
}
