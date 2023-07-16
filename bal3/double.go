package bal3

//------------------------------------------------------------------------------

type double[T Unsigned] struct {
	Hi, Lo T
}

func makeDouble[T Unsigned](hi, lo T) double[T] {
	return double[T]{
		Hi: hi,
		Lo: lo,
	}
}

//------------------------------------------------------------------------------

type doubleCore[T Unsigned] struct {
	tc TryteCore[T]
}

func makeDoubleCore[T Unsigned](tc TryteCore[T]) doubleCore[T] {
	return doubleCore[T]{
		tc: tc,
	}
}

//------------------------------------------------------------------------------

// TotalTrits returns the total number of trits.
func (dc doubleCore[T]) TotalTrits() int {
	return 2 * dc.tc.n
}

//------------------------------------------------------------------------------

func (dc doubleCore[T]) wantShl(a double[T], i int) double[T] {
	checkShiftAmount(i)
	m := dc.TotalTrits()
	var b double[T]
	for j := i; j < m; j++ {
		b = dc.SetTrit(b, j, dc.GetTrit(a, j-i))
	}
	return b
}

func (dc doubleCore[T]) wantShr(a double[T], i int) double[T] {
	checkShiftAmount(i)
	m := dc.TotalTrits()
	var b double[T]
	for j := m - 1 - i; j >= 0; j-- {
		b = dc.SetTrit(b, j, dc.GetTrit(a, j+i))
	}
	return b
}

func (dc doubleCore[T]) Shl(a double[T], i int) double[T] {
	checkShiftAmount(i)
	var (
		tc = dc.tc
		n  = tc.n
	)
	var hi, lo T
	switch {
	case i < n:
		hi = tc.Shl(a.Hi, i)
		hi |= tc.Shr(a.Lo, n-i)
		lo = tc.Shl(a.Lo, i)
	case i < (2 * n):
		hi = tc.Shl(a.Lo, i-n)
	default:
		// set zero
	}
	return makeDouble(hi, lo)
}

func (dc doubleCore[T]) Shr(a double[T], i int) double[T] {
	checkShiftAmount(i)
	var (
		tc = dc.tc
		n  = tc.n
	)
	var hi, lo T
	switch {
	case i < n:
		lo = tc.Shr(a.Lo, i)
		lo |= tc.Shl(a.Hi, n-i)
		hi = tc.Shr(a.Hi, i)
	case i < (2 * n):
		lo = tc.Shr(a.Hi, i-n)
	default:
		// set zero
	}
	return makeDouble(hi, lo)
}

//------------------------------------------------------------------------------

func (dc doubleCore[T]) SetTrit(a double[T], i int, t int) double[T] {
	var (
		tc = dc.tc
		n  = tc.n
	)
	switch {
	case i < n:
		a.Lo = tc.setTrit(a.Lo, i, t)
	case i < (2 * n):
		a.Hi = tc.setTrit(a.Hi, i-n, t)
	}
	return a
}

func (dc doubleCore[T]) GetTrit(a double[T], i int) int {
	var (
		tc = dc.tc
		n  = tc.n
	)
	var t int
	switch {
	case i < n:
		t = tc.getTrit(a.Lo, i)
	case i < (2 * n):
		t = tc.getTrit(a.Hi, i-n)
	}
	return t
}

func (dc doubleCore[T]) IntToDouble(v int) (d double[T], rest int) {
	tc := dc.tc
	var lo, hi T
	lo, v = tc.IntToTrite(v)
	hi, v = tc.IntToTrite(v)
	d = makeDouble(hi, lo)
	rest = v
	return d, rest
}

func (dc doubleCore[T]) DoubleToInt(a double[T], rest int) int {
	tc := dc.tc
	v := rest
	v = tc.TryteToInt(a.Hi, v)
	v = tc.TryteToInt(a.Lo, v)
	return v
}

func (dc doubleCore[T]) ToStringAll(a double[T]) string {
	tc := dc.tc
	var (
		hi = tc.FormatAllTrits(a.Hi)
		lo = tc.FormatAllTrits(a.Lo)
	)
	return hi + "_" + lo
}

func (dc doubleCore[T]) Invert(a double[T]) double[T] {
	tc := dc.tc
	return double[T]{
		Hi: tc.Invert(a.Hi),
		Lo: tc.Invert(a.Lo),
	}
}

//------------------------------------------------------------------------------

func (dc doubleCore[T]) Compare(a, b double[T]) int {
	tc := dc.tc
	c := tc.Compare(a.Hi, b.Hi)
	if c != 0 {
		return c
	}
	return tc.Compare(a.Lo, b.Lo)
}

// a == n
func (dc doubleCore[T]) Equal(a, b double[T]) bool {
	return dc.Compare(a, b) == 0
}

// a < b
func (dc doubleCore[T]) Less(a, b double[T]) bool {
	return dc.Compare(a, b) == -1
}

// a > b
func (dc doubleCore[T]) Greater(a, b double[T]) bool {
	return dc.Compare(a, b) == 1
}

//------------------------------------------------------------------------------

func (dc doubleCore[T]) Sign(x double[T]) int {
	tc := dc.tc
	sign := tc.Sign(x.Hi)
	if sign != 0 {
		return sign
	}
	return tc.Sign(x.Lo)
}

//------------------------------------------------------------------------------

func (dc doubleCore[T]) Rand(r *Rand) double[T] {
	tc := dc.tc
	return double[T]{
		Hi: tc.RandSh(r),
		Lo: tc.RandSh(r),
	}
}

func (dc doubleCore[T]) RandSh(r *Rand) double[T] {
	a := dc.Rand(r)
	return dc.Shr(a, r.Intn(dc.TotalTrits()))
}

//------------------------------------------------------------------------------

func (dc doubleCore[T]) Add(a, b double[T], carryIn int) (c double[T], carryOut int) {
	tc := dc.tc
	var lo, hi T
	carry := carryIn
	lo, carry = tc.Add(a.Lo, b.Lo, carry)
	hi, carry = tc.Add(a.Hi, b.Hi, carry)
	return makeDouble(hi, lo), carry
}

func (dc doubleCore[T]) Sub(a, b double[T], carryIn int) (c double[T], carryOut int) {
	tc := dc.tc
	var lo, hi T
	carry := carryIn
	lo, carry = tc.Sub(a.Lo, b.Lo, carry)
	hi, carry = tc.Sub(a.Hi, b.Hi, carry)
	return makeDouble(hi, lo), carry
}

func (dc doubleCore[T]) Bounds() (min, max int) {
	return tryteBounds(dc.TotalTrits())
}

func (dc doubleCore[T]) Mul(a, b double[T]) (hi, lo double[T]) {

	tc := dc.tc

	var (
		hi_00, lo_00 = tc.Mul(a.Lo, b.Lo)
		hi_01, lo_01 = tc.Mul(a.Lo, b.Hi)
		hi_10, lo_10 = tc.Mul(a.Hi, b.Lo)
		hi_11, lo_11 = tc.Mul(a.Hi, b.Hi)
	)

	var (
		a0 = lo_00
		a1 = hi_00
		a2 = hi_01
		a3 = hi_11
	)

	carry := 0
	a1, carry = tc.Add(a1, lo_01, carry)
	a2, carry = tc.Add(a2, hi_10, carry)
	a3, carry = tc.Add(a3, 0, carry)

	carry = 0
	a1, carry = tc.Add(a1, lo_10, carry)
	a2, carry = tc.Add(a2, lo_11, carry)
	a3, carry = tc.Add(a3, 0, carry)

	hi = makeDouble(a3, a2)
	lo = makeDouble(a1, a0)

	return
}
