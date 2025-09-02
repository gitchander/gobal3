package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/gitchander/gobal3/bal3"
	"github.com/gitchander/gobal3/bal3/base27"
	"github.com/gitchander/gobal3/utils/random"
)

func main() {
	//testToString()
	//testIncTC4()
	//testIncTC6()
	//testIncTC9()
	testIncTC27()
	//testFormatBase27()
	//testParseBase27()
	//testLimits()
	//testQuoRemT32Random()
	//testBig()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mergeSlices[T any](as ...[]T) []T {
	var bs []T
	for _, a := range as {
		bs = append(bs, a...)
	}
	return bs
}

func makeInt64s(min, max int64) []int64 {
	var (
		n  = max - min
		as = make([]int64, n)
	)
	for i := int64(0); i < n; i++ {
		as[i] = min + i
	}
	return as
}

func makeValues(min, max int64) []int64 {
	const m int64 = 27
	return mergeSlices[int64](
		makeInt64s(min, min+m),
		makeInt64s(-m, +m+1),
		makeInt64s(max-m+1, max+1),
	)
}

func testToString() {
	tc := bal3.TC6
	t, _ := tc.Int64ToTrite(-50)
	fmt.Println(t.ToInt64())
}

func testIncTC4() {
	tc := bal3.TC4
	min, max := tc.LimitsInt64()
	var a bal3.Tryte4
	a, _ = tc.Int64ToTrite(min)
	for i := min; i <= max; i++ {
		s := base27.FormatBase27(tc, a)
		ai64 := a.ToInt64()
		fmt.Printf("%+4d %4s %2s\n", ai64, a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testIncTC6() {
	tc := bal3.TC6
	min, max := tc.LimitsInt64()
	var a bal3.Tryte6
	a, _ = tc.Int64ToTrite(min)
	for i := min; i <= max; i++ {
		s := base27.FormatBase27(tc, a)
		fmt.Printf("%4d %6s %2s\n", tc.ToInt64(a), a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testIncTC9() {
	tc := bal3.TC9
	min, max := tc.LimitsInt64()
	var a bal3.Tryte9
	a, _ = tc.Int64ToTrite(min)
	for i := min; i <= max; i++ {
		s := base27.FormatBase27(tc, a)
		fmt.Printf("%5d %9s %3s\n", tc.ToInt64(a), a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testIncTC27() {
	tc := bal3.TC27
	min, max := tc.LimitsInt64()
	values := makeValues(min, max)
	for _, value := range values {
		a, _ := tc.Int64ToTrite(value)
		s := base27.FormatBase27(tc, a)
		fmt.Printf("%15d %28s %10s\n", tc.ToInt64(a), a, s)
	}
}

func testFormatBase27() {
	tc := bal3.TC32

	const (
		// min int64 = -40
		// max int64 = +40

		min int64 = 10000000
		max int64 = 10000100
	)

	for i := min; i <= max; i++ {
		a, _ := tc.Int64ToTrite(i)
		s := base27.FormatBase27(tc, a)
		fmt.Printf("%3d %5s %3s\n", i, a, s)
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

	a, err := base27.ParseBase27(tc, "4DD")
	checkError(err)
	s := base27.FormatBase27(tc, a)
	fmt.Println(tc.ToInt64(a), a, s)
}

func testLimits() {

	printLimits := func(typeName string, min, max int64) {
		w := max - min + 1
		fmt.Printf("%q: { min: %d, max: %d } %d\n", typeName, min, max, w)
	}

	var min, max int64

	min, max = bal3.TC4.LimitsInt64()
	printLimits("tryte4", min, max)

	min, max = bal3.TC6.LimitsInt64()
	printLimits("tryte6", min, max)

	min, max = bal3.TC8.LimitsInt64()
	printLimits("tryte8", min, max)

	min, max = bal3.TC9.LimitsInt64()
	printLimits("tryte9", min, max)

	min, max = bal3.TC16.LimitsInt64()
	printLimits("tryte16", min, max)

	min, max = bal3.TC32.LimitsInt64()
	printLimits("tryte32", min, max)

	fmt.Println()
	printLimits("int32", math.MinInt32, math.MaxInt32)
	printLimits("int64", math.MinInt64, math.MaxInt64)
}

func testQuoRemT16Samples() {

	tc := bal3.TC16

	intToTrite := func(v int64) bal3.Tryte16 {
		t, _ := tc.Int64ToTrite(v)
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

			av, _ = tc.TryteToInt64(a, 0)
			bv, _ = tc.TryteToInt64(b, 0)
		)

		quo, rem := tc.QuoRem(a, b)

		var (
			haveQuo = quo.ToInt64()
			haveRem = rem.ToInt64()

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
			av, _ = tc.TryteToInt64(a, 0)
			bv, _ = tc.TryteToInt64(b, 0)
		)

		quo, rem := tc.QuoRem(a, b)

		var (
			haveQuo = tc.ToInt64(quo)
			haveRem = tc.ToInt64(rem)

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
		av int64 = 40
		bv int64 = 40
	)

	tc := bal3.TC4

	a, _ := tc.Int64ToTrite(av)
	b, _ := tc.Int64ToTrite(bv)

	hi, lo := tc.Mul(a, b)
	fmt.Println(hi, lo)
	fmt.Println("1>", tc.ToInt64(hi)*81+tc.ToInt64(lo))
	fmt.Println("2>", av*bv)

	cLo := a.Mul(b)
	fmt.Println(cLo, tc.ToInt64(cLo))
}

func testTryte8() {
	tc := bal3.TC8

	x, _ := tc.Int64ToTrite(7)
	y, _ := tc.Int64ToTrite(7)

	xyLo := x.Mul(y)
	fmt.Println(xyLo, tc.ToInt64(xyLo))

	for i := int64(-3280); i <= 3280; i++ {
		a, _ := tc.Int64ToTrite(i)
		fmt.Printf("%3d: %5s\n", i, a)
	}
}

func testQuoRemT16() {

	tc := bal3.TC8

	var (
		a, _ = tc.Int64ToTrite(-21523360)
		b, _ = tc.Int64ToTrite(7)
	)

	quo, rem := tc.QuoRem(a, b)
	fmt.Println(quo, rem)
	fmt.Println(tc.ToInt64(quo), tc.ToInt64(rem))
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

func testBig() {
	var (
		//s = "0"
		s = "000T101T01TTT10T10001T01010101010101010101010TTTTTTT00000TTTT"
	)
	b, err := bal3.ParseBt(s)
	checkError(err)
	fmt.Println(b)
	b.Neg()
	fmt.Println(b)

	fmt.Println(b.IsNegative())
	fmt.Println(b.IsZero())
	fmt.Println(b.IsPositive())
}
