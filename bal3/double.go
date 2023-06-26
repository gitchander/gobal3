package bal3

//------------------------------------------------------------------------------

type Double[T Unsigned] struct {
	Hi, Lo T
}

func MakeDouble[T Unsigned](hi, lo T) Double[T] {
	return Double[T]{
		Hi: hi,
		Lo: lo,
	}
}

//------------------------------------------------------------------------------

type DoubleCore[T Unsigned] struct {
	tc TryteCore[T]
}

func MakeDoubleCore[T Unsigned](tc TryteCore[T]) DoubleCore[T] {
	return DoubleCore[T]{
		tc: tc,
	}
}

//------------------------------------------------------------------------------

// TotalTrits returns the total number of trits.
func (dc DoubleCore[T]) TotalTrits() int {
	return 2 * dc.tc.n
}

//------------------------------------------------------------------------------

func (dc DoubleCore[T]) wantShl(a Double[T], i int) Double[T] {
	checkShiftAmount(i)
	m := dc.TotalTrits()
	var b Double[T]
	for j := i; j < m; j++ {
		b = dc.SetTrit(b, j, dc.GetTrit(a, j-i))
	}
	return b
}

func (dc DoubleCore[T]) wantShr(a Double[T], i int) Double[T] {
	checkShiftAmount(i)
	m := dc.TotalTrits()
	var b Double[T]
	for j := m - 1 - i; j >= 0; j-- {
		b = dc.SetTrit(b, j, dc.GetTrit(a, j+i))
	}
	return b
}

func (dc DoubleCore[T]) Shl(a Double[T], i int) Double[T] {
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
	return MakeDouble(hi, lo)
}

func (dc DoubleCore[T]) Shr(a Double[T], i int) Double[T] {
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
	return MakeDouble(hi, lo)
}

//------------------------------------------------------------------------------

func (dc DoubleCore[T]) SetTrit(a Double[T], i int, t int) Double[T] {
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

func (dc DoubleCore[T]) GetTrit(a Double[T], i int) int {
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

func (dc DoubleCore[T]) FromInt(v int) Double[T] {
	tc := dc.tc
	var lo, hi T
	lo, v = tc.FromIntRest(v)
	hi, v = tc.FromIntRest(v)
	return MakeDouble(hi, lo)
}

func (dc DoubleCore[T]) ToInt(a Double[T]) int {
	tc := dc.tc
	var (
		hi = tc.ToInt(a.Hi)
		lo = tc.ToInt(a.Lo)
	)
	return hi*powersOfThree[tc.n] + lo
}

func (dc DoubleCore[T]) ToStringAll(a Double[T]) string {
	tc := dc.tc
	var (
		hi = tc.FormatAllTrits(a.Hi)
		lo = tc.FormatAllTrits(a.Lo)
	)
	return hi + "_" + lo
}

func (dc DoubleCore[T]) Invert(a Double[T]) Double[T] {
	tc := dc.tc
	return Double[T]{
		Hi: tc.Invert(a.Hi),
		Lo: tc.Invert(a.Lo),
	}
}

//------------------------------------------------------------------------------

func (dc DoubleCore[T]) Compare(a, b Double[T]) int {
	tc := dc.tc
	c := tc.Compare(a.Hi, b.Hi)
	if c != 0 {
		return c
	}
	return tc.Compare(a.Lo, b.Lo)
}

// a == n
func (dc DoubleCore[T]) Equal(a, b Double[T]) bool {
	return dc.Compare(a, b) == 0
}

// a < b
func (dc DoubleCore[T]) Less(a, b Double[T]) bool {
	return dc.Compare(a, b) == -1
}

// a > b
func (dc DoubleCore[T]) Greater(a, b Double[T]) bool {
	return dc.Compare(a, b) == 1
}

//------------------------------------------------------------------------------

func (dc DoubleCore[T]) Sign(x Double[T]) int {
	tc := dc.tc
	sign := tc.Sign(x.Hi)
	if sign != 0 {
		return sign
	}
	return tc.Sign(x.Lo)
}

//------------------------------------------------------------------------------

func (dc DoubleCore[T]) Rand(r *Rand) Double[T] {
	tc := dc.tc
	return Double[T]{
		Hi: tc.RandSh(r),
		Lo: tc.RandSh(r),
	}
}

func (dc DoubleCore[T]) RandSh(r *Rand) Double[T] {
	a := dc.Rand(r)
	return dc.Shr(a, r.Intn(dc.TotalTrits()))
}

//------------------------------------------------------------------------------

func (dc DoubleCore[T]) Add(a, b Double[T], carryIn int) (c Double[T], carryOut int) {
	tc := dc.tc
	var lo, hi T
	carry := carryIn
	lo, carry = tc.Add(a.Lo, b.Lo, carry)
	hi, carry = tc.Add(a.Hi, b.Hi, carry)
	return MakeDouble(hi, lo), carry
}

func (dc DoubleCore[T]) Sub(a, b Double[T], carryIn int) (c Double[T], carryOut int) {
	tc := dc.tc
	var lo, hi T
	carry := carryIn
	lo, carry = tc.Sub(a.Lo, b.Lo, carry)
	hi, carry = tc.Sub(a.Hi, b.Hi, carry)
	return MakeDouble(hi, lo), carry
}

func (dc DoubleCore[T]) Mul(a, b Double[T]) (hi, lo Double[T]) {

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

	hi = MakeDouble(a3, a2)
	lo = MakeDouble(a1, a0)

	return
}

func (dc DoubleCore[T]) Bounds() (min, max int) {
	return tryteBounds(dc.TotalTrits())
}
