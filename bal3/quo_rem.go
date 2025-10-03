package bal3

import (
	"errors"
	"fmt"
	"math/big"
)

// https://en.wikipedia.org/wiki/Long_division
// Algorithm for arbitrary base

// https://en.wikipedia.org/wiki/Balanced_ternary
// Multi-trit division

const (
	// EnableCorrectionQuoRem = false
	EnableCorrectionQuoRem = true
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

//------------------------------------------------------------------------------

type divParams[Tryte GenericTryte] struct {
	dc doubleCore[Tryte]

	d   double[Tryte] // d - divisor: { hi: 0, lo: b }
	hd  double[Tryte] // hd - half divisor (d/2)
	hdn double[Tryte] // hdn - half divisor negative (-d/2)

	signD int

	r double[Tryte]
}

func makeDivParams[Tryte GenericTryte](tc TryteCore[Tryte], a, b Tryte) *divParams[Tryte] {

	dc := makeDoubleCore[Tryte](tc)

	var (
		r = makeDouble[Tryte](0, a)
		d = makeDouble[Tryte](0, b)
	)

	var (
		//		d   Double[Tryte] // d - divisor: { hi: 0, lo: b }
		hd  double[Tryte] // hd - half divisor (d/2)
		hdn double[Tryte] // hdn - half divisor negative (-d/2)
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
	hdn = dc.Inverse(hd)

	return &divParams[Tryte]{
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
		return +1
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

func correctionQuoRemV1[T GenericTryte](tc TryteCore[T], a, b T, q, r T) (cq, cr T) {
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

func correctionQuoRemV2[T GenericTryte](tc TryteCore[T], a, b T, q, r T) (cq, cr T) {

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

func correctionQuoRem[Tryte GenericTryte](tc TryteCore[Tryte], a, b Tryte, q, r Tryte) (cq, cr Tryte) {
	return correctionQuoRemV1(tc, a, b, q, r)
	//return correctionQuoRemV2(tc, a, b, q, r)
}

//------------------------------------------------------------------------------

func tryteQuoRem[Tryte GenericTryte](tc TryteCore[Tryte], a, b Tryte) (q, r Tryte) {

	if tc.IsZero(b) {
		panic(errDivisionByZero)
	}

	dp := makeDivParams(tc, a, b)

	for i := tc.n; i > 0; { // backward iterate
		i--

		t := dp.divIter()
		q = setTrit(q, i, Trit(t)) // q[i] = t
	}

	r = dp.outRem()

	// QuoRem correction
	if EnableCorrectionQuoRem {
		q, r = correctionQuoRem(tc, a, b, q, r)
	}

	return
}

//------------------------------------------------------------------------------
