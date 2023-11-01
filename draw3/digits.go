package draw3

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

func calcSubDigits(v int) []int {
	const (
		min = -1 // digit min
		max = +1 // digit max
	)
	ds := make([]int, 4)
	CalcDigits(v, min, max, ds)
	return ds
}
