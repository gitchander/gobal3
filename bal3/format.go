package bal3

func tritToChar(t int) (c byte, ok bool) {
	switch t {
	case tv_T:
		c = tc_T
	case tv_0:
		c = tc_0
	case tv_1:
		c = tc_1
	default:
		return 0, false
	}
	return c, true
}

func charToTrit(c byte) (t int, ok bool) {
	switch c {
	case tc_T:
		t = tv_T
	case tc_0:
		t = tv_0
	case tc_1:
		t = tv_1
	default:
		return 0, false
	}
	return t, true
}

func mustTritToChar(t int) byte {
	c, ok := tritToChar(t)
	if !ok {
		panic(errInvalidTrit(t))
	}
	return c
}
