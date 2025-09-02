package bal3

// tryteSetAllTrits

func setTritsN[Tryte CoreTryte](n int, t Trit) Tryte {
	var x Tryte
	for i := 0; i < n; i++ {
		x = setTrit(x, i, t)
	}
	return x
}
