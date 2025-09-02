package bal3

import (
	"fmt"
	"testing"
)

func testMul[T CoreTryte](tc TryteCore[T], a, b T) error {

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

func testMulBounds[T CoreTryte](tc TryteCore[T]) error {
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

func testMulRand[T CoreTryte](tc TryteCore[T]) error {
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

func TestMulT27Rand(t *testing.T) {
	err := testMulRand(TC27)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrytesMul(t *testing.T) {
	samples := [][3]int64{
		{0, 0, 0},
		{0, 1, 0},
		{1, 3, 3},
		{5, 7, 35},
		{423542, 223424, 94629447808},
	}
	tc := TC9
	for i, v := range samples {
		var (
			a, _ = tc.Int64ToTrite(v[0])
			b, _ = tc.Int64ToTrite(v[1])
			c, _ = tc.Int64ToTrite(v[2])
		)
		// t.Log(a, b, c)

		ab := a.Mul(b)
		if !(c.Equal(ab)) {
			t.Fatalf("sample[%d]: wrong %v", i, v)
		}
	}
}
