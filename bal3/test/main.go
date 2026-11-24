package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/gitchander/gobal3/bal3"
	"github.com/gitchander/gobal3/bal3/base27/v1/base27"
	"github.com/gitchander/gobal3/utils/random"
)

func main() {
	//testToString()
	//testIncTC4()
	//testIncTC6()
	//testIncTC9()
	//testIncTC27()
	//testFormatBase27()
	//testParseBase27()
	//testLimits()
	//testQuoRemSimple()
	//testQuoRemT32Random()
	//testBig()
	//testTryteLen()
	//testMulMax()
	//exampleMulLo()
	exampleBig()
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
	t := tc.Int64ToTrite(-50)
	fmt.Println(t.ToInt64())
}

func testIncTC4() {
	tc := bal3.TC4
	min, max := tc.LimitsInt64()
	var a bal3.Tryte4
	a = tc.Int64ToTrite(min)
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
	a = tc.Int64ToTrite(min)
	for i := min; i <= max; i++ {
		s := base27.FormatBase27(tc, a)
		fmt.Printf("%4d %6s %2s\n", tc.TryteToInt64(a), a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testIncTC9() {
	tc := bal3.TC9
	min, max := tc.LimitsInt64()
	var a bal3.Tryte9
	a = tc.Int64ToTrite(min)
	for i := min; i <= max; i++ {
		s := base27.FormatBase27(tc, a)
		fmt.Printf("%5d %9s %3s\n", tc.TryteToInt64(a), a, s)
		a, _ = tc.Add(a, 0, 1) // inc 1 trit
	}
}

func testIncTC27() {
	tc := bal3.TC27
	min, max := tc.LimitsInt64()
	values := makeValues(min, max)
	for _, value := range values {
		a := tc.Int64ToTrite(value)
		s := base27.FormatBase27(tc, a)
		fmt.Printf("%15d %28s %10s\n", tc.TryteToInt64(a), a, s)
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
		a := tc.Int64ToTrite(i)
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
	fmt.Println(tc.TryteToInt64(a), a, s)
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

	min, max = bal3.TC27.LimitsInt64()
	printLimits("tryte27", min, max)

	min, max = bal3.TC32.LimitsInt64()
	printLimits("tryte32", min, max)

	fmt.Println()
	printLimits("int32", math.MinInt32, math.MaxInt32)
	printLimits("int64", math.MinInt64, math.MaxInt64)
}

func testQuoRemT16Samples() {

	tc := bal3.TC16

	intToTrite := func(v int64) bal3.Tryte16 {
		t := tc.Int64ToTrite(v)
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

			av = tc.TryteToInt64(a)
			bv = tc.TryteToInt64(b)
		)

		quo, rem := tc.QuoRem(0, a, b)

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

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func testQuoRemSimple() {
	var (
		tc = bal3.TC32

		a = tc.Int64ToTrite(8)
		b = tc.Int64ToTrite(3)
	)

	quo, rem := tc.QuoRem(0, a, b)
	fmt.Printf("(%s / %s) = (quo,rem) = bal3:(%s,%s)\n", a, b, quo, rem)
	fmt.Printf("(quo,rem) = base10:(%d,%d)\n", tc.TryteToInt64(quo), tc.TryteToInt64(rem))

}

func testQuoRemRange() {

	const (
		b = 8

		a0 = -20
		a1 = 20
	)

	for a := a0; a < a1; a++ {

		quo, rem := quoRem(a, b)

		tc := bal3.TC32

		var (
			ta = tc.Int64ToTrite(int64(a))
			tb = tc.Int64ToTrite(int64(b))
		)

		quoT, remT := tc.QuoRem(0, ta, tb)

		//fmt.Printf("%s / %s\n", ta, tb)
		//fmt.Printf("%d / %d = [ (%d,%d), (%d,%d) ]\n", a, b, quo, rem, quoT.ToInt64(), remT.ToInt64())
		fmt.Printf("%3d / %3d = [ (%3d,%3d), (%3d,%3d) ]\n", a, b, quo, rem, quoT.ToInt64(), remT.ToInt64())

		// fmt.Println(quo, quo.ToInt64())
		// fmt.Println(rem, rem.ToInt64())
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
			av = tc.TryteToInt64(a)
			bv = tc.TryteToInt64(b)
		)

		quo, rem := tc.QuoRem(0, a, b)

		var (
			haveQuo = tc.TryteToInt64(quo)
			haveRem = tc.TryteToInt64(rem)

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

	a := tc.Int64ToTrite(av)
	b := tc.Int64ToTrite(bv)

	hi, lo := tc.Mul(a, b)
	fmt.Println(hi, lo)
	fmt.Println("1>", tc.TryteToInt64(hi)*81+tc.TryteToInt64(lo))
	fmt.Println("2>", av*bv)

	cLo := a.Mul(b)
	fmt.Println(cLo, tc.TryteToInt64(cLo))
}

func testTryte8() {
	tc := bal3.TC8

	x := tc.Int64ToTrite(7)
	y := tc.Int64ToTrite(7)

	xyLo := x.Mul(y)
	fmt.Println(xyLo, tc.TryteToInt64(xyLo))

	for i := int64(-3280); i <= 3280; i++ {
		a := tc.Int64ToTrite(i)
		fmt.Printf("%3d: %5s\n", i, a)
	}
}

func testQuoRemT16() {

	tc := bal3.TC8

	var (
		a = tc.Int64ToTrite(-21523360)
		b = tc.Int64ToTrite(7)
	)

	quo, rem := tc.QuoRem(0, a, b)
	fmt.Println(quo, rem)
	fmt.Println(tc.TryteToInt64(quo), tc.TryteToInt64(rem))
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
		//s = "111111"
		s = "001T111TT1TT1T101T01TTT10T10001T01010101010101010101010TTTTTTT00000TTTT000TTTTTT0T00T"
	)
	b, err := bal3.ParseBigTryte(s)
	checkError(err)
	fmt.Println(b)
	b.Inverse()
	fmt.Println(b)

	fmt.Println(b.ToBigInt())

	fmt.Println(b.IsNegative())
	fmt.Println(b.IsZero())
	fmt.Println(b.IsPositive())
}

func testTryteLen() {

	tc := bal3.TC9

	x, err := tc.Parse("TT1_10T_11T")
	checkError(err)

	fmt.Println(x)
	fmt.Println(tc.Len(x))
}

func testTryte() {

	tc := bal3.TC9

	x, err := tc.Parse("TT1_10T_11T")
	checkError(err)

	fmt.Println(x)
	fmt.Println(tc.Len(x))
}

func testMulMax() {

	var (
		// tc = bal3.TC9
		tc = bal3.TC27
	)

	var (
		// a = tc.SetAllTrits(-1)
		// b = tc.SetAllTrits(-1)

		a = tc.SetAllTrits(1)
		b = tc.SetAllTrits(1)

		// a = tc.Int64ToTrite(1000)
		// b = tc.Int64ToTrite(11)
	)

	printTryteField("a", tc, a)
	printTryteField("b", tc, b)

	hi, lo := tc.Mul(a, b)

	printTryteField("(a * b).hi", tc, hi)
	printTryteField("(a * b).lo", tc, lo)

	quo, rem := tc.QuoRem(hi, lo, b)

	printTryteField("quo", tc, quo)
	printTryteField("rem", tc, rem)
}

func printTryteField[T bal3.GenericTryte](fieldName string, tc bal3.TryteCore[T], t T) {
	fmt.Printf("%s = (base10) %d = (bal3) %s\n", fieldName, tc.TryteToInt64(t), tc.Format(t))
}

func exampleMulLo() {
	tc := bal3.TC9
	minValue, maxValue := tc.LimitsInt64() // -9841 9841
	fmt.Println("Limits:", minValue, maxValue)

	var (
		a = tc.Int64ToTrite(3)
		b = tc.Int64ToTrite(4920)
	)

	fmt.Println(a, b)

	c, overflow := tc.MulLo(a, b)
	fmt.Println(c, overflow)
	fmt.Println(tc.TryteToInt64(c))

	hi, lo := tc.Mul(a, b)
	fmt.Println(a, b)
	fmt.Println(hi, lo)
}

func exampleBig() {
	{
		b := bal3.NewBigTryteInt64(1245)
		fmt.Println(b)
		fmt.Println()
	}

	{
		values := []int64{
			math.MinInt64 + 0,
			math.MinInt64 + 1,
			math.MinInt64 + 2,
			math.MinInt64 + 3,

			-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5,

			math.MaxInt64 - 2,
			math.MaxInt64 - 1,
			math.MaxInt64 - 0,
		}

		for _, value := range values {
			b := bal3.NewBigTryteInt64(value)
			y, err := b.ToInt64Err()
			checkError(err)

			if y != value {
				panic(fmt.Sprintf("%d != %d", y, value))
			}
			fmt.Println(b, y)
		}
		fmt.Println()
	}

	{
		var (
			a int64 = 1243454654635
			b int64 = 763423534534452

			ba = bal3.NewBigTryteInt64(a)
			bb = bal3.NewBigTryteInt64(b)
			bc = bal3.NewBigTryte()
		)
		bc.Add(ba, bb)

		fmt.Println(bc, bc.ToBigInt())

		fmt.Println(a + b)
		fmt.Println()
	}
}
