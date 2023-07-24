package digits

import (
	"fmt"
	"math"
	"testing"
)

func testSample(a int, min, max int, dn int) error {
	d := NewDigiter(min, max)
	digits := make([]int, dn)
	rest := d.IntToDigits(a, digits)
	b, err := d.DigitsToInt(digits, rest)
	if err != nil {
		return err
	}
	if b != a {
		return fmt.Errorf("%d != %d", b, a)
	}
	return nil
}

func TestDigitsSamples(t *testing.T) {

	var samples []int
	samples = appendIntsRange(samples, math.MinInt, math.MinInt+100)
	samples = appendIntsRange(samples, -100, +100)
	samples = appendIntsRange(samples, math.MaxInt-100, math.MaxInt)
	samples = append(samples, math.MaxInt)

	r := randNow()
	//r := randBySeed(0)
	for i := 0; i < 100; i++ {
		min, max := randomBaseMinMax(r)
		t.Log(min, max)
		for _, sample := range samples {
			//fmt.Println(sample)
			err := testSample(sample, min, max, 40)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func appendIntsRange(as []int, min, max int) []int {
	for a := min; a < max; a++ {
		as = append(as, a)
	}
	return as
}
