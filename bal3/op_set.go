package bal3

// tryteSetAllTrits

func setTritsN[T CoreTryte](n int, t Trit) T {
	var x T
	for i := 0; i < n; i++ {
		x = setTrit(x, i, t)
	}
	return x
}
