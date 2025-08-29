package bal3

//------------------------------------------------------------------------------

func tritsCompare(a, b Trit) int {
	if a < b {
		return -1
	}
	if a > b {
		return +1
	}
	return 0 // a == b
}

//------------------------------------------------------------------------------

//            /
//            | -1: a < b
// cmp(a,b) = |  0: a = b
//            | +1: a > b
//            \

func trytesCompare[T coreTryte](n int, a, b T) int {

	for i := n; i > 0; { // backward iterate
		i--

		var (
			ta = getTrit(a, i)
			tb = getTrit(b, i)
		)
		c := tritsCompare(ta, tb)
		if c != 0 {
			return c
		}
	}
	return 0
}
