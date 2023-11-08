package main

import (
	"fmt"
	"math"

	. "github.com/gitchander/gobal3/ternary"
)

func main() {
	testFunc()
	//testModBal3()
	testQuoRem()
}

func testFunc() {

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

	ts := [3][3]Tri{
		{1, -1, 0},
		{-1, 0, 1},
		{0, 1, -1},
	}

	// ts := [3][3]Tri{
	// 	{-1, 0, 0},
	// 	{0, 0, 0},
	// 	{0, 0, 1},
	// }

	f := func(a, b Tri) Tri {
		return MaxN(
			Min3(Is(a, -1), Is(b, -1), ts[0][0]),
			Min3(Is(a, -1), Is(b, 0), ts[0][1]),
			Min3(Is(a, -1), Is(b, 1), ts[0][2]),

			Min3(Is(a, 0), Is(b, -1), ts[1][0]),
			Min3(Is(a, 0), Is(b, 0), ts[1][1]),
			Min3(Is(a, 0), Is(b, 1), ts[1][2]),

			Min3(Is(a, 1), Is(b, -1), ts[2][0]),
			Min3(Is(a, 1), Is(b, 0), ts[2][1]),
			Min3(Is(a, 1), Is(b, 1), ts[2][2]),
		)
	}

	s := PrintableBinaryTable("\t", f)
	fmt.Print(s)
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

	base := 3

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
