package bal3

import (
	"fmt"
	"testing"
)

func testDoubleAdd[T coreTryte](dc doubleCore[T], a, b double[T], carryIn Trit) error {

	res, carryOut := dc.Add(a, b, carryIn)

	var (
		have = dc.doubleToInt64(res, int64(carryOut))
		want = (dc.doubleToInt64(a, 0) + dc.doubleToInt64(b, 0)) + int64(carryIn)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testDoubleSub[T coreTryte](dc doubleCore[T], a, b double[T], carryIn Trit) error {

	res, carryOut := dc.Sub(a, b, carryIn)

	var (
		have = dc.doubleToInt64(res, int64(carryOut))
		want = (dc.doubleToInt64(a, 0) - dc.doubleToInt64(b, 0)) + int64(carryIn)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testDoubleAddSub[T coreTryte](dc doubleCore[T], a, b double[T], carryIn Trit) error {
	err := testDoubleAdd(dc, a, b, carryIn)
	if err != nil {
		return err
	}
	return testDoubleSub(dc, a, b, carryIn)
}

//------------------------------------------------------------------------------

func TestDoubleShl(t *testing.T) {

	var (
		tc = TC8
		dc = makeDoubleCore(tc)
	)

	r := newRandNext()
	for i := 0; i < 100; i++ {

		a := dc.Rand(r)
		offset := r.Intn(2 * tc.n)

		var (
			have = dc.Shl(a, offset)
			want = dc.wantShl(a, offset)
		)

		if dc.Compare(have, want) != 0 {
			t.Fatalf("invalid value: have %d, want %d",
				dc.doubleToInt64(have, 0), dc.doubleToInt64(want, 0))
		}
	}
}

func TestDoubleShr(t *testing.T) {

	var (
		tc = TC8
		dc = makeDoubleCore(tc)
	)

	r := newRandNext()
	for i := 0; i < 100; i++ {

		a := dc.Rand(r)
		offset := r.Intn(2 * tc.n)

		var (
			have = dc.Shr(a, offset)
			want = dc.wantShr(a, offset)
		)

		if dc.Compare(have, want) != 0 {
			t.Fatalf("invalid value: have %d, want %d",
				dc.doubleToInt64(have, 0), dc.doubleToInt64(want, 0))
		}
	}
}

func TestDoubleAddT8(t *testing.T) {

	r := newRandNext()

	var (
		tc = TC8
		dc = makeDoubleCore(tc)
	)

	for i := 0; i < 1000; i++ {
		var (
			a       = dc.RandSh(r)
			b       = dc.RandSh(r)
			carryIn = randTrit(r)
		)
		err := testDoubleAddSub(dc, a, b, carryIn)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDoubleSubT8(t *testing.T) {

	r := newRandNext()

	var (
		tc = TC8
		dc = makeDoubleCore(tc)
	)

	for i := 0; i < 1000; i++ {
		var (
			a       = dc.RandSh(r)
			b       = dc.RandSh(r)
			carryIn = randTrit(r)
		)
		err := testDoubleAddSub(dc, a, b, carryIn)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDoubleAddSubT8Samples(t *testing.T) {

	type tritsType = Tryte8

	var (
		tc = TC8
		dc = makeDoubleCore(tc)
	)

	min, max := dc.Limits()

	type sample[T coreTryte] struct {
		a double[T]
		b double[T]
	}
	samples := []sample[tritsType]{
		{a: min, b: min},
		{a: min, b: max},
		{a: max, b: min},
		{a: max, b: max},
	}
	var err error
	for _, sample := range samples {
		for _, carryIn := range tritsAll {
			err = testDoubleAddSub(dc, sample.a, sample.b, carryIn)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func testDoubleMul[T coreTryte](dc doubleCore[T], a, b double[T]) error {

	hi, lo := dc.Mul(a, b)

	var (
		have = dc.doubleToInt64(lo, dc.doubleToInt64(hi, 0))
		want = dc.doubleToInt64(a, 0) * dc.doubleToInt64(b, 0)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testDoubleMulRand[T coreTryte](tc TryteCore[T]) error {
	dc := makeDoubleCore(tc)
	r := newRandNext()
	for i := 0; i < 1000; i++ {
		var (
			a = dc.RandSh(r)
			b = dc.RandSh(r)
		)
		err := testDoubleMul(dc, a, b)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestDoubleMulTryte6(t *testing.T) {
	err := testDoubleMulRand(TC6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleMulTryte9(t *testing.T) {
	err := testDoubleMulRand(TC9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleMulT8(t *testing.T) {
	err := testDoubleMulRand(TC8)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleMulT16(t *testing.T) {
	err := testDoubleMulRand(TC16)
	if err != nil {
		t.Fatal(err)
	}
}
