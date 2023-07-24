package digits

import (
	"fmt"

	"github.com/gitchander/gobal3/utils/overflows"
)

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

func (d *Digiter) checkDigit(digit int) error {
	if (d.min <= digit) && (digit < d.max) {
		return nil
	}
	return fmt.Errorf("invalid digit %d, want interval [%d...%d)", digit, d.min, d.max)
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

func (d *Digiter) Digit(x int) (digit, rest int) {

	q, r := quoRem(x, d.base)

	for r < d.min {
		q--
		r += d.base
	}
	for r >= d.max {
		q++
		r -= d.base
	}

	digit = r
	rest = q
	return
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
		digit, v = d.Digit(v)
		ds[i] = digit
	}
	rest = v
	return rest
}

func (d *Digiter) IntToDigitsN(v int, n int) (ds []int, rest int) {
	var digit int
	ds = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if (v == 0) && (len(ds) > 0) {
			break
		}
		digit, v = d.Digit(v)
		ds = append(ds, digit)
	}
	rest = v
	return ds, rest
}

func (d *Digiter) digitsToInt_(ds []int, rest int) int {
	v := rest
	for i := len(ds) - 1; i >= 0; i-- {
		v = (v * d.base) + ds[i]
	}
	return v
}

func (d *Digiter) DigitsToInt(digits []int, rest int) (int, error) {
	base := d.base
	v := rest
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		err := d.checkDigit(digit)
		if err != nil {
			return 0, err
		}

		//fmt.Println(v)

		// v = (v * d.base) + digit
		// v = (v * d.base) + digit + (k*d.base - k*d.base)
		// v = (v + k)*d.base - k*d.base + digit

		k := 0
		switch {
		case v < 0:
			k = +base
		case v > 0:
			k = -base
		}

		vb, ok := overflows.MulInt((v + k), base)
		if !ok {
			return 0, fmt.Errorf("mul overflow")
		}
		s1, ok := overflows.AddInt(-(k * base), digit)
		if !ok {
			return 0, fmt.Errorf("add overflow: s1")
		}
		s2, ok := overflows.AddInt(vb, s1)
		if !ok {
			return 0, fmt.Errorf("add overflow: s2")
		}
		v = s2
	}
	return v, nil
}
