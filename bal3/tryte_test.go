package bal3

import (
	"fmt"
	"testing"
)

func testTritsAdd[T CoreTryte](tc TryteCore[T], a, b T, carryIn Trit) error {

	res, carryOut := tc.Add(a, b, carryIn)

	var (
		have, _ = tc.TryteToInt64(res, int64(carryOut))
		want    = tc.ToInt64(a) + tc.ToInt64(b) + int64(carryIn)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testTritsSub[T CoreTryte](tc TryteCore[T], a, b T, carryIn Trit) error {

	res, carryOut := tc.Sub(a, b, carryIn)

	var (
		have, _ = tc.TryteToInt64(res, int64(carryOut))
		want    = tc.ToInt64(a) - tc.ToInt64(b) + int64(carryIn)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testTritsAddSub[T CoreTryte](tc TryteCore[T], a, b T, carryIn Trit) error {
	err := testTritsAdd(tc, a, b, carryIn)
	if err != nil {
		return err
	}
	return testTritsSub(tc, a, b, carryIn)
}

func testAddSubRand[T CoreTryte](tc TryteCore[T]) error {
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
	min, max := tc.LimitsInt64()
	for av := min; av <= max; av++ {
		a, _ := tc.Int64ToTrite(av)
		for bv := min; bv <= max; bv++ {
			b, _ := tc.Int64ToTrite(bv)
			for _, carryIn := range tritsAll {
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

	min, max := tc.Limits()

	type sample[T CoreTryte] struct {
		a       T
		b       T
		carryIn Trit
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

	test := func(av, bv int64) {

		want := 0
		switch {
		case av < bv:
			want = -1
		case av > bv:
			want = 1
		}

		var (
			a, _ = tc.Int64ToTrite(av)
			b, _ = tc.Int64ToTrite(bv)
		)

		have := tc.Compare(a, b)

		if have != want {
			t.Fatalf("invalid value: have %d, want %d", have, want)
		}
	}

	min, max := tc.LimitsInt64()
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
