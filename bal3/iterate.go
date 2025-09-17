package bal3

//------------------------------------------------------------------------------

func forwardIterate(n int, f func(i int) bool) {
	for i := 0; i < n; i++ { // forward iterate
		if !f(i) {
			return
		}
	}
}

func backwardIterate(n int, f func(i int) bool) {
	for i := n; i > 0; { // backward iterate
		i--

		if !f(i) {
			return
		}
	}
}

//------------------------------------------------------------------------------

type TritIterateFunc = func(i int, t Trit) bool

// Forward iterate by trits.
func tryteForward[Tryte GenericTryte](n int, x Tryte, f TritIterateFunc) {
	for i := 0; i < n; i++ { // forward iterate
		t := getTrit(x, i)
		if !f(i, t) {
			return
		}
	}
}

// Backward iterate by trits.
func tryteBackward[Tryte GenericTryte](n int, x Tryte, f TritIterateFunc) {
	for i := n; i > 0; { // backward iterate
		i--

		t := getTrit(x, i)
		if !f(i, t) {
			return
		}
	}
}
