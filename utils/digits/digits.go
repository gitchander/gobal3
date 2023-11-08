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
	if min > max {
		panic("interval is empty")
	}
	return &Digiter{
		min:  min,
		max:  max,
		base: (max - min + 1),
	}
}

func (d *Digiter) checkDigit(digit int) error {
	if (d.min <= digit) && (digit <= d.max) {
		return nil
	}
	return fmt.Errorf("invalid digit %d, want interval [%d .. %d]", digit, d.min, d.max)
}

func (d *Digiter) Base() int {
	return d.base
}

func (d *Digiter) RestDigit(x int) (rest, digit int) {

	q, r := quoRem(x, d.base)

	for r < d.min {
		q--
		r += d.base
	}
	for r > d.max {
		q++
		r -= d.base
	}

	rest = q
	digit = r

	return
}

// dl - digit interval
func (d *Digiter) IntToDigits(v int, ds []int) (rest int) {
	var digit int
	for i := range ds {
		v, digit = d.RestDigit(v)
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
		v, digit = d.RestDigit(v)
		ds = append(ds, digit)
	}
	rest = v
	return ds, rest
}

func (d *Digiter) DigitsToInt(digits []int, rest int) (int, error) {
	v := d.digitsToIntV1(digits, rest)
	return v, nil
	//return d.digitsToIntV2(digits, rest)
}

func (d *Digiter) digitsToIntV1(ds []int, rest int) int {
	v := rest
	for i := len(ds) - 1; i >= 0; i-- {
		v = (v * d.base) + ds[i]
	}
	return v
}

func (d *Digiter) digitsToIntV2(digits []int, rest int) (int, error) {
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
			fmt.Println("digits:", digits, len(digits))
			return 0, fmt.Errorf("mul overflow: %d * %d", v, k)
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
