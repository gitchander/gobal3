package bal3

type Trit int

func tritsCompare(a, b Trit) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

var tritValues = [...]Trit{
	tv_T,
	tv_0,
	tv_1,
}
