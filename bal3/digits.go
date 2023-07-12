package bal3

import (
	"github.com/gitchander/gobal3/utils/digits"
)

func intToTriteV1[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	var t int
	for i := 0; i < tc.n; i++ {
		v, t = digits.QuoRemMinMax(v, tritMin, tritMax)
		a = tc.setTrit(a, i, t)
	}
	rest = v
	return
}

func intToTriteV2[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	ds := make([]int, tc.n)
	rest = digits.CalcDigits(v, tritMin, tritMax, ds)
	for i, d := range ds {
		a = tc.setTrit(a, i, d)
	}
	return
}

func intToTriteV3[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	var t int
	for i := 0; i < tc.n; i++ {
		v, t = quoRemBal3(v)
		a = tc.setTrit(a, i, t)
	}
	rest = v
	return
}

func intToTriteRest[T Unsigned](tc TryteCore[T], v int) (a T, rest int) {
	//return intToTriteV1(tc, v)
	//return intToTriteV2(tc, v)
	return intToTriteV3(tc, v)
}

func tryteToIntV1[T Unsigned](tc TryteCore[T], a T) int {
	var v int
	p := 1
	for i := 0; i < tc.n; i++ {
		v += p * tc.getTrit(a, i)
		p *= base
	}
	return v
}

func tryteToIntV2[T Unsigned](tc TryteCore[T], a T) int {
	ds := make([]int, tc.n)
	for i := range ds {
		ds[i] = tc.getTrit(a, i)
	}
	return digits.CalcNumber(tritMin, tritMax, ds, 0)
}

func tryteToInt[T Unsigned](tc TryteCore[T], a T) int {
	return tryteToIntV1(tc, a)
	//return tryteToIntV2(tc, a)
}
