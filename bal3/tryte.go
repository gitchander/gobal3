package bal3

//------------------------------------------------------------------------------

// big.BitLen(), TritLen()

// Len returns the minimum number of trits required to represent x; the result is 0 for x == 0.
func tryteLen[T GenericTryte](n int, x T) int {
	for i := n; i > 0; { // backward iterate
		i--

		t := getTrit(x, i)
		if t != 0 {
			return i + 1
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

func tryteSetAllTrits[Tryte GenericTryte](n int, t Trit) Tryte {
	var x Tryte
	for i := 0; i < n; i++ {
		x = setTrit(x, i, t)
	}
	return x
}

//------------------------------------------------------------------------------

func tryteLimits[T GenericTryte](n int) (min, max T) {
	min = tryteSetAllTrits[T](n, -1)
	max = tryteSetAllTrits[T](n, +1)
	return min, max
}

//------------------------------------------------------------------------------
