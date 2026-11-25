package bal3

func tritSign(t Trit) int {
	if t < 0 {
		return -1
	}
	if t > 0 {
		return +1
	}
	return 0
}

func signCanonize(sign int) int {
	if sign < 0 {
		return -1
	}
	if sign > 0 {
		return +1
	}
	return 0
}

//------------------------------------------------------------------------------

//           /
//           | -1: x < 0
// sign(x) = |  0: x = 0
//           | +1: x > 0
//           \

func tryteSign[Tryte GenericTryte](n int, x Tryte) int {
	for i := range backward(n) {
		sign := tritSign(getTrit(x, i))
		if sign != 0 {
			return sign
		}
	}
	return 0
}

//------------------------------------------------------------------------------

func tryteSetSign[Tryte GenericTryte](n int, x Tryte, sign int) (Tryte, bool) {

	var (
		haveSign    int
		indexOfSign int = -1

		wantSign = signCanonize(sign)
	)

	for i := range backward(n) {
		haveSign = tritSign(getTrit(x, i))
		if haveSign != 0 {
			indexOfSign = i
			break
		}
	}

	if haveSign == 0 {
		return x, (wantSign == 0)
	}

	if wantSign == 0 {
		return tryteSetAllTrits[Tryte](n, 0), true // Zeroing always succeeds and sets trits to zero
		// return x, false
	}

	if haveSign == wantSign {
		return x, true
	}

	for i := range backward(indexOfSign + 1) {
		t := getTrit(x, i)
		t = tritInverse(t)
		x = setTrit(x, i, t)
	}
	return x, true
}
