package bal3

import (
	"fmt"
)

// numerator / denominator

func quoRemT16_v1(a, b Tryte16) (quo, rem Tryte16) {
	var (
		da = Tryte32(a) // double 'a'
		db = Tryte32(b) // double 'b'
	)
	tc := TC32
	q, r := tryteQuoRemLo(tc, da, db)
	quo = Tryte16(q)
	rem = Tryte16(r)
	return
}

func quoRemT16_v2(a, b Tryte16) (quo, rem Tryte16) {
	tc := TC16
	return tc.QuoRem(a, b)
}

//     |<---n--->|
// A: 0.111...111|111... = (1/2)_base10 = (1/1T)_bal3
// B: 0.000...000|111... = ((1/2)/(3^n))_base10 = (1/(2*(3^n)))_base10
// C: 0.111...111|000... = (1/2) - (1/(2*(3^n)))

// d*(1/2)_base10 = d*0.5_base10 = d*A
// d*A = d*B + d*C

func tryteQuoRemLo[Tryte GenericTryte](tc TryteCore[Tryte], a, b Tryte) (q, r Tryte) {

	k := tc.n / 2

	var (
		signA = tc.Sign(a)
		signB = tc.Sign(b)
	)

	switch signB {
	case 0: // b == 0
		panic(errDivisionByZero)
	case -1: // b < 0
		a = tc.Inverse(a)
		b = tc.Inverse(b)
	}

	//---------------------------------------------------------------
	dd := b              // double divisor
	var halfDivPos Tryte // positive

	{
		// 0.5 (base10) = 0.1111111... (bal3)

		// |<---n--->|<---n--->|
		// |000...000,111...111|
		var factor Tryte // set lo n trits.
		for i := 0; i < k; i++ {
			factor = setTrit(factor, i, 1)
		}

		//      |<---n--->|
		// d * 0,111...111|000...
		//d1 := tc.MulLo(dd, factor)
		_, d1 := tc.Mul(dd, factor)

		//      |<---n--->|
		// d * 0,000...000|111...
		// d2 = d1 >> n
		d2 := tc.Shr(d1, k)

		halfDivPos, _ = tc.Add(d1, d2, 0)
	}

	dd = tc.Shl(dd, k)
	halfDivNeg := tc.Inverse(halfDivPos) // negative

	//---------------------------------------------------------------

	rr := a // (0, a.lo)

	iter := func() int {

		c1 := tc.Compare(rr, halfDivNeg)
		if c1 == -1 { // (rr < halfDivNeg)
			rr, _ = tc.Add(rr, dd, 0)
			return -1
		}

		c2 := tc.Compare(rr, halfDivPos)
		if c2 == 1 { // (rr > halfDivPos)
			rr, _ = tc.Sub(rr, dd, 0)
			return 1
		}

		return 0
	}

	if true {
		for i := k; i > 0; { // backward iterate
			i--

			rr = tc.Shl(rr, 1) // r = r << 1
			t := Trit(iter())
			q = setTrit(q, i, t)
		}
	} else {

		fmt.Println("divisor:", tc.Format(dd))
		fmt.Println("hd_neg: ", tc.Format(halfDivNeg))
		fmt.Println("hd_pos: ", tc.Format(halfDivPos))
		fmt.Println()

		for i := k; i > 0; { // backward iterate
			i--

			rr = tc.Shl(rr, 1) // r = r << 1
			rr_prev := rr
			_ = rr_prev

			t := Trit(iter())

			fmt.Println(tc.Format(rr_prev))
			fmt.Println(tc.Format(rr))
			fmt.Println(t)
			fmt.Println()

			q = setTrit(q, i, t)
		}
	}

	//---------------------------------------------------------------------

	rr = tc.Shr(rr, k) // r = r >> n
	r = rr             // r = rr.lo

	// b < 0
	if signB == -1 {
		r = tc.Inverse(r)
	}

	// QuoRem correction
	if EnableCorrectionQuoRem {
		one, _ := tc.Int64ToTrite(1)

		if (signA == 1) && tc.IsNegative(r) {
			r, _ = tc.Add(r, b, 0)
			if tc.IsPositive(q) {
				q, _ = tc.Sub(q, one, 0)
			} else {
				q, _ = tc.Add(q, one, 0)
			}
		}

		if (signA == -1) && tc.IsPositive(r) {
			r, _ = tc.Sub(r, b, 0)
			if tc.IsPositive(q) {
				q, _ = tc.Sub(q, one, 0)
			} else {
				q, _ = tc.Add(q, one, 0)
			}
		}
	}

	return
}
