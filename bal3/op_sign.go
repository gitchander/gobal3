package bal3

//           /
//           | -1: x < 0
// sign(x) = |  0: x = 0
//           | +1: x > 0
//           \

func tryteSign[T CoreTryte](n int, x T) int {

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
