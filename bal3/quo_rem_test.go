package bal3

import (
	"fmt"
	"math"
	"testing"
)

// func testQuoRem_v1[T GenericTryte](tc TryteCore[T], hi, lo, b T) error {

// 	quo, rem := tc.QuoRem(hi, lo, b)

// 	//hi, lo := tc.Mul(quo, b)

// 	checkAll := false

// 	if checkAll {
// 		if !(tc.IsZero(hi)) {
// 			return fmt.Errorf("(%s / %s): hi is not zero: it is %s",
// 				tc.Format(a),
// 				tc.Format(b),
// 				tc.Format(hi))
// 		}
// 	}

// 	aa, c1 := tc.Add(lo, rem, 0)

// 	if checkAll {
// 		if c1 != 0 {
// 			return fmt.Errorf("c1 is not zero: it is %d", c1)
// 		}
// 	}

// 	if !tc.Equal(aa, a) {
// 		return fmt.Errorf("%s != %s", tc.Format(aa), tc.Format(a))
// 	}

// 	return nil
// }

func testQuoRem_v2[T GenericTryte](tc TryteCore[T], hi, lo, y T) error {

	dc := makeDoubleCore(tc)

	ab := makeDouble(hi, lo)

	// quo = ab / y
	// rem = ab % y
	quo, rem := tc.QuoRem(ab.Hi, ab.Lo, y)

	v1 := makeDouble(tc.Mul(quo, y)) // quo * y
	v2 := makeDouble(0, rem)
	v3, _ := dc.Add(v1, v2, 0) // quo * y + rem

	cmp := dc.Compare(ab, v3)
	if cmp != 0 {

		if true {
			dab, _ := dc.DoubleToInt64(ab, 0)
			fmt.Println("dab:", dab)

			fmt.Println("y:", tc.TryteToInt64(y))
			fmt.Println("quo:", tc.TryteToInt64(quo))
			fmt.Println("rem:", tc.TryteToInt64(rem))
		}
		return fmt.Errorf("%v != %v", ab, v3)
	}

	return nil
}

func testQuoRem[T GenericTryte](tc TryteCore[T], hi, lo, y T) error {
	// return testQuoRem_v1(tc, hi, lo, y)
	return testQuoRem_v2(tc, hi, lo, y)
}

func TestQuoRemT9Random(t *testing.T) {

	// tc := TC9
	// type Tryte = Tryte9

	tc := TC27
	type Tryte = Tryte27

	var a, b Tryte

	r := newRandNext()
	for i := 0; i < 1000; i++ {

		a = tc.randTryteForTest(r)
		b = 0
		for b.IsZero() {
			b = tc.randTryteForTest(r)
		}

		hi, lo := tc.Mul(a, b)

		err := testQuoRem(tc, hi, lo, b)
		if err != nil {
			t.Fatal(err)
		}
	}
}

//------------------------------------------------------------------------------

func TestQuoRemT9(t *testing.T) {

	tc := TC9
	type Tryte = Tryte9

	type testSample struct {
		hi, lo, y Tryte
		quo, rem  Tryte
	}
	samples := []testSample{
		{
			hi: 0, lo: tc.Int64ToTrite(7), y: tc.Int64ToTrite(5),
			quo: tc.Int64ToTrite(1), rem: tc.Int64ToTrite(2), // 7 = 5 * 1 + 2
		},
		{
			hi: 0, lo: tc.Int64ToTrite(8), y: tc.Int64ToTrite(5),
			quo: tc.Int64ToTrite(2), rem: tc.Int64ToTrite(-2), // 8 = 5 * 2 - 2
		},
		{
			hi: 0, lo: tc.Int64ToTrite(9), y: tc.Int64ToTrite(5),
			quo: tc.Int64ToTrite(2), rem: tc.Int64ToTrite(-1), // 9 = 5 * 2 - 1
		},
		{
			hi: 0, lo: tc.Int64ToTrite(10), y: tc.Int64ToTrite(5),
			quo: tc.Int64ToTrite(2), rem: tc.Int64ToTrite(0), // 10 = 5 * 2 + 0
		},
		{
			hi: tc.Int64ToTrite(-4920), lo: tc.Int64ToTrite(-4921), // 9841 * 9841 = 96845281 = 4920 * (3^9) + 4921
			y:   tc.Int64ToTrite(9841),
			quo: tc.Int64ToTrite(-9841), rem: tc.Int64ToTrite(0),
		},
		{
			hi: tc.Int64ToTrite(4920), lo: tc.Int64ToTrite(4921), // 9841 * 9841 = 96845281 = 4920 * (3^9) + 4921
			y:   tc.Int64ToTrite(9841),
			quo: tc.Int64ToTrite(9841), rem: tc.Int64ToTrite(0),
		},
	}
	for i, sample := range samples {
		//hi, lo := tc.Mul(sample.a, sample.y)
		quo, rem := tc.QuoRem(sample.hi, sample.lo, sample.y)
		if !(quo.Equal(sample.quo)) {
			t.Fatalf("sample[%d] invalid %s: have %s, want %s", i, "quo", quo, sample.quo)
		}
		if !(rem.Equal(sample.rem)) {
			t.Fatalf("sample[%d] invalid %s: have %s, want %s", i, "quo", rem, sample.rem)
		}
	}
}

