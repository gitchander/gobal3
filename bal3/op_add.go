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

// c0 - carryIn (input carry trit)
// c1 - carryOut (output carry trit)

// z = (x + y) + c0

func trytesAdd[Tryte GenericTryte](n int, x, y Tryte, c0 Trit) (z Tryte, c1 Trit) {
	var (
		carry, t Trit
	)
	carry = c0
	for i := 0; i < n; i++ {
		carry, t = tritsAdd3(getTrit(x, i), getTrit(y, i), carry)
		z = setTrit(z, i, t)
	}
	c1 = carry
	return z, c1
}
