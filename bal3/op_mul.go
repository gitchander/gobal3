package bal3

// mul - multiplication

//------------------------------------------------------------------------------

// Mul table

// +---+---+---+---+
// | * | T | 0 | 1 |
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

// a * b = (hi << n) + lo

func trytesMul_v1[Tryte GenericTryte](n int, a, b Tryte) (hi, lo Tryte) {
	var (
		w      Tryte
		w1, w2 Tryte
		carry  Trit
	)
	for i := 0; i < n; i++ {
		ai := getTrit(a, i)
		for j := 0; j < n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}

		w1 = tryteShl(n, w, i)       // w << i
		w2 = tryteShr(n, w, (n - i)) // w >> (n - i)

		carry = 0
		lo, carry = trytesAdd(n, lo, w1, carry)
		hi, carry = trytesAdd(n, hi, w2, carry)
	}
	return hi, lo
}

func trytesMul_v2[Tryte GenericTryte](n int, a, b Tryte) (hi, lo Tryte) {
	var (
		w      Tryte
		w1, w2 Tryte
		carry  Trit
	)
	for i := 0; i < n; i++ {

		ai := getTrit(a, i)

		switch ai {
		case -1:
			w = tryteInverse(n, b)
		case 0:
			w = 0
		case +1:
			w = b
		}

		if w != 0 {
			w1 = tryteShl(n, w, i)       // w << i
			w2 = tryteShr(n, w, (n - i)) // w >> (n - i)

			carry = 0
			lo, carry = trytesAdd(n, lo, w1, carry)
			hi, carry = trytesAdd(n, hi, w2, carry)
		}
	}
	return hi, lo
}

func trytesMul[Tryte GenericTryte](n int, a, b Tryte) (hi, lo Tryte) {
	// return trytesMul_v1(n, a, b)
	return trytesMul_v2(n, a, b)
}

//------------------------------------------------------------------------------

func trytesMulLo_v1[Tryte GenericTryte](n int, a, b Tryte) (lo Tryte) {
	var (
		w  Tryte
		w1 Tryte
	)
	for i := 0; i < n; i++ {
		ai := getTrit(a, i)
		for j := 0; j < n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}

		w1 = tryteShl(n, w, i) // w << i

		lo, _ = trytesAdd(n, lo, w1, 0)
	}
	return lo
}

func trytesMulLo_v2[Tryte GenericTryte](n int, a, b Tryte) (lo Tryte) {
	var (
		w  Tryte
		w1 Tryte
	)
	for i := 0; i < n; i++ {

		ai := getTrit(a, i)

		switch ai {
		case -1:
			w = tryteInverse(n, b)
		case 0:
			w = 0
		case +1:
			w = b
		}

		if w != 0 {
			w1 = tryteShl(n, w, i) // w << i

			lo, _ = trytesAdd(n, lo, w1, 0)
		}
	}
	return lo
}

// trytesMulLo returns only lo
func trytesMulLo[Tryte GenericTryte](n int, a, b Tryte) (lo Tryte) {
	// return trytesMulLo_v1(n, a, b)
	return trytesMulLo_v2(n, a, b)
}

//------------------------------------------------------------------------------