func TestQuoRemT16Samples(t *testing.T) {

	type tritsType = Tryte16
	type sampleType [2]tritsType
	tc := TC16

	var (
		fromStr = func(as, bs string) sampleType {
			var (
				a = tc.MustParse(as)
				b = tc.MustParse(bs)
			)
			return sampleType{a, b}
		}

		fromInt64 = func(av, bv int64) sampleType {
			var (
				a = tc.Int64ToTrite(av)
				b = tc.Int64ToTrite(bv)
			)
			return sampleType{a, b}
		}
	)

	ps := []sampleType{
		fromStr("T0000T10T", "T11T1"),
		fromStr("1T01T", "11"),
		fromStr("111T", "11"),

		// custom samples
		fromStr("1", "1T"),
		fromStr("1T", "10"),
		fromInt64(2, 4),
		fromInt64(5, 3),
		fromInt64(-8, -5),
		fromInt64(1316, -191),
		fromInt64(2, 3),
	}

	for _, p := range ps {
		//err := testQuoRemT16(p[0], p[1])
		err := testQuoRem(tc, 0, p[0], p[1])
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestQuoRemT4All(t *testing.T) {
	tc := TC4
	min, max := tc.LimitsInt64()
	for ai := min; ai <= max; ai++ {
		a := tc.Int64ToTrite(ai)
		for bi := min; bi <= max; bi++ {
			if bi == 0 {
				continue
			}
			b := tc.Int64ToTrite(bi)
			err := testQuoRem(tc, 0, a, b)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestQuoRemT16Random(t *testing.T) {
	tc := TC16
	r := newRandNext()
	for i := 0; i < 1000; i++ {

		a := tc.randTryteForTest(r)
		var b Tryte16
		for b.IsZero() {
			b = tc.randTryteForTest(r)
		}

		err := testQuoRem(tc, 0, a, b)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestQuoRemT32Random(t *testing.T) {

	type tritsType = Tryte32
	tc := TC32

	err := testQuoRemRange[tritsType](tc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuoRemT32Samples(t *testing.T) {

	type tritsType = Tryte32
	tc := TC32

	min, max := tc.Limits()

	samples := [][2]tritsType{
		{0, max},
		{0, min},
		{min, min},
		{min, max},
		{max, min},
		{max, max},
	}

	for _, sample := range samples {
		var (
			a = sample[0]
			b = sample[1]
		)
		err := testQuoRem(tc, 0, a, b)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testQuoRemRange[T GenericTryte](tc TryteCore[T]) error {
	r := newRandNext()
	for i := 0; i < 1000; i++ {
		a := tc.randTryteForTest(r)
		var b T
		for tc.IsZero(b) {
			b = tc.randTryteForTest(r)
		}
		err := testQuoRem(tc, 0, a, b)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestQuoRemBal3Samples(t *testing.T) {
	samples := []struct {
		val      int64
		quo, rem int64
	}{
		{val: math.MinInt64 + 0, quo: -3074457345618258603, rem: 1},
		{val: math.MinInt64 + 1, quo: -3074457345618258602, rem: -1},
		{val: math.MinInt64 + 2, quo: -3074457345618258602, rem: 0},
		{val: math.MinInt64 + 3, quo: -3074457345618258602, rem: 1},
		{val: math.MinInt64 + 4, quo: -3074457345618258601, rem: -1},
		{val: math.MinInt64 + 5, quo: -3074457345618258601, rem: 0},
		{val: math.MinInt64 + 6, quo: -3074457345618258601, rem: 1},
		{val: math.MinInt64 + 7, quo: -3074457345618258600, rem: -1},
		{val: math.MinInt64 + 8, quo: -3074457345618258600, rem: 0},
		{val: math.MinInt64 + 9, quo: -3074457345618258600, rem: 1},

		{val: -7, quo: -2, rem: -1},
		{val: -6, quo: -2, rem: 0},
		{val: -5, quo: -2, rem: 1},
		{val: -4, quo: -1, rem: -1},
		{val: -3, quo: -1, rem: 0},
		{val: -2, quo: -1, rem: 1},
		{val: -1, quo: 0, rem: -1},
		{val: 0, quo: 0, rem: 0},
		{val: 1, quo: 0, rem: 1},
		{val: 2, quo: 1, rem: -1},
		{val: 3, quo: 1, rem: 0},
		{val: 4, quo: 1, rem: 1},
		{val: 5, quo: 2, rem: -1},
		{val: 6, quo: 2, rem: 0},
		{val: 7, quo: 2, rem: 1},

		{val: math.MaxInt64 - 8, quo: 3074457345618258600, rem: -1},
		{val: math.MaxInt64 - 7, quo: 3074457345618258600, rem: 0},
		{val: math.MaxInt64 - 6, quo: 3074457345618258600, rem: 1},
		{val: math.MaxInt64 - 5, quo: 3074457345618258601, rem: -1},
		{val: math.MaxInt64 - 4, quo: 3074457345618258601, rem: 0},
		{val: math.MaxInt64 - 3, quo: 3074457345618258601, rem: 1},
		{val: math.MaxInt64 - 2, quo: 3074457345618258602, rem: -1},
		{val: math.MaxInt64 - 1, quo: 3074457345618258602, rem: 0},
		{val: math.MaxInt64 - 0, quo: 3074457345618258602, rem: 1},
	}
	for _, sample := range samples {
		quo, rem := quoRemBal3(sample.val)
		if quo != sample.quo {
			t.Fatalf("invalid %q (val=%d): have %d, want %d", "quo", sample.val, quo, sample.quo)
		}
		if rem != sample.rem {
			t.Fatalf("invalid %q (val=%d): have %d, want %d", "rem", sample.val, rem, sample.rem)
		}
	}
}
