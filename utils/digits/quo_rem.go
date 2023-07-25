package digits

import (
	"math"
)

// quoRemMinMax
// a <= b

// val: ....................... | a ... b | .......................
// quo: ... |   -2    |   -1    |    0    |    1    |    2    | ...
// rem: ... | a ... b | a ... b | a ... b | a ... b | a ... b | ...

func QuoRemMinMax1(x int, min, max int) (q, r int) {

	base := max - min + 1

	if x < min {
		q, r = quoRem(x-max, base)
		r += max
	} else {
		q, r = quoRem(x-min, base)
		r += min
	}

	return q, r
}

func QuoRemMinMax2(x int, min, max int) (q, r int) {

	base := max - min + 1

	q, r = quoRem(x, base)

	for r < min {
		if q == math.MinInt {
			panic("overflow min")
		}
		q--
		r += base
	}
	for r > max {
		if q == math.MaxInt {
			panic("overflow max")
		}
		q++
		r -= base
	}

	return q, r
}
