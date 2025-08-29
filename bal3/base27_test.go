package bal3

import (
	"fmt"
	"testing"
)

func testBase27[T coreTryte](tc TryteCore[T], a T) error {

	sa := FormatBase27(tc, a)

	b, err := ParseBase27(tc, sa)
	if err != nil {
		return err
	}

	if !(tc.Equal(a, b)) {
		var (
			fa = tc.Format(a)
			fb = tc.Format(b)
		)
		return fmt.Errorf("testBase27: %s %s %s", fa, sa, fb)
	}

	return nil
}

func testBase27Bounds[T coreTryte](tc TryteCore[T]) error {
	min, max := tc.LimitsInt64()
	for av := min; av <= max; av++ {
		a, _ := tc.Int64ToTrite(av)
		err := testBase27(tc, a)
		if err != nil {
			return err
		}
	}
	return nil
}

func testBase27Rand[T coreTryte](tc TryteCore[T]) error {
	r := newRandNext()
	for i := 0; i < 1000; i++ {
		a := tc.RandSh(r)
		err := testBase27(tc, a)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestBase27T4Bounds(t *testing.T) {
	err := testBase27Bounds(TC4)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBase27T32Rand(t *testing.T) {
	err := testBase27Rand(TC32)
	if err != nil {
		t.Fatal(err)
	}
}
