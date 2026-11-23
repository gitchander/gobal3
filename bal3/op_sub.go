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

// b0 - borrowIn (input borrow trit)
// b1 - borrowOut (output borrow trit)
// diff - difference

// diff - 3 * b1 = x - y - b0
// diff + 3 * (-b1) = x + (-y) + (-b0)

func tritsSub(x, y Trit, b0 Trit) (diff, b1 Trit) {

	y = tritInverse(y)
	b0 = tritInverse(b0)

	diff, b1 = tritsAdd3(x, y, b0)

	b1 = tritInverse(b1)

	return diff, b1
}

//------------------------------------------------------------------------------

// b0 - borrowIn (input borrow trit)
// b1 - borrowOut (output borrow trit)
// diff - difference

// diff - (3^n) * b1 = x - y - b0
// diff + (3^n) * (-b1) = x + (-y) + (-b0)

func trytesSub[Tryte GenericTryte](n int, x, y Tryte, b0 Trit) (diff Tryte, b1 Trit) {
	var (
		borrow, t Trit
	)
	borrow = b0
	for i := 0; i < n; i++ {
		t, borrow = tritsSub(getTrit(x, i), getTrit(y, i), borrow)
		diff = setTrit(diff, i, t)
	}
	b1 = borrow
	return diff, b1
}

//------------------------------------------------------------------------------
