package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/gitchander/gobal3/utils/digits"
)

func main() {
	testCalcDigits()
	//testCalcDigitsN()
	//testDigits()
}

func testCalcDigits() {
	var (
		// digiter      = digits.NewDigiter(0, 2)
		// digitWidth   = 1
		// digitsNumber = 64

		// digiter      = digits.NewDigiter(0, 10)
		// digitWidth   = 3
		// digitsNumber = 25

		// digiter      = digits.NewDigiter(-1, 2)
		// digitWidth   = 3
		// digitsNumber = 41

		// digiter      = digits.NewDigiter(5, 9)
		// digitWidth   = 1
		// digitsNumber = 40

		// digiter      = digits.NewDigiter(-4, 5)
		// digitWidth   = 3
		// digitsNumber = 21

		// digiter      = digits.NewDigiter(-40, 41)
		// digitWidth   = 4
		// digitsNumber = 11

		digiter      = digits.NewDigiter(4, 13)
		digitWidth   = 3
		digitsNumber = 25
	)
	var as []int
	as = appendIntsMinMax(as, math.MinInt, math.MinInt+15)
	as = appendIntsMinMax(as, -20, 20)
	as = appendIntsMinMax(as, math.MaxInt-15, math.MaxInt)
	as = append(as, math.MaxInt)
	ds := make([]int, digitsNumber)
	for _, a := range as {
		rest := digiter.IntToDigits(a, ds)
		b := digiter.DigitsToInt(ds, rest)
		if b != a {
			panic(fmt.Errorf("%d != %d", b, a))
		}
		fmt.Printf("%21d %21d %s\n", a, rest, formatDigits(ds, digitWidth))
	}
}

func appendIntsMinMax(as []int, min, max int) []int {
	for a := min; a < max; a++ {
		as = append(as, a)
	}
	return as
}

func testCalcDigitsN() {
	const (
		digitWidth = 3
	)
	digiter := digits.NewDigiter(0, 10)
	x := 123404534
	ds, rest := digiter.IntToDigitsN(x, 10)
	fmt.Println(x, rest, formatDigits(ds, digitWidth))
}

func testDigits() {
	const (
		digitWidth = 3
	)
	digiter := digits.NewDigiter(-1, 2)
	ds := make([]int, 10)
	for x := 0; x < 100; x++ {
		rest := digiter.IntToDigits(x, ds)
		fmt.Printf("% 4d %3d %s\n", x, rest, formatDigits(ds, digitWidth))
	}
}

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
