package digits

type Digiter struct {
	min, max int
	base     int
}

func NewDigiter(min, max int) *Digiter {
	if min >= max {
		panic("interval is empty")
	}
	return &Digiter{
		min:  min,
		max:  max,
		base: max - min,
	}
}

func (d *Digiter) Base() int {
	return d.base
}

// QuoRemInterval
// min < max
// a = min
// b = max - 1

// val: ....................... | a ... b | .......................
// quo: ... |-2 ...-2 |-1 ...-1 | 0 ... 0 | 1 ... 1 | 2 ... 2 | ...
// rem: ... | a ... b | a ... b | a ... b | a ... b | a ... b | ...

func (d *Digiter) QuoRem(x int) (q, r int) {

	q, r = quoRem(x, d.base)

	for r < d.min {
		q--
		r += d.base
	}
	for r >= d.max {
		q++
		r -= d.base
	}

	return q, r
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

// dl - digit interval
func (d *Digiter) IntToDigits(v int, ds []int) (rest int) {
	var digit int
	for i := range ds {
		v, digit = d.QuoRem(v)
		ds[i] = digit
	}
	rest = v
	return rest
}

func (d *Digiter) DigitsToInt(ds []int, rest int) int {
	v := rest
	for i := len(ds) - 1; i >= 0; i-- {
		v = (v * d.base) + ds[i]
	}
	return v
}

func (d *Digiter) IntToDigitsN(v int, n int) (ds []int, rest int) {
	var digit int
	ds = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if (v == 0) && (len(ds) > 0) {
			break
		}
		v, digit = d.QuoRem(v)
		ds = append(ds, digit)
	}
	rest = v
	return ds, rest
}
