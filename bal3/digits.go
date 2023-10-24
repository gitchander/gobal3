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
		a = tc.setTrit(a, i, Trit(t))
	}
	rest = v
	return
}

func tryteToInt64V1[T Unsigned](tc TryteCore[T], a T, rest int64) (int64, bool) {
	base := int64(digiter.Base())
	v := rest
	for i := (tc.n - 1); i >= 0; i-- {
		v = (v * base) + int64(tc.getTrit(a, i))
	}
	return v, true
}

//------------------------------------------------------------------------------

func intToTriteV2[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	ds := make([]int, tc.n)
	rest = digiter.IntToDigits(v, ds)
	for i, d := range ds {
		t := Trit(d)
		a = tc.setTrit(a, i, t)
	}
	return
}

func tryteToIntV2[T Unsigned](tc TryteCore[T], a T, rest int) int {
	ds := make([]int, tc.n)
	for i := range ds {
		t := tc.getTrit(a, i)
		ds[i] = int(t)
	}
	v, err := digiter.DigitsToInt(ds, rest)
	if err != nil {
		panic(err)
	}
	return v
}

//------------------------------------------------------------------------------

func int64ToTriteV3[T Unsigned](tc TryteCore[T], v int64) (a T, rest int64) {
	var t int64
	for i := 0; i < tc.n; i++ {
		v, t = quoRemBal3(v)
		a = tc.setTrit(a, i, Trit(t))
	}
	rest = v
	return
}

//------------------------------------------------------------------------------

func int64ToTrite[T Unsigned](tc TryteCore[T], v int64) (a T, rest int64) {
	//return intToTriteV1(tc, v)
	//return intToTriteV2(tc, v)
	return int64ToTriteV3(tc, v)
}

func tryteToInt64[T Unsigned](tc TryteCore[T], a T, rest int64) (int64, bool) {
	return tryteToInt64V1(tc, a, rest)
	//return tryteToIntV2(tc, a, rest)
}
