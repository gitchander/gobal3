package bal3

import (
	"github.com/gitchander/gobal3/utils/digits"
)

var digiter = digits.NewDigiter(-1, +1)

//------------------------------------------------------------------------------

func intToTriteRestV1[T GenericTryte](tc TryteCore[T], v int) (a T, rest int) {
	var t int
	for i := 0; i < tc.n; i++ {
		v, t = digiter.RestDigit(v)
		a = setTrit(a, i, Trit(t))
	}
	rest = v
	return
}

func tryteRestToInt64V1[T GenericTryte](tc TryteCore[T], a T, rest int64) (int64, bool) {
	base := int64(digiter.Base())
	v := rest
	for i := tc.n; i > 0; { // backward iterate
		i--

		v = (v * base) + int64(getTrit(a, i))
	}
	return v, true
}

//------------------------------------------------------------------------------

func intToTriteRestV2[T GenericTryte](tc TryteCore[T], v int) (a T, rest int) {
	ds := make([]int, tc.n)
	rest = digiter.IntToDigits(v, ds)
	for i, d := range ds {
		t := Trit(d)
		a = setTrit(a, i, t)
	}
	return
}

func tryteRestToIntV2[T GenericTryte](tc TryteCore[T], a T, rest int) int {
	ds := make([]int, tc.n)
	for i := range ds {
		t := getTrit(a, i)
		ds[i] = int(t)
	}
	v, err := digiter.DigitsToInt(ds, rest)
	if err != nil {
		panic(err)
	}
	return v
}

//------------------------------------------------------------------------------

func int64ToTriteRestV3[T GenericTryte](tc TryteCore[T], v int64) (a T, rest int64) {
	var t int64
	for i := 0; i < tc.n; i++ {
		v, t = quoRemBal3(v)
		a = setTrit(a, i, Trit(t))
	}
	rest = v
	return
}

//------------------------------------------------------------------------------

func int64ToTriteRest[T GenericTryte](tc TryteCore[T], v int64) (a T, rest int64) {
	//return intToTriteRestV1(tc, v)
	//return intToTriteRestV2(tc, v)
	return int64ToTriteRestV3(tc, v)
}

func tryteRestToInt64[T GenericTryte](tc TryteCore[T], a T, rest int64) (int64, bool) {
	return tryteRestToInt64V1(tc, a, rest)
	//return tryteRestToIntV2(tc, a, rest)
}

//------------------------------------------------------------------------------
