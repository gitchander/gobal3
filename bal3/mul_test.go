package bal3

import (
	"fmt"
	"testing"
)

func testMul[T Unsigned](tc TryteCore[T], a, b T) error {

	hi, lo := tc.Mul(a, b)

	hiFactor := powersOfThree[tc.n]

	var (
		have = tc.ToInt(hi)*hiFactor + tc.ToInt(lo)
		want = tc.ToInt(a) * tc.ToInt(b)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testMulBounds[T Unsigned](tc TryteCore[T]) error {
	var min, max = tc.Bounds()
	for av := min; av <= max; av++ {
		for bv := min; bv <= max; bv++ {
			var (
				a = tc.FromInt(av)
				b = tc.FromInt(bv)
			)
			err := testMul(tc, a, b)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func testMulRand[T Unsigned](tc TryteCore[T]) error {
	r := newRandNext()
	for i := 0; i < 1000; i++ {
		var (
			a = tc.RandSh(r)
			b = tc.RandSh(r)
		)
		err := testMul(tc, a, b)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestMulT4Bounds(t *testing.T) {
	err := testMulBounds(T4C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT8Rand(t *testing.T) {
	err := testMulRand(T8C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT16Rand(t *testing.T) {
	err := testMulRand(T16C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT32Rand(t *testing.T) {
	err := testMulRand(T32C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT6Rand(t *testing.T) {
	err := testMulRand(T6C)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT9Rand(t *testing.T) {
	err := testMulRand(T9C)
	if err != nil {
		t.Fatal(err)
	}
}
