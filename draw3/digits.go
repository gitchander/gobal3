package draw3

import (
	"github.com/gitchander/gobal3/utils/digits"
	"github.com/gitchander/gobal3/utils/random"
)

func calcSubDigits(v int) []int {
	const (
		min = -1 // digit min
		max = +1 // digit max
	)
	ds := make([]int, 4)
	var d int
	for i := range ds {
		v, d = digits.RestDigit(v, min, max)
		ds[i] = d
	}
	return ds
}

func RandTrit(r *random.Rand) int {
	return r.Intn(3) - 1 // [-1, 0, +1]
}

func RandDigits(r *random.Rand, n int) []int {
	digits := make([]int, n)
	for i := range digits {
		digits[i] = RandTrit(r)
	}
	return digits
}
