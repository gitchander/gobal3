package main

import (
	"fmt"
	"math"

	"github.com/gitchander/gobal3/ternary"
)

func main() {
	testFunc()
	//testModBal3()
	//testQuoRem()
}

func testFunc() {

	var (
		neg = ternary.Neg
		min = ternary.Min
		max = ternary.Max
		xor = ternary.Xor
		inc = ternary.Inc
		dec = ternary.Dec
	)

	_, _, _, _, _, _ = neg, min, max, xor, inc, dec

	// +---+---+---+---+
	// |   | T | 0 | 1 |
	// +---+---+---+---+
	// | T | 1 | T | 0 |
	// +---+---+---+---+
	// | 0 | T | 0 | 1 |
	// +---+---+---+---+
	// | 1 | 0 | 1 | T |
	// +---+---+---+---+

	// http://homepage.divms.uiowa.edu/%7Ejones/ternary/logic.shtml#minimization

	f := func(a, b int) int {

		return maxN(
			minInt3(ternary.Is(a, -1), ternary.Is(b, -1), 1),
			minInt3(ternary.Is(a, -1), ternary.Is(b, 0), -1),
			minInt3(ternary.Is(a, -1), ternary.Is(b, 1), 0),

			minInt3(ternary.Is(a, 0), ternary.Is(b, -1), -1),
			minInt3(ternary.Is(a, 0), ternary.Is(b, 0), 0),
			minInt3(ternary.Is(a, 0), ternary.Is(b, 1), 1),

			minInt3(ternary.Is(a, 1), ternary.Is(b, -1), 0),
			minInt3(ternary.Is(a, 1), ternary.Is(b, 0), 1),
			minInt3(ternary.Is(a, 1), ternary.Is(b, 1), -1),
		)

		// return maxN(
		// 	minInt3(ternary.Is(a, -1), ternary.Is(b, -1), -1),
		// 	minInt3(ternary.Is(a, -1), ternary.Is(b, 0), 0),
		// 	minInt3(ternary.Is(a, -1), ternary.Is(b, 1), 0),
		// 	minInt3(ternary.Is(a, 0), ternary.Is(b, -1), 0),
		// 	minInt3(ternary.Is(a, 0), ternary.Is(b, 0), 0),
		// 	minInt3(ternary.Is(a, 0), ternary.Is(b, 1), 0),
		// 	minInt3(ternary.Is(a, 1), ternary.Is(b, -1), 0),
		// 	minInt3(ternary.Is(a, 1), ternary.Is(b, 0), 0),
		// 	minInt3(ternary.Is(a, 1), ternary.Is(b, 1), 1),
		// )
	}

	s := ternary.PrintableBinaryTable("\t", f)
	fmt.Print(s)
}

// OR
func maxN(as ...int) int {
	n := len(as)
	if n == 0 {
		return 0
	}
	m := as[0]
	for i := 1; i < n; i++ {
		m = maxInt(m, as[i])
	}
	return m
}

// And
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Or
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// And
func minInt3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func testModBal3() {
	for i := -20; i <= 20; i++ {
		fmt.Printf("%3d: %3d\n", i, modBal3(i))
	}

	a := math.MinInt
	for i := 0; i < 7; i++ {
		fmt.Printf("%3d: %3d\n", a, modBal3(a))
		a++
	}

	{
		const n = 5
		a = math.MaxInt - n
		for i := 0; i <= n; i++ {
			fmt.Printf("%3d: %3d\n", a, modBal3(a))
			a++
		}
	}
}

func modBal3_v2(x int) int {
	const (
		min  = -1
		max  = +1
		base = max - min + 1
	)
	var r int
	if x < min {
		r = (x - max) % base
		r += max
	} else {
		r = (x - min) % base
		r += min
	}
	return r
}

func modBal3_v3(x int) int {
	const (
		min  = -1
		max  = +1
		base = max - min + 1
	)
	r := x % base
	if r > max {
		r -= base
	}
	if r < min {
		r += base
	}
	return r
}

var (
	//modBal3 = modBal3_v2
	modBal3 = modBal3_v3
)

func testQuoRem() {

	testSample := func(a int) {
		const base = 3
		q, r := quoRemBal3(a)
		b := q*3 + r
		if b != a {
			err := fmt.Errorf("invalid value: have %d, want %d", b, a)
			panic(err)
		}
		fmt.Printf("%3d %3d %3d\n", a, q, r)
	}

	for i := -20; i <= 20; i++ {
		testSample(i)
	}

	{
		const n = 7
		a := math.MinInt
		for i := 0; i < n; i++ {
			testSample(a)
			a++
		}
	}

	{
		const n = 7
		a := math.MaxInt - n + 1
		for i := 0; i < n; i++ {
			testSample(a)
			a++
		}
	}
}

func quoRemBal3(a int) (q, r int) {

	const (
		min = -1
		max = +1

		base = max - min + 1
	)

	q, r = quoRem(a, base)

	if r > max {
		q++
		r -= base
	}
	if r < min {
		q--
		r += base
	}

	return q, r
}

func quoRem(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
