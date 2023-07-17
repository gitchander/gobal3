package overflows

import (
	"math"
)

// https://groups.google.com/g/golang-nuts/c/h5oSN5t3Au4/m/KaNQREhZh0QJ

func MulUint64(a, b uint64) (uint64, bool) {
	c := a * b
	if (a <= 1) || (b <= 1) {
		return c, true
	}
	if (c / b) != a {
		return 0, false
	}
	return c, true
}

func MulInt64(a, b int64) (int64, bool) {
	c := a * b
	if (a == 0) || (b == 0) || (a == 1) || (b == 1) {
		return c, true
	}
	if a == math.MinInt64 || b == math.MaxInt64 {
		return 0, false
	}
	if (c / b) != a {
		return 0, false
	}
	return c, true
}

func MulInt(a, b int) (int, bool) {
	c := a * b
	if (a == 0) || (b == 0) || (a == 1) || (b == 1) {
		return c, true
	}
	if a == math.MinInt || b == math.MaxInt {
		return 0, false
	}
	if (c / b) != a {
		return 0, false
	}
	return c, true
}
