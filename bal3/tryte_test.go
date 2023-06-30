package bal3

import (
	"fmt"
	"testing"
)

func testTritsAdd[T Unsigned](tc TryteCore[T], a, b T, carryIn int) error {

	carryFactor := powersOfThree[tc.n]

	res, carryOut := tc.Add(a, b, carryIn)

	var (
		have = tc.ToInt(res) + (carryOut * carryFactor)
		want = tc.ToInt(a) + tc.ToInt(b) + carryIn
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testTritsSub[T Unsigned](tc TryteCore[T], a, b T, carryIn int) error {

	carryFactor := powersOfThree[tc.n]

	res, carryOut := tc.Sub(a, b, carryIn)

	var (
		have = tc.ToInt(res) + (carryOut * carryFactor)
		want = tc.ToInt(a) - tc.ToInt(b) + carryIn
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testTritsAddSub[T Unsigned](tc TryteCore[T], a, b T, carryIn int) error {
	err := testTritsAdd(tc, a, b, carryIn)
	if err != nil {
		return err
	}
	return testTritsSub(tc, a, b, carryIn)
}

func testAddSubRand[T Unsigned](tc TryteCore[T]) error {
	r := newRandNext()
	for i := 0; i < 1000; i++ {
		var (
			a       = tc.RandSh(r)
			b       = tc.RandSh(r)
			carryIn = randTrit(r)
		)
		err := testTritsAddSub(tc, a, b, carryIn)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestAddTryte4(t *testing.T) {
	tc := TC4
	min, max := tc.Bounds()
	for av := min; av <= max; av++ {
		a := tc.FromInt(av)
		for bv := min; bv <= max; bv++ {
			b := tc.FromInt(bv)
			for _, carryIn := range tritValues {
				err := testTritsAddSub(tc, a, b, carryIn)
				if err != nil {
					t.Fatal(err)
				}
			}
		}
	}
}

func TestAddT6Rand(t *testing.T) {
	err := testAddSubRand(TC6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddT9Rand(t *testing.T) {
	err := testAddSubRand(TC9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddTryte8(t *testing.T) {
	err := testAddSubRand(TC8)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddT8Samples(t *testing.T) {

	type tritsType = Tryte8
	tc := TC8

	minInt, maxInt := tc.Bounds()
	var (
		min = tc.FromInt(minInt)
		max = tc.FromInt(maxInt)
	)

	type sample[T Unsigned] struct {
		a       T
		b       T
		carryIn int
	}
	samples := []sample[tritsType]{
		{a: min, b: min, carryIn: -1},
		{a: min, b: min, carryIn: +1},
		{a: min, b: max, carryIn: -1},
		{a: min, b: max, carryIn: +1},
		{a: max, b: min, carryIn: -1},
		{a: max, b: min, carryIn: +1},
		{a: max, b: max, carryIn: -1},
		{a: max, b: max, carryIn: +1},
	}

	var err error
	for _, sample := range samples {
		err = testTritsAddSub(tc, sample.a, sample.b, sample.carryIn)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSubTryte8(t *testing.T) {
	tc := TC8
	r := newRandNext()
	for i := 0; i < 1000; i++ {
		var (
			a       = tc.RandSh(r)
			b       = tc.RandSh(r)
			carryIn = randTrit(r)
		)
		err := testTritsAddSub(tc, a, b, carryIn)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCompareTryte4(t *testing.T) {

	tc := TC4

	test := func(av, bv int) {

		want := 0
		switch {
		case av < bv:
			want = -1
		case av > bv:
			want = 1
		}

		var (
			a = tc.FromInt(av)
			b = tc.FromInt(bv)
		)

		have := tc.Compare(a, b)

		if have != want {
			t.Fatalf("invalid value: have %d, want %d", have, want)
		}
	}

	min, max := tc.Bounds()
	for av := min; av <= max; av++ {
		for bv := min; bv <= max; bv++ {
			test(av, bv)
		}
	}
}

func TestShiftT6(t *testing.T) {
	tc := TC6
	a := tc.SetAllTrits(1) // all trits set 1.
	b := a.Shl(1)          // b = a << 1
	b = b.Shr(1)           // b = b >> 1
	// last trit musk be 0.
	if a.Equal(b) {
		t.Fatalf("shift error: last trit is not reset: %s = %s", a, b)
	}
}
