package digits

import (
	"fmt"
	"math"
	"testing"

	"github.com/gitchander/gobal3/utils/random"
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

	samples := makeLimitInts(10)

	r := randNow()
	//r := randBySeed(0)
	for i := 0; i < 100; i++ {
		min, max := randomBaseMinMax(r)
		// t.Log(min, max)
		for _, sample := range samples {
			err := testSample(sample, min, max, 1)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestRestDigitTri(t *testing.T) {

	// base 3
	const (
		min = -1
		max = +1
	)

	type sampleTypes struct {
		value       int
		rest, digit int
	}

	samples := []sampleTypes{
		{value: math.MinInt + 0, rest: -3074457345618258603, digit: 1},

		{value: math.MinInt + 1, rest: -3074457345618258602, digit: -1},
		{value: math.MinInt + 2, rest: -3074457345618258602, digit: 0},
		{value: math.MinInt + 3, rest: -3074457345618258602, digit: 1},

		{value: math.MinInt + 4, rest: -3074457345618258601, digit: -1},
		{value: math.MinInt + 5, rest: -3074457345618258601, digit: 0},
		{value: math.MinInt + 6, rest: -3074457345618258601, digit: 1},

		{value: -7, rest: -2, digit: -1},
		{value: -6, rest: -2, digit: 0},
		{value: -5, rest: -2, digit: +1},

		{value: -4, rest: -1, digit: -1},
		{value: -3, rest: -1, digit: 0},
		{value: -2, rest: -1, digit: +1},

		{value: -1, rest: 0, digit: -1},
		{value: 0, rest: 0, digit: 0},
		{value: 1, rest: 0, digit: +1},

		{value: 2, rest: 1, digit: -1},
		{value: 3, rest: 1, digit: 0},
		{value: 4, rest: 1, digit: +1},

		{value: 5, rest: 2, digit: -1},
		{value: 6, rest: 2, digit: 0},
		{value: 7, rest: 2, digit: +1},

		{value: math.MaxInt - 5, rest: 3074457345618258601, digit: -1},
		{value: math.MaxInt - 4, rest: 3074457345618258601, digit: 0},
		{value: math.MaxInt - 3, rest: 3074457345618258601, digit: 1},

		{value: math.MaxInt - 2, rest: 3074457345618258602, digit: -1},
		{value: math.MaxInt - 1, rest: 3074457345618258602, digit: 0},
		{value: math.MaxInt - 0, rest: 3074457345618258602, digit: 1},
	}

	for _, sample := range samples {
		rest, digit := RestDigit(sample.value, min, max)
		checkHaveWant(t, "rest", rest, sample.rest)
		checkHaveWant(t, "digit", digit, sample.digit)
	}
}

func TestRestDigitRand(t *testing.T) {
	r := randNow()
	randBase := func() (min, max int) {
		min = random.RandIntIn(r, -100, 100+1)
		max = random.RandIntIn(r, -100, 100+1)
		if min > max {
			min, max = max, min
		}
		return min, max
	}
	randValue := func() int {
		return random.RandIntIn(r, -1000, 1000+1)
	}
	wantRestDigit := func(x int, min, max int) (rest, digit int) {
		base := max - min + 1
		digit = x
		for digit < min {
			rest--
			digit += base
		}
		for digit > max {
			rest++
			digit -= base
		}
		return rest, digit
	}
	for i := 0; i < 1000; i++ {
		var (
			value    = randValue()
			min, max = randBase() // rand base
		)
		var (
			haveRest, haveDigit = RestDigit(value, min, max)
			wantRest, wantDigit = wantRestDigit(value, min, max)
		)
		checkHaveWant(t, "rest", haveRest, wantRest)
		checkHaveWant(t, "digit", haveDigit, wantDigit)
	}
}

func checkHaveWant(t *testing.T, name string, have, want int) {
	if have != want {
		err := fmt.Errorf("invalid %s: have %d, want %d", name, have, want)
		t.Fatal(err)
	}
}

func makeLimitInts(d int) []int {
	if d < 0 {
		return nil
	}
	var as []int
	as = appendIntsMinMax(as, math.MinInt, math.MinInt+d+1)
	as = appendIntsMinMax(as, -d, d+1)
	if d > 0 {
		as = appendIntsMinMax(as, math.MaxInt-d+1, math.MaxInt)
		as = append(as, math.MaxInt)
	}
	return as
}

func appendIntsMinMax(as []int, min, max int) []int {
	for a := min; a < max; a++ {
		as = append(as, a)
	}
	return as
}
