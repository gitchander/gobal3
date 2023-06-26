package bal3

import (
	"fmt"
	"testing"
)

func TestQuoRemT16Samples(t *testing.T) {

	type tritsType = Tryte16
	type sampleType [2]tritsType
	tc := T16C

	var (
		fromStr = func(as, bs string) sampleType {
			var (
				a = tc.MustParse(as)
				b = tc.MustParse(bs)
			)
			return sampleType{a, b}
		}

		fromInt = func(av, bv int) sampleType {
			var (
				a = tc.FromInt(av)
				b = tc.FromInt(bv)
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
		fromInt(2, 4),
		fromInt(5, 3),
		fromInt(-8, -5),
		fromInt(1316, -191),
		fromInt(2, 3),
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
	tc := T4C
	min, max := tc.Bounds()
	for ai := min; ai <= max; ai++ {
		a := tc.FromInt(ai)
		for bi := min; bi <= max; bi++ {
			if bi == 0 {
				continue
			}
			b := tc.FromInt(bi)
			err := testQuoRemDouble(tc, a, b)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestQuoRemT16Random(t *testing.T) {
	tc := T16C
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
	tc := T32C

	err := testQuoRemRange[tritsType](tc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuoRemT32Samples(t *testing.T) {

	type tritsType = Tryte32
	tc := T32C

	minInt, maxInt := tc.Bounds()

	var (
		min = tc.FromInt(minInt)
		max = tc.FromInt(maxInt)
	)

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
		av = a.Int()
		bv = b.Int()
	)

	var (
		quoRemT16 = quoRemT16_v1
		//quoRemT16 = quoRemT16_v2
	)

	quo, rem := quoRemT16(a, b)
	var (
		haveQuo = quo.Int()
		haveRem = rem.Int()
	)

	wantQuo, wantRem := quoRemInt(av, bv)

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
		av = tc.ToInt(a)
		bv = tc.ToInt(b)
	)

	quo, rem := tc.QuoRem(a, b)
	var (
		haveQuo = tc.ToInt(quo)
		haveRem = tc.ToInt(rem)
	)

	wantQuo, wantRem := quoRemInt(av, bv)

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
