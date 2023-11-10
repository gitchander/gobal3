package digits

import (
	"fmt"
	"math"
)

// RestDigit
// a <= b

// value: ....................... | a ... b | .......................
// rest:  ... |   -2    |   -1    |    0    |    1    |    2    | ...
// digit: ... | a ... b | a ... b | a ... b | a ... b | a ... b | ...

// min <= max
// base = max - min + 1
// x = rest * base + digit

func calcRestDigit1(x int, min, max int) (rest, digit int) {

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

	rest = q
	digit = r

	return
}

func calcRestDigit1_mod(x int, min, max int) (rest, digit int) {

	checkBaseRange(min, max)

	base := max - min + 1

	var q, r int

	switch {
	case x < min:
		dx := min - 1
		q, r = quoRem(x-dx, base)
		q--
		r += dx + base
	case x > max:
		dx := max + 1
		q, r = quoRem(x-dx, base)
		q++
		r += dx - base
	default:
		q = 0
		r = x
	}

	rest = q
	digit = r

	return
}

func calcRestDigit2(x int, min, max int) (rest, digit int) {

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

func calcRestDigit3(x int, min, max int) (rest, digit int) {

	checkBaseRange(min, max)

	base := max - min + 1

	rest, digit = quoRem(x, base)

	if digit < min {
		k := ceilDiv(min-digit, base)
		rest -= k
		digit += base * k
	}
	if digit > max {
		k := ceilDiv(digit-max, base)
		rest += k
		digit -= base * k
	}

	return
}

func RestDigit(x int, min, max int) (rest, digit int) {
	//return calcRestDigit1(x, min, max)
	//return calcRestDigit1_mod(x, min, max)
	//return calcRestDigit2(x, min, max)
	return calcRestDigit3(x, min, max)
}

func checkBaseRange(min, max int) {
	if min > max {
		panic(fmt.Errorf("invalid base range [%d..%d]", min, max))
	}
}
