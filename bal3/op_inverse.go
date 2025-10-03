package bal3

// opposite
// reverse
// inverse

// https://en.wikipedia.org/wiki/Additive_inverse

func tritInverse1(a Trit) Trit {
	return -a
}

func tritInverse2(a Trit) Trit {
	return trico.Inverse(a)
}

var (
	tritInverse = tritInverse1
	// tritInverse = tritInverse2
)

// Opp, Opposite
// Inv, Inverse
// Rev, Reverse

func tryteInverse[Tryte GenericTryte](n int, a Tryte) Tryte {
	var b Tryte
	for i := 0; i < n; i++ {
		t := getTrit(a, i)
		t = tritInverse(t)
		b = setTrit(b, i, t)
	}
	return b
}
