package bal3

// neg - negative

func tritNeg1(a Trit) Trit {
	return -a
}

func tritNeg2(a Trit) Trit {
	return trico.Neg(a)
}

var (
	tritNeg = tritNeg1
	// tritNeg = tritNeg2
)

func tryteNeg[T CoreTryte](n int, a T) T {
	var b T
	for i := 0; i < n; i++ {
		t := getTrit(a, i)
		t = tritNeg(t)
		b = setTrit(b, i, t)
	}
	return b
}
