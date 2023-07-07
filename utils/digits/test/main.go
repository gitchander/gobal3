package main

import (
	"fmt"
	"strings"

	"github.com/gitchander/gobal3/utils/digits"
)

func main() {
	testCalcDigits()
	//testCalcDigitsN()
	//testDigits()
}

func testCalcDigits() {
	const (
		// min, max   = 0, 1
		// digitWidth = 1

		// min, max   = -1, 1
		// digitWidth = 3

		// min, max   = 0, 9
		// digitWidth = 1

		// min, max   = 5, 7
		// digitWidth = 1

		min, max   = -1, 5
		digitWidth = 3
	)
	ds := make([]int, 10)
	for x := -20; x <= 20; x++ {
		rest := digits.CalcDigits(x, min, max, ds)
		fmt.Printf("%4d %4d %s\n", x, rest, formatDigits(ds, digitWidth))
	}
}

func testCalcDigitsN() {
	const (
		min, max   = 0, 9
		digitWidth = 3
	)
	x := 123404534
	ds, rest := digits.CalcDigitsN(x, min, max, 10)
	fmt.Println(x, rest, formatDigits(ds, digitWidth))
}

func testDigits() {
	const (
		min, max   = -1, 1
		digitWidth = 3
	)
	ds := make([]int, 10)
	for x := 0; x < 100; x++ {
		rest := digits.CalcDigits(x, min, max, ds)
		fmt.Printf("% 4d %3d %s\n", x, rest, formatDigits(ds, digitWidth))
	}
}

// func reverse[T any](a []T) {
// 	i, j := 0, (len(a) - 1)
// 	for i < j {
// 		a[i], a[j] = a[j], a[i]
// 		i, j = i+1, j-1
// 	}
// }

func formatDigits(ds []int, digitWidth int) string {
	var b strings.Builder
	for i := len(ds); i > 0; i-- {
		digit := ds[i-1]
		fmt.Fprintf(&b, "%[1]*[2]d", digitWidth, digit)
	}
	return frameSquare(b.String())
}

func frameSquare(s string) string {
	return "[" + s + "]"
}

func frame(s string) string {
	return "(" + s + ")"
	//return "[" + s + "]"
}
