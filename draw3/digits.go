package draw3

import (
	"github.com/gitchander/gobal3/utils/digits"
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
