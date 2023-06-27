package bal3

import (
	"fmt"
	"testing"
)

func testDoubleAdd[T Unsigned](dc doubleCore[T], a, b double[T], carryIn int) error {

	hiFactor := powersOfThree[2*dc.tc.n]

	res, carryOut := dc.Add(a, b, carryIn)

	var (
		have = dc.ToInt(res) + (carryOut * hiFactor)
		want = (dc.ToInt(a) + dc.ToInt(b)) + carryIn
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testDoubleSub[T Unsigned](dc doubleCore[T], a, b double[T], carryIn int) error {

	hiFactor := powersOfThree[2*dc.tc.n]

	res, carryOut := dc.Sub(a, b, carryIn)

	var (
		have = dc.ToInt(res) + (carryOut * hiFactor)
		want = (dc.ToInt(a) - dc.ToInt(b)) + carryIn
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testDoubleAddSub[T Unsigned](dc doubleCore[T], a, b double[T], carryIn int) error {
	err := testDoubleAdd(dc, a, b, carryIn)
	if err != nil {
		return err
	}
	return testDoubleSub(dc, a, b, carryIn)
}

//------------------------------------------------------------------------------

func TestDoubleShl(t *testing.T) {

	var (
		tc = T8C
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
			t.Fatalf("invalid value: have %d, want %d", dc.ToInt(have), dc.ToInt(want))
		}
	}
}

func TestDoubleShr(t *testing.T) {

	var (
		tc = T8C
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
			t.Fatalf("invalid value: have %d, want %d", dc.ToInt(have), dc.ToInt(want))
		}
	}
}

func TestDoubleAddT8(t *testing.T) {

	r := newRandNext()

	var (
		tc = T8C
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
		tc = T8C
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
		tc = T8C
		dc = makeDoubleCore(tc)
	)

	minInt, maxInt := dc.Bounds()
	var (
		min = dc.FromInt(minInt)
		max = dc.FromInt(maxInt)
	)

	type sample[T Unsigned] struct {
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
		for _, carryIn := range tritValues {
			err = testDoubleAddSub(dc, sample.a, sample.b, carryIn)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func testDoubleMul[T Unsigned](dc doubleCore[T], a, b double[T]) error {

	hi, lo := dc.Mul(a, b)

	hiFactor := powersOfThree[2*dc.tc.n]

	var (
		have = dc.ToInt(hi)*hiFactor + dc.ToInt(lo)
		want = dc.ToInt(a) * dc.ToInt(b)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testDoubleMulRand[T Unsigned](tc TryteCore[T]) error {
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
	err := testDoubleMulRand(T6C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleMulTryte9(t *testing.T) {
	err := testDoubleMulRand(T9C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleMulT8(t *testing.T) {
	err := testDoubleMulRand(T8C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoubleMulT16(t *testing.T) {
	err := testDoubleMulRand(T16C)
	if err != nil {
		t.Fatal(err)
	}
}
