package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/gitchander/gobal3/utils/digits"
)

func main() {
	//testCalcDigits()
	//testCalcDigitsN()
	//testDigits()
	//testCalcDigits2()
	testRestDigit()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testCalcDigits() {
	var (
		digiter      = digits.NewDigiter(0, 1)
		digitWidth   = 1
		digitsNumber = 64

		// digiter      = digits.NewDigiter(0, 9)
		// digitWidth   = 1
		// digitsNumber = 25

		// digiter      = digits.NewDigiter(-1, 1)
		// digitWidth   = 3
		// digitsNumber = 41

		// digiter      = digits.NewDigiter(5, 8)
		// digitWidth   = 1
		// digitsNumber = 40

		// digiter      = digits.NewDigiter(-4, 5)
		// digitWidth   = 3
		// digitsNumber = 21

		// digiter      = digits.NewDigiter(-40, 41)
		// digitWidth   = 4
		// digitsNumber = 11

		// digiter      = digits.NewDigiter(4, 13)
		// digitWidth   = 3
		// digitsNumber = 25

		// digiter      = digits.NewDigiter(17, 36)
		// digitWidth   = 3
		// digitsNumber = 20

		// digiter      = digits.NewDigiter(-36, -17)
		// digitWidth   = 4
		// digitsNumber = 20

		// digiter      = digits.NewDigiter(-1, +1)
		// digitWidth   = 3
		// digitsNumber = 43
	)
	as := makeInts1(20)
	ds := make([]int, digitsNumber)
	for _, a := range as {
		rest := digiter.IntToDigits(a, ds)
		fmt.Printf("%21d %21d %s\n", a, rest, formatDigits(ds, digitWidth))

		b, err := digiter.DigitsToInt(ds, rest)
		checkError(err)
		if b != a {
			panic(fmt.Errorf("%d != %d", b, a))
		}
	}
}

func testCalcDigitsN() {
	const (
		digitWidth = 3
	)
	digiter := digits.NewDigiter(0, 9)
	x := 123404534
	ds, rest := digiter.IntToDigitsN(x, 10)
	fmt.Println(x, rest, formatDigits(ds, digitWidth))
}

func testDigits() {
	const (
		digitWidth = 3
	)
	digiter := digits.NewDigiter(-1, +1)
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

func testCalcDigits2() {

	var (
		// digiter      = digits.NewDigiter(0, 9)
		// digitWidth   = 3
		// digitsNumber = 21

		// digiter      = digits.NewDigiter(-1, +1)
		// digitWidth   = 3
		// digitsNumber = 41

		// digiter      = digits.NewDigiter(50, 108)
		// digitWidth   = 3
		// digitsNumber = 40

		// digiter      = digits.NewDigiter(-13, -9)
		// digitWidth   = 4
		// digitsNumber = 30

		digiter      = digits.NewDigiter(2, 2)
		digitWidth   = 3
		digitsNumber = 30
	)

	ds := make([]int, digitsNumber)

	var (
		//a = -1
		a = math.MinInt
		//a = math.MinInt + 8
		//a = math.MaxInt
	)

	rest := digiter.IntToDigits(a, ds)
	fmt.Printf("%d %d %s\n", a, rest, formatDigits(ds, digitWidth))

	b, err := digiter.DigitsToInt(ds, rest)
	checkError(err)
	if b != a {
		panic(fmt.Errorf("%d != %d", b, a))
	}
}

func testRestDigit() {

	var (
		dmin, dmax = -1, 1
		//dmin, dmax = 0, 1
		//dmin, dmax = 0, 19
		//dmin, dmax = -10, -4
		//dmin, dmax = 7, 9
	)

	as := makeInts1(20)

	for _, a := range as {
		rest, digit := digits.RestDigit(a, dmin, dmax)
		fmt.Printf("%20d %20d %3d\n", a, rest, digit)
	}
}

func makeInts1(d int) []int {
	var as []int
	as = appendIntsMinMax(as, math.MinInt, math.MinInt+d)
	as = appendIntsMinMax(as, -d, d+1)
	as = appendIntsMinMax(as, math.MaxInt-d, math.MaxInt)
	as = append(as, math.MaxInt)
	return as
}

func appendIntsMinMax(as []int, min, max int) []int {
	for a := min; a < max; a++ {
		as = append(as, a)
	}
	return as
}
