package bal3

import (
	"github.com/gitchander/gobal3/utils/digits"
)

var digiter = digits.NewDigiter(tritMin, tritMax)

//------------------------------------------------------------------------------

func intToTriteV1[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	var t int
	for i := 0; i < tc.n; i++ {
		t, v = digiter.Digit(v)
		a = tc.setTrit(a, i, t)
	}
	rest = v
	return
}

func tryteToIntV1[T Unsigned](tc TryteCore[T], a T, rest int) int {
	base := digiter.Base()
	v := rest
	for i := (tc.n - 1); i >= 0; i-- {
		v = (v * base) + tc.getTrit(a, i)
	}
	return v
}

//------------------------------------------------------------------------------

func intToTriteV2[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	ds := make([]int, tc.n)
	rest = digiter.IntToDigits(v, ds)
	for i, d := range ds {
		a = tc.setTrit(a, i, d)
	}
	return
}

func tryteToIntV2[T Unsigned](tc TryteCore[T], a T, rest int) int {
	ds := make([]int, tc.n)
	for i := range ds {
		ds[i] = tc.getTrit(a, i)
	}
	v, err := digiter.DigitsToInt(ds, rest)
	if err != nil {
		panic(err)
	}
	return v
}

//------------------------------------------------------------------------------

func intToTriteV3[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	var t int
	for i := 0; i < tc.n; i++ {
		v, t = quoRemBal3(v)
		a = tc.setTrit(a, i, t)
	}
	rest = v
	return
}

//------------------------------------------------------------------------------

func intToTrite[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	//return intToTriteV1(tc, v)
	//return intToTriteV2(tc, v)
	return intToTriteV3(tc, v)
}

func tryteToInt[T Unsigned](tc TryteCore[T], a T, rest int) int {
	return tryteToIntV1(tc, a, rest)
	//return tryteToIntV2(tc, a, rest)
}
