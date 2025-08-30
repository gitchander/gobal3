package bal3

// mul - multiplication

//------------------------------------------------------------------------------

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
	return a * b
}

func tritsMulV2(a, b Trit) Trit {
	return tritByTable(tableMul, a, b)
}

func tritsMulV3(a, b Trit) Trit {
	return trico.Xamax(a, b)
}

var (
	tritsMul = tritsMulV1
	// tritsMul = tritsMulV2
	// tritsMul = tritsMulV3
)

//------------------------------------------------------------------------------

func trytesMul[T CoreTryte](n int, a, b T) (hi, lo T) {
	var (
		w        T
		w_lo, w_hi T
		carry    Trit
	)
	for i := 0; i < n; i++ {
		ai := getTrit(a, i)
		for j := 0; j < n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}

		w_lo = tryteShiftLeft(n, w, i)        // w << i
		w_hi = tryteShiftRight(n, w, (n - i)) // w >> (n - i)

		carry = 0
		lo, carry = trytesAdd(n, lo, w_lo, carry)
		hi, carry = trytesAdd(n, hi, w_hi, carry)
	}
	return hi, lo
}

func trytesMulLo[T CoreTryte](n int, a, b T) (lo T) {
	var (
		w     T
		w_lo   T
		carry Trit
	)
	for i := 0; i < n; i++ {
		ai := getTrit(a, i)
		for j := 0; j < n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}

		w_lo = tryteShiftLeft(n, w, i) // w << i

		carry = 0
		lo, carry = trytesAdd(n, lo, w_lo, carry)
	}
	return lo
}
