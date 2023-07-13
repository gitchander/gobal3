package main

import (
	"fmt"
	"math"

	"github.com/gitchander/gobal3/ternary"
	"github.com/gitchander/gobal3/utils/digits"
)

func main() {
	//testFunc()
	//testModBal3()
	testQuoRem()
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

	ts := [3][3]int{
		{1, -1, 0},
		{-1, 0, 1},
		{0, 1, -1},
	}

	// ts := [3][3]int{
	// 	{-1, 0, 0},
	// 	{0, 0, 0},
	// 	{0, 0, 1},
	// }

	f := func(a, b int) int {
		return maxInts(
			minInt3(ternary.Is(a, -1), ternary.Is(b, -1), ts[0][0]),
			minInt3(ternary.Is(a, -1), ternary.Is(b, 0), ts[0][1]),
			minInt3(ternary.Is(a, -1), ternary.Is(b, 1), ts[0][2]),

			minInt3(ternary.Is(a, 0), ternary.Is(b, -1), ts[1][0]),
			minInt3(ternary.Is(a, 0), ternary.Is(b, 0), ts[1][1]),
			minInt3(ternary.Is(a, 0), ternary.Is(b, 1), ts[1][2]),

			minInt3(ternary.Is(a, 1), ternary.Is(b, -1), ts[2][0]),
			minInt3(ternary.Is(a, 1), ternary.Is(b, 0), ts[2][1]),
			minInt3(ternary.Is(a, 1), ternary.Is(b, 1), ts[2][2]),
		)
	}

	s := ternary.PrintableBinaryTable("\t", f)
	fmt.Print(s)
}

//------------------------------------------------------------------------------

func minInt2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func maxInt3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
	} else {
		if b > c {
			return b
		}
	}
	return c
}

func minInts(as ...int) int {
	n := len(as)
	if n == 0 {
		return 0 // default value
	}
	j := 0
	for i := 1; i < n; i++ {
		if as[j] > as[i] {
			j = i
		}
	}
	return as[j]
}

func maxInts(as ...int) int {
	n := len(as)
	if n == 0 {
		return 0 // default value
	}
	j := 0
	for i := 1; i < n; i++ {
		if as[j] < as[i] {
			j = i
		}
	}
	return as[j]
}

//------------------------------------------------------------------------------

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

	const (
		min = -1
		max = +1

		base = max - min + 1
	)

	quoRemMinMax := digits.QuoRemMinMax
	_ = quoRemMinMax

	testSample := func(a int) {
		q, r := quoRemBal3(a)
		//q, r := quoRemMinMax(a, min, max)

		b := q*base + r
		if b != a {
			err := fmt.Errorf("invalid value: have %d, want %d", b, a)
			panic(err)
		}
		fmt.Printf("%3d %3d %3d\n", a, q, r)
	}

	{
		const n = 7
		a := math.MinInt
		for i := 0; i < n; i++ {
			testSample(a)
			a++
		}
	}

	for i := -7; i <= 7; i++ {
		testSample(i)
	}

	{
		const n = 6
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
