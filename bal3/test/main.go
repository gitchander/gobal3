package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/gitchander/gobal3/bal3"
	"github.com/gitchander/gobal3/utils/random"
)

func main() {

	//testToString()

	//testIncTC4()
	//testIncTC6()
	//testIncTC9()

	//testFormatBase27()
	//testParseBase27()

	//testBounds()
	testQuoRemT32Random()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testToString() {
	tc := bal3.TC6
	fmt.Println(tc.IntToTrite(5))
}

func testIncTC4() {
	tc := bal3.TC4
	min, max := tc.Bounds()
	var a bal3.Tryte4
	a, _ = tc.IntToTrite(min)
	for i := min; i <= max; i++ {
		s := bal3.FormatBase27(tc, a)
		fmt.Printf("%3d %4s %2s\n", a.Int(), a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testIncTC6() {
	tc := bal3.TC6
	min, max := tc.Bounds()
	var a bal3.Tryte6
	a, _ = tc.IntToTrite(min)
	for i := min; i <= max; i++ {
		s := bal3.FormatBase27(tc, a)
		fmt.Printf("%4d %6s %2s\n", a.Int(), a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testIncTC9() {
	tc := bal3.TC9
	min, max := tc.Bounds()
	var a bal3.Tryte9
	a, _ = tc.IntToTrite(min)
	for i := min; i <= max; i++ {
		s := bal3.FormatBase27(tc, a)
		fmt.Printf("%5d %9s %3s\n", a.Int(), a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testFormatBase27() {
	tc := bal3.TC32

	const (
		min = 10000000
		max = 10000100
	)

	for i := min; i <= max; i++ {
		a, _ := tc.IntToTrite(i)
		s := bal3.FormatBase27(tc, a)
		fmt.Println(i, a, s)
	}
}

func testParseBase27() {

	// -41  T1111 SD
	// -40   TTTT TX
	//  40   1111 1D
	//  41  1TTTT 2X

	// -3281 NDD
	// -3280 PXX
	//  3280 4DD
	//  3281 5XX

	tc := bal3.TC8

	a, err := bal3.ParseBase27(tc, "4DD")
	checkError(err)
	s := bal3.FormatBase27(tc, a)
	fmt.Println(a.Int(), a, s)
}

func testBounds() {

	type tryteType = bal3.Tryte4

	printBounds := func(typeName string, min, max tryteType) {
		fmt.Printf("%q: { min: %d, max: %d }\n", typeName, min, max)
	}
	var min, max tryteType

	min, max = bal3.TC4.Limits()
	printBounds("tryte4", min, max)

	min, max = bal3.TC8.Limits()
	printBounds("tryte8", min, max)

	min, max = bal3.TC16.Limits()
	printBounds("tryte16", min, max)

	min, max = bal3.TC32.Limits()
	printBounds("tryte32", min, max)

	min, max = bal3.TC6.Limits()
	printBounds("tryte6", min, max)

	min, max = bal3.TC9.Limits()
	printBounds("tryte9", min, max)

	printBounds("int32", math.MinInt32, math.MaxInt32)
	printBounds("int64", math.MinInt64, math.MaxInt64)
}

func testQuoRemT16Samples() {

	tc := bal3.TC16

	intToTrite := func(v int) bal3.Tryte16 {
		t, _ := tc.IntToTrite(v)
		return t
	}

	samples := [][2]bal3.Tryte16{
		{tc.MustParse("T0000T10T"), tc.MustParse("T11T1")},
		{tc.MustParse("1T01T"), tc.MustParse("11")},
		{tc.MustParse("111T"), tc.MustParse("11")},
		{intToTrite(38), intToTrite(4)},
		{intToTrite(6580), intToTrite(-47)},
		{intToTrite(392771), intToTrite(-186)},
		{intToTrite(280), intToTrite(8)},
	}

	for _, sample := range samples {
		var (
			a = sample[0]
			b = sample[1]

			av = tc.TryteToInt(a, 0)
			bv = tc.TryteToInt(b, 0)
		)

		quo, rem := tc.QuoRem(a, b)

		var (
			haveQuo = quo.Int()
			haveRem = rem.Int()

			wantQuo = av / bv
			wantRem = av % bv
		)

		fmt.Printf("have: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, haveQuo, haveRem)
		fmt.Printf("want: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, wantQuo, wantRem)
		fmt.Println()
	}
}

func testQuoRemT32Random() {

	tc := bal3.TC32

	r := random.NewRandNow()

	for i := 0; i < 100; i++ {

		a := tc.RandSh(r)
		var b bal3.Tryte32
		for b.IsZero() {
			b = tc.RandSh(r)
		}

		var (
			av = tc.TryteToInt(a, 0)
			bv = tc.TryteToInt(b, 0)
		)

		quo, rem := tc.QuoRem(a, b)

		var (
			haveQuo = quo.Int()
			haveRem = rem.Int()

			wantQuo = av / bv
			wantRem = av % bv
		)

		fmt.Printf("have: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, haveQuo, haveRem)
		fmt.Printf("want: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, wantQuo, wantRem)
		fmt.Println()
	}
}

func testTryte4() {
	var (
		av = 40
		bv = 40
	)

	tc := bal3.TC4

	a, _ := tc.IntToTrite(av)
	b, _ := tc.IntToTrite(bv)

	hi, lo := tc.Mul(a, b)
	fmt.Println(hi, lo)
	fmt.Println("1>", hi.Int()*81+lo.Int())
	fmt.Println("2>", av*bv)

	cLo := a.Mul(b)
	fmt.Println(cLo, cLo.Int())
}

func testTryte8() {
	tc := bal3.TC8

	x, _ := tc.IntToTrite(7)
	y, _ := tc.IntToTrite(7)

	xyLo := x.Mul(y)
	fmt.Println(xyLo, xyLo.Int())

	for i := -3280; i <= 3280; i++ {
		a, _ := tc.IntToTrite(i)
		fmt.Printf("%3d: %5s\n", i, a)
	}
}

func testQuoRemT16() {

	tc := bal3.TC8

	var (
		a, _ = tc.IntToTrite(-21523360)
		b, _ = tc.IntToTrite(7)
	)

	quo, rem := tc.QuoRem(a, b)
	fmt.Println(quo, rem)
	fmt.Println(quo.Int(), rem.Int())
}

func printPowersOfThree() {
	var (
		p = 1
		n = 3
	)
	for i := 0; i < 40; i++ {
		fmt.Printf("%d: %d,\n", i, p)
		p *= n
	}
}

func printPowersOfThreeBig() {
	var (
		p = big.NewInt(1)
		n = big.NewInt(3)
	)
	for i := 0; i < 100; i++ {
		fmt.Printf("%d: %d,\n", i, p)
		p.Mul(p, n)
	}
}

//
// Add table
//
// +----+----+----+
// |    |  0 |  1 |
// +----+----+----+
// |  0 |  0 |  1 |
// +----+----+----+
// |  1 |  1 | 10 |
// +----+----+----+

func bitsOpAdd(a, b int) (hi, lo int) {

	sum := a + b

	switch sum {
	case 0:
		hi, lo = 0, 0 // 00
	case 1:
		hi, lo = 0, 1 // 01
	case 2:
		hi, lo = 1, 0 // 10
	default:
		panic(fmt.Errorf("there is invalid sum %d", sum))
	}
	return
}

func bitsOpAdd3(a, b, c int) (hi, lo int) {

	var (
		hi1, lo1 = bitsOpAdd(a, b)
		hi2, lo2 = bitsOpAdd(lo1, c)
		hi3, lo3 = bitsOpAdd(hi1, hi2)
	)

	_ = hi3

	return lo3, lo2
}
