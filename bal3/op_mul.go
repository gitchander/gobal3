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

func trytesMul[T coreTryte](n int, a, b T) (hi, lo T) {
	var (
		w        T
		wLo, wHi T
		carry    Trit
	)
	for i := 0; i < n; i++ {
		ai := getTrit(a, i)
		for j := 0; j < n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}

		wLo = tryteShiftLeft(n, w, i)        // w << i
		wHi = tryteShiftRight(n, w, (n - i)) // w >> (n - i)

		carry = 0
		lo, carry = trytesAdd(n, lo, wLo, carry)
		hi, carry = trytesAdd(n, hi, wHi, carry)
	}
	return hi, lo
}

func trytesMulLo[T coreTryte](n int, a, b T) (lo T) {
	var (
		w     T
		wLo   T
		carry Trit
	)
	for i := 0; i < n; i++ {
		ai := getTrit(a, i)
		for j := 0; j < n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}

		wLo = tryteShiftLeft(n, w, i) // w << i

		carry = 0
		lo, carry = trytesAdd(n, lo, wLo, carry)
	}
	return lo
}
