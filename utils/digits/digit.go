package digits

import (
	"fmt"
	"math"
)

// quoRemMinMax
// a <= b

// val: ....................... | a ... b | .......................
// rem: ... | a ... b | a ... b | a ... b | a ... b | a ... b | ...
// quo: ... |   -2    |   -1    |    0    |    1    |    2    | ...

//------------------------------------------------------------------------------

// a = -1
// b = +1
// base = (b - a) + 1 = (1 - 1) + 1 = 3

//     __
// ...|10|-9|-8|-7|-6|-5|-4|-3|-2|-1| 0| 1| 2| 3| 4| 5| 6| 7| 8| 9|10|...
// ...+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+...
// ...|-1  0  1|-1  0  1|-1  0  1|-1  0  1|-1  0  1|-1  0  1|-1  0  1|... <- digits
// ...|   -3   |   -2   |   -1   |    0   |    1   |    2   |    3   |... <- rest

//------------------------------------------------------------------------------

func calcDigit1(x int, min, max int) (digit, rest int) {

	checkBaseRange(min, max)

	base := max - min + 1

	var q, r int

	if x < min {
		q, r = quoRem(x-max, base)
		r += max
	} else {
		q, r = quoRem(x-min, base)
		r += min
	}

	digit = r
	rest = q
	return
}

func calcDigit2(x int, min, max int) (digit, rest int) {

	checkBaseRange(min, max)

	base := max - min + 1

	rest, digit = quoRem(x, base)

	for digit < min {
		if rest == math.MinInt {
			panic("overflow min")
		}
		rest--
		digit += base
	}
	for digit > max {
		if rest == math.MaxInt {
			panic("overflow max")
		}
		rest++
		digit -= base
	}

	return
}

func Digit(x int, min, max int) (digit, rest int) {
	//return calcDigit1(x, min, max)
	return calcDigit2(x, min, max)
}

func checkBaseRange(min, max int) {
	if min > max {
		panic(fmt.Errorf("invalid base range [%d..%d]", min, max))
	}
}
