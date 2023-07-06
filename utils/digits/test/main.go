package main

import (
	"fmt"

	"github.com/gitchander/gobal3/utils/digits"
)

func main() {

	const (
		vmin = -15
		vmax = 15
	)

	const (
		min = -1
		max = 4
	)

	var ds = make([]int, 10)
	for x := vmin; x <= vmax; x++ {
		rest := digits.CalcDigits(x, min, max, ds)
		y := digits.CalcNumber(min, max, ds, rest)
		if y != x {
			panic(fmt.Errorf("%d != %d", y, x))
		}
		reverse(ds)
		fmt.Printf("%3d %3d %v\n", x, rest, ds)
	}
}

func reverse[T any](a []T) {
	i, j := 0, (len(a) - 1)
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
}
