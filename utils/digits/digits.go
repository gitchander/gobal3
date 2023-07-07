package digits

// QuoRemMinMax
// (a <= b)
// Z:   ....................... | a ... b | .......................
// quo: ... |-2 ...-2 |-1 ...-1 | 0 ... 0 | 1 ... 1 | 2 ... 2 | ...
// rem: ... | a ... b | a ... b | a ... b | a ... b | a ... b | ...

func QuoRemMinMax(x int, min, max int) (q, r int) {
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

func CalcDigitsN(v int, min, max int, n int) (ds []int, rest int) {
	var d int
	ds = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if (v == 0) && (len(ds) > 0) {
			break
		}
		v, d = QuoRemMinMax(v, min, max)
		ds = append(ds, d)
	}
	rest = v
	return ds, rest
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
