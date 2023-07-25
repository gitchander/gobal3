package bal3

import (
	"fmt"
	"math"
	"testing"
)

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
				a, _ = tc.Int64ToTrite(av)
				b, _ = tc.Int64ToTrite(bv)
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
		err := testQuoRemDouble(tc, p[0], p[1])
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestQuoRemT4All(t *testing.T) {
	tc := TC4
	min, max := tc.LimitsInt64()
	for ai := min; ai <= max; ai++ {
		a, _ := tc.Int64ToTrite(ai)
		for bi := min; bi <= max; bi++ {
			if bi == 0 {
				continue
			}
			b, _ := tc.Int64ToTrite(bi)
			err := testQuoRemDouble(tc, a, b)
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

		a := tc.RandSh(r)
		var b Tryte16
		for b.IsZero() {
			b = tc.RandSh(r)
		}

		err := testQuoRemT16(a, b)
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
		err := testQuoRemDouble(tc, a, b)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testQuoRemRange[T Unsigned](tc TryteCore[T]) error {
	r := newRandNext()
	for i := 0; i < 1000; i++ {
		a := tc.RandSh(r)
		var b T
		for tc.IsZero(b) {
			b = tc.RandSh(r)
		}
		err := testQuoRemDouble(tc, a, b)
		if err != nil {
			return err
		}
	}
	return nil
}

func testQuoRemT16(a, b Tryte16) error {

	var (
		av, _ = a.ToInt64()
		bv, _ = b.ToInt64()
	)

	var (
		quoRemT16 = quoRemT16_v1
		//quoRemT16 = quoRemT16_v2
	)

	quo, rem := quoRemT16(a, b)
	var (
		haveQuo, _ = quo.ToInt64()
		haveRem, _ = rem.ToInt64()
	)

	wantQuo, wantRem := quoRemInt64(av, bv)

	printAll := func() {
		fmt.Printf("have: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, haveQuo, haveRem)
		fmt.Printf("want: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, wantQuo, wantRem)
	}

	var (
		haveA = bv*haveQuo + haveRem
		wantA = av
	)
	if haveA != wantA {
		printAll()
		return fmt.Errorf("wrong (b*quo + rem) value: have %d, want %d", haveA, wantA)
	}

	if haveQuo != wantQuo {
		printAll()
		return fmt.Errorf("wrong quo: have %d, want %d", haveQuo, wantQuo)
	}

	if haveRem != wantRem {
		printAll()
		return fmt.Errorf("wrong rem: have %d, want %d", haveRem, wantRem)
	}

	return nil
}

func testQuoRemDouble[T Unsigned](tc TryteCore[T], a, b T) error {

	var (
		av = tc.tryteToInt64(a, 0)
		bv = tc.tryteToInt64(b, 0)
	)

	quo, rem := tc.QuoRem(a, b)
	var (
		haveQuo = tc.tryteToInt64(quo, 0)
		haveRem = tc.tryteToInt64(rem, 0)
	)

	wantQuo, wantRem := quoRemInt64(av, bv)

	printAll := func() {
		fmt.Printf("have: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, haveQuo, haveRem)
		fmt.Printf("want: quoRem(%d, %d) => { quo: %d, rem: %d }\n", av, bv, wantQuo, wantRem)
	}

	var (
		haveA = bv*haveQuo + haveRem
		wantA = av
	)
	if haveA != wantA {
		printAll()
		return fmt.Errorf("wrong (b*quo + rem) value: have %d, want %d", haveA, wantA)
	}

	if haveQuo != wantQuo {
		printAll()
		return fmt.Errorf("wrong quo: have %d, want %d", haveQuo, wantQuo)
	}

	if haveRem != wantRem {
		printAll()
		return fmt.Errorf("wrong rem: have %d, want %d", haveRem, wantRem)
	}

	return nil
}

func TestQuoRemBal3Samples(t *testing.T) {
	samples := []struct {
		a        int64
		quo, rem int64
	}{
		{a: math.MinInt64 + 0, quo: -3074457345618258603, rem: 1},
		{a: math.MinInt64 + 1, quo: -3074457345618258602, rem: -1},
		{a: math.MinInt64 + 2, quo: -3074457345618258602, rem: 0},
		{a: math.MinInt64 + 3, quo: -3074457345618258602, rem: 1},
		{a: math.MinInt64 + 4, quo: -3074457345618258601, rem: -1},
		{a: math.MinInt64 + 5, quo: -3074457345618258601, rem: 0},
		{a: math.MinInt64 + 6, quo: -3074457345618258601, rem: 1},

		{a: -7, quo: -2, rem: -1},
		{a: -6, quo: -2, rem: 0},
		{a: -5, quo: -2, rem: 1},
		{a: -4, quo: -1, rem: -1},
		{a: -3, quo: -1, rem: 0},
		{a: -2, quo: -1, rem: 1},
		{a: -1, quo: 0, rem: -1},
		{a: 0, quo: 0, rem: 0},
		{a: 1, quo: 0, rem: 1},
		{a: 2, quo: 1, rem: -1},
		{a: 3, quo: 1, rem: 0},
		{a: 4, quo: 1, rem: 1},
		{a: 5, quo: 2, rem: -1},
		{a: 6, quo: 2, rem: 0},
		{a: 7, quo: 2, rem: 1},

		{a: math.MaxInt64 - 5, quo: 3074457345618258601, rem: -1},
		{a: math.MaxInt64 - 4, quo: 3074457345618258601, rem: 0},
		{a: math.MaxInt64 - 3, quo: 3074457345618258601, rem: 1},
		{a: math.MaxInt64 - 2, quo: 3074457345618258602, rem: -1},
		{a: math.MaxInt64 - 1, quo: 3074457345618258602, rem: 0},
		{a: math.MaxInt64 - 0, quo: 3074457345618258602, rem: 1},
	}
	for _, sample := range samples {
		quo, rem := quoRemBal3(sample.a)
		if quo != sample.quo {
			t.Fatalf("invalid %q (a=%d): have %d, want %d", "quo", sample.a, quo, sample.quo)
		}
		if rem != sample.rem {
			t.Fatalf("invalid %q (a=%d): have %d, want %d", "rem", sample.a, rem, sample.rem)
		}
	}
}
