package bal3

//------------------------------------------------------------------------------

// big.BitLen(), TritLen()

// Len returns the minimum number of trits required to represent x; the result is 0 for x == 0.
func tryteLen[T GenericTryte](n int, x T) int {
	for i := n; i > 0; { // backward iterate
		i--

		t := getTrit(x, i)
		if t != 0 {
			return i
		}
	}
	return 0
}

//------------------------------------------------------------------------------

func tryteDoUnary[T GenericTryte](n int, a T, f UnaryFunc) T {
	var b T
	for i := 0; i < n; i++ {
		var (
			ai = getTrit(a, i)
			bi = f(ai)
		)
		b = setTrit(b, i, bi)
	}
	return b
}

func tryteDoBinary[T GenericTryte](n int, a, b T, f BinaryFunc) T {
	var c T
	for i := 0; i < n; i++ {
		var (
			ai = getTrit(a, i)
			bi = getTrit(b, i)
			ci = f(ai, bi)
		)
		c = setTrit(c, i, ci)
	}
	return c
}

//------------------------------------------------------------------------------

//           /
//           | -1: x < 0
// sign(x) = |  0: x = 0
//           | +1: x > 0
//           \

func tryteSign[Tryte GenericTryte](n int, x Tryte) int {

	for i := n; i > 0; { // backward iterate
		i--

		t := getTrit(x, i)
		switch {
		case t < 0:
			return -1
		case t > 0:
			return +1
		}
	}
	return 0
}

//------------------------------------------------------------------------------

func tryteSetAllTrits[Tryte GenericTryte](n int, t Trit) Tryte {
	var x Tryte
	for i := 0; i < n; i++ {
		x = setTrit(x, i, t)
	}
	return x
}

//------------------------------------------------------------------------------

func tryteLimits[T GenericTryte](n int) (min, max T) {
	min = tryteSetAllTrits[T](n, tv_T)
	max = tryteSetAllTrits[T](n, tv_1)
	return min, max
}

//------------------------------------------------------------------------------
