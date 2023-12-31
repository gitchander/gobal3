package bal3

import (
	"fmt"
	"testing"
)

func testMul[T Unsigned](tc TryteCore[T], a, b T) error {

	hi, lo := tc.Mul(a, b)

	var (
		have, _ = tc.TryteToInt64(lo, tc.ToInt64(hi))
		want    = tc.ToInt64(a) * tc.ToInt64(b)
	)

	if have != want {
		return fmt.Errorf("invalid value: have %d, want %d", have, want)
	}
	return nil
}

func testMulBounds[T Unsigned](tc TryteCore[T]) error {
	min, max := tc.LimitsInt64()
	for av := min; av <= max; av++ {
		for bv := min; bv <= max; bv++ {
			var (
				a, _ = tc.Int64ToTrite(av)
				b, _ = tc.Int64ToTrite(bv)
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
	err := testMulBounds(TC4)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT8Rand(t *testing.T) {
	err := testMulRand(TC8)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT16Rand(t *testing.T) {
	err := testMulRand(TC16)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT32Rand(t *testing.T) {
	err := testMulRand(TC32)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT6Rand(t *testing.T) {
	err := testMulRand(TC6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMulT9Rand(t *testing.T) {
	err := testMulRand(TC9)
	if err != nil {
		t.Fatal(err)
	}
}
