package bal3

import (
	"errors"
	"fmt"
	"math/big"
)

var errDivisionByZero = errors.New("division by zero")

func _() {

	x := big.NewInt(2)
	y := big.NewInt(3)
	var r big.Int

	//x.SetBit()
	//x.Bit()
	//x.BitLen()

	z, rr := x.QuoRem(x, y, &r)
	_, _ = z, rr

	//x.BitLen()

	fmt.Println()

	// natdiv.go
	// func (q nat) divBasic(u, v nat) {
}

// numerator / denominator

// https://en.wikipedia.org/wiki/Long_division
// Algorithm for arbitrary base

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

func tryteQuoRemLo[T Unsigned](tc TryteCore[T], a, b T) (q, r T) {

	k := tc.n / 2

	var (
		signA = tc.Sign(a)
		signB = tc.Sign(b)
	)

	switch signB {
	case 0: // b == 0
		panic(errDivisionByZero)
	case -1: // b < 0
		a = tc.Neg(a)
		b = tc.Neg(b)
	}

	//---------------------------------------------------------------
	dd := b          // double divisor
	var halfDivPos T // positive

	{
		// 0.5_base10 = (0.1111111...)_bal3

		// |<---n--->|<---n--->|
		// |000...000,111...111|
		var factor T // set lo n trits.
		for i := 0; i < k; i++ {
			factor = tc.setTrit(factor, i, 1)
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
	halfDivNeg := tc.Neg(halfDivPos) // negative

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
		for i := k - 1; i >= 0; i-- {
			rr = tc.Shl(rr, 1) // r = r << 1
			t := Trit(iter())
			q = tc.setTrit(q, i, t)
		}
	} else {

		fmt.Println("divisor:", tc.Format(dd))
		fmt.Println("hd_neg: ", tc.Format(halfDivNeg))
		fmt.Println("hd_pos: ", tc.Format(halfDivPos))
		fmt.Println()

		for i := k - 1; i >= 0; i-- {

			rr = tc.Shl(rr, 1) // r = r << 1
			rr_prev := rr
			_ = rr_prev

			t := Trit(iter())

			fmt.Println(tc.Format(rr_prev))
			fmt.Println(tc.Format(rr))
			fmt.Println(t)
			fmt.Println()

			q = tc.setTrit(q, i, t)
		}
	}

	//---------------------------------------------------------------------

	rr = tc.Shr(rr, k) // r = r >> n
	r = rr             // r = rr.lo

	// b < 0
	if signB == -1 {
		r = tc.Neg(r)
	}

	// correction
	if true {
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

//------------------------------------------------------------------------------

func tryteQuoRem[T Unsigned](tc TryteCore[T], a, b T) (q, r T) {

	if tc.IsZero(b) {
		panic(errDivisionByZero)
	}

	dp := makeDivParams(tc, a, b)

	for i := tc.n - 1; i >= 0; i-- {
		t := dp.divIter()
		q = tc.setTrit(q, i, Trit(t)) // q[i] = t
	}

	r = dp.outRem()

	// correction q and r
	q, r = correctionQuoRem(tc, a, b, q, r)

	return
}

//------------------------------------------------------------------------------

type divParams[T Unsigned] struct {
	dc doubleCore[T]

	d   double[T] // d - divisor: { hi: 0, lo: b }
	hd  double[T] // hd - half divisor (d/2)
	hdn double[T] // hdn - half divisor negative (-d/2)

	signD int

	r double[T]
}

func makeDivParams[T Unsigned](tc TryteCore[T], a, b T) *divParams[T] {

	dc := makeDoubleCore[T](tc)

	var (
		r = makeDouble[T](0, a)
		d = makeDouble[T](0, b)
	)

	var (
		//		d   Double[T] // d - divisor: { hi: 0, lo: b }
		hd  double[T] // hd - half divisor (d/2)
		hdn double[T] // hdn - half divisor negative (-d/2)
	)

	// 0.5_base10 = (0.1111111...)_bal3

	// |<---n--->|<---n--->|
	// |000...000,111...111|
	factor := makeDouble(0, tc.SetAllTrits(1)) // set lo n trits.

	//      |<---n--->|
	// d * 0,111...111|000...
	_, d1 := dc.Mul(d, factor)

	//      |<---n--->|
	// d * 0,000...000|111...
	// d2 = d1 >> n
	d2 := dc.Shr(d1, tc.n)

	d = dc.Shl(d, tc.n) // d = d << n
	hd, _ = dc.Add(d1, d2, 0)
	hdn = dc.Neg(hd)

	return &divParams[T]{
		dc: dc,

		d:   d,
		hd:  hd,
		hdn: hdn,

		signD: dc.Sign(d),

		r: r,
	}
}

func (p *divParams[T]) divIterV1() int {

	dc := p.dc

	if p.signD == 1 { // d > 0

		// hdn < 0 < hd < d

		c1 := dc.Compare(p.r, p.hd)
		if c1 == 1 { // rr > hd
			p.r, _ = dc.Sub(p.r, p.d, 0)
			return 1
		}

		c2 := dc.Compare(p.r, p.hdn)
		if c2 == -1 { // rr < hdn
			p.r, _ = dc.Add(p.r, p.d, 0)
			return -1
		}

		return 0
	}

	if p.signD == -1 { // d < 0

		// d < hd < 0 < hdn

		c1 := dc.Compare(p.r, p.hd)
		if c1 == -1 { // rr < hd
			p.r, _ = dc.Sub(p.r, p.d, 0)
			return 1
		}

		c2 := dc.Compare(p.r, p.hdn)
		if c2 == 1 { // rr > hdn
			p.r, _ = dc.Add(p.r, p.d, 0)
			return -1
		}

		return 0
	}

	return 0
}

func (p *divParams[T]) divIterV2() int {

	dc := p.dc

	c1 := dc.Compare(p.r, p.hd)
	if c1 == p.signD {
		p.r, _ = dc.Sub(p.r, p.d, 0)
		return 1
	}

	c2 := dc.Compare(p.r, p.hdn)
	if c2 == -p.signD {
		p.r, _ = dc.Add(p.r, p.d, 0)
		return -1
	}
	return 0
}

func (p *divParams[T]) divIterV3() int {

	dc := p.dc

	if dc.Compare(p.r, p.hd) == p.signD {
		p.r, _ = dc.Sub(p.r, p.d, 0)
		return 1
	}
	if dc.Compare(p.hdn, p.r) == p.signD {
		p.r, _ = dc.Add(p.r, p.d, 0)
		return -1
	}
	return 0
}

func (p *divParams[T]) divIter() int {

	p.r = p.dc.Shl(p.r, 1) // r = r << 1

	//return p.divIterV1()
	//return p.divIterV2()
	return p.divIterV3()
}

func (p *divParams[T]) outRem() T {
	return p.r.Hi
}

//------------------------------------------------------------------------------

// correction q and r

func correctionQuoRemV1[T Unsigned](tc TryteCore[T], a, b T, q, r T) (cq, cr T) {
	one, _ := tc.Int64ToTrite(1)
	var (
		signA = tc.Sign(a)
		signB = tc.Sign(b)
		signR = tc.Sign(r)
	)
	if (signA == 1) && (signR == -1) { // (a > 0) && (r < 0)
		if signB == 1 { // b > 0
			q, _ = tc.Sub(q, one, 0) // q -= 1
			r, _ = tc.Add(r, b, 0)
		}
		if signB == -1 { // b < 0
			q, _ = tc.Add(q, one, 0) // q += 1
			r, _ = tc.Sub(r, b, 0)
		}
	}
	if (signA == -1) && (signR == 1) { // (a < 0) && (r > 0)
		if signB == 1 { // b > 0
			q, _ = tc.Add(q, one, 0) // q += 1
			r, _ = tc.Sub(r, b, 0)
		}
		if signB == -1 { // b < 0
			q, _ = tc.Sub(q, one, 0) // q -= 1
			r, _ = tc.Add(r, b, 0)
		}
	}
	return q, r
}

func correctionQuoRemV2[T Unsigned](tc TryteCore[T], a, b T, q, r T) (cq, cr T) {

	var (
		signA = tc.Sign(a)
		signB = tc.Sign(b)
		signR = tc.Sign(r)
	)

	if (signA == 1) && (signR == -1) { // (a > 0) && (r < 0)
		if signB == 1 { // b > 0
			q, _ = tc.Sub(q, 0, -1) // q -= 1
			r, _ = tc.Add(r, b, 0)
		}
		if signB == -1 { // b < 0
			q, _ = tc.Add(q, 0, 1) // q += 1
			r, _ = tc.Sub(r, b, 0)
		}
	}

	if (signA == -1) && (signR == 1) { // (a < 0) && (r > 0)
		if signB == 1 { // b > 0
			q, _ = tc.Add(q, 0, 1) // q += 1
			r, _ = tc.Sub(r, b, 0)
		}
		if signB == -1 { // b < 0
			q, _ = tc.Sub(q, 0, -1) // q -= 1
			r, _ = tc.Add(r, b, 0)
		}
	}

	return q, r
}

func correctionQuoRem[T Unsigned](tc TryteCore[T], a, b T, q, r T) (cq, cr T) {
	return correctionQuoRemV1(tc, a, b, q, r)
	//return correctionQuoRemV2(tc, a, b, q, r)
}
