package digits

// QuoRemMinMax
// (a <= b)
// Z:   ....................... | a ... b | .......................
// quo: ... |-2 ...-2 |-1 ...-1 | 0 ... 0 | 1 ... 1 | 2 ... 2 | ...
// rem: ... | a ... b | a ... b | a ... b | a ... b | a ... b | ...

func QuoRemMinMax(x int, min, max int) (q, r int) {

	if min > max {
		panic("min > max")
	}

	// radix
	base := max - min + 1

	if (min <= x) && (x <= max) {
		q, r = 0, x
	}

	if x < min {
		q, r = quoRem(x-max, base)
		r += max
	}

	if x > max {
		q, r = quoRem(x-min, base)
		r += min
	}

	return q, r
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func CalcDigits(v int, min, max int, ds []int) (rest int) {
	var d int
	for i := range ds {
		v, d = QuoRemMinMax(v, min, max)
		ds[i] = d
	}
	rest = v
	return rest
}

func CalcNumber(min, max int, ds []int, rest int) int {
	base := max - min + 1
	var v int
	p := 1
	for _, d := range ds {
		v += d * p
		p *= base
	}
	v += rest * p
	return v
}
