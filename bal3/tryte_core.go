package bal3

import (
	"fmt"

	"github.com/gitchander/gobal3/ternary"
)

type TryteCore[T Unsigned] struct {
	n int
}

// n - number of trits
func MakeTryteCore[T Unsigned](n int) TryteCore[T] {
	bits := bitsPerUnsigned[T]()
	maxTrits := bits / 2
	if n > maxTrits {
		panic("invalid number of trits")
	}
	return TryteCore[T]{n}
}

// TotalTrits returns the total number of trits.
func (tc TryteCore[T]) TotalTrits() int {
	return tc.n
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) setTrit(x T, i int, t int) T {
	return setTrit(x, i, t)
}

func (tc TryteCore[T]) getTrit(x T, i int) int {
	return getTrit(x, i)
}

//------------------------------------------------------------------------------

// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (tc TryteCore[T]) Compare(a, b T) int {
	for i := tc.n; i > 0; {
		i--
		var (
			ta = tc.getTrit(a, i)
			tb = tc.getTrit(b, i)
		)
		c := tritsCompare(ta, tb)
		if c != 0 {
			return c
		}
	}
	return 0
}

// a == n
func (tc TryteCore[T]) Equal(a, b T) bool {
	return tc.Compare(a, b) == 0
}

// a < b
func (tc TryteCore[T]) Less(a, b T) bool {
	return tc.Compare(a, b) == -1
}

// a > b
func (tc TryteCore[T]) Greater(a, b T) bool {
	return tc.Compare(a, b) == 1
}

// The "less than or equal to" sign: <=
// a <= b
func (tc TryteCore[T]) LessOrEqual(a, b T) bool {
	c := tc.Compare(a, b)
	return (c == -1) || (c == 0)
}

// The "greater than or equal to" sign: >=
// a >= b
func (tc TryteCore[T]) GreaterOrEqual(a, b T) bool {
	c := tc.Compare(a, b)
	return (c == 1) || (c == 0)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Sign(x T) int {
	for i := tc.n; i > 0; {
		i--
		t := tc.getTrit(x, i)
		if t != 0 {
			return t
		}
	}
	return 0
}

// x == 0
func (tc TryteCore[T]) IsZero(x T) bool {
	return tc.Sign(x) == 0
}

// x < 0
func (tc TryteCore[T]) IsNegative(x T) bool {
	return tc.Sign(x) == -1
}

// x > 0
func (tc TryteCore[T]) IsPositive(x T) bool {
	return tc.Sign(x) == 1
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) SetAllTrits(t int) T {
	var a T
	for i := 0; i < tc.n; i++ {
		a = tc.setTrit(a, i, t)
	}
	return a
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) FromIntRest(v int) (a T, rest int) {
	for i := 0; i < tc.n; i++ {
		var t, rem int
		switch {
		case v == 0:
			t = 0
		case v > 0:
			v, rem = quoRemInt(v+1, base)
			t = rem - 1 // rem: {0, 1, 2} -> (rem - 1): {-1, 0, 1}
		case v < 0:
			v, rem = quoRemInt(v-1, base)
			t = rem + 1 // rem: {0,-1,-2} -> (rem + 1): {-1, 0, 1}
		}
		a = tc.setTrit(a, i, t)
	}
	rest = v
	return
}

func (tc TryteCore[T]) FromInt(v int) T {
	a, _ := tc.FromIntRest(v)
	return a
}

func (tc TryteCore[T]) ToInt(a T) int {
	var v int
	p := 1
	for i := 0; i < tc.n; i++ {
		v += p * tc.getTrit(a, i)
		p *= base
	}
	return v
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Format(a T) string {
	var (
		bs = make([]byte, tc.n)
		j  = tc.n - 1
		k  = j
	)
	for i := 0; i < tc.n; i++ {
		t := tc.getTrit(a, i)
		if t != 0 {
			k = j
		}
		bs[j] = mustTritToChar(t)
		j--
	}
	return string(bs[k:])
}

func (tc TryteCore[T]) FormatAllTrits(a T) string {
	var (
		bs = make([]byte, tc.n)
		j  = tc.n - 1
	)
	for i := 0; i < tc.n; i++ {
		t := tc.getTrit(a, i)
		bs[j] = mustTritToChar(t)
		j--
	}
	return string(bs)
}

func (tc TryteCore[T]) Parse(s string) (T, error) {
	var v T
	var count int
	bs := []byte(s)
	for _, b := range bs {
		if b == '_' {
			continue
		}
		t, ok := charToTrit(b)
		if !ok {
			return v, fmt.Errorf("invalid trit char %q", b)
		}
		v = tc.Shl(v, 1)        // v = v << 1
		v = tc.setTrit(v, 0, t) // v[0] = t
		count++
	}
	if not(valueIn(count, 1, tc.n+1)) {
		return v, fmt.Errorf("invalid number of trits: have %d, want [%d..%d]", count, 1, tc.n)
	}
	return v, nil
}

func (tc TryteCore[T]) MustParse(s string) T {
	a, err := tc.Parse(s)
	if err != nil {
		panic(err)
	}
	return a
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Invert(a T) (b T) {
	for i := 0; i < tc.n; i++ {
		t := tc.getTrit(a, i)
		t = ternary.Neg(t)
		b = tc.setTrit(b, i, t)
	}
	return b
}

//------------------------------------------------------------------------------

// Shl - shift left
// a << i
func (tc TryteCore[T]) Shl(a T, i int) T {
	return shiftLeft(tc, a, i)
}

// Shr - shift right
// a >> i
func (tc TryteCore[T]) Shr(a T, i int) T {
	return shiftRight(tc, a, i)
}

//------------------------------------------------------------------------------

// carryIn  - input carry trit
// carryOut - output carry trit

func (tc TryteCore[T]) Add(x, y T, carryIn int) (res T, carryOut int) {
	var s int
	carry := carryIn

	for i := 0; i < tc.n; i++ {
		s, carry = tritsAdd(tc.getTrit(x, i), tc.getTrit(y, i), carry)
		res = tc.setTrit(res, i, s)
	}
	return res, carry
}

func (tc TryteCore[T]) Sub(x, y T, carryIn int) (res T, carryOut int) {
	var s int
	carry := carryIn

	for i := 0; i < tc.n; i++ {
		s, carry = tritsSub(tc.getTrit(x, i), tc.getTrit(y, i), carry)
		res = tc.setTrit(res, i, s)
	}
	return res, carry
}

func (tc TryteCore[T]) Mul(a, b T) (hi, lo T) {
	for i := 0; i < tc.n; i++ {
		ai := tc.getTrit(a, i)
		var w T
		for j := 0; j < tc.n; j++ {
			bj := tc.getTrit(b, j)
			w = tc.setTrit(w, j, tritsMul(ai, bj))
		}
		var (
			wLo   = tc.Shl(w, i)          // w << i
			wHi   = tc.Shr(w, (tc.n - i)) // w >> (n-i)
			carry = 0
		)
		lo, carry = tc.Add(lo, wLo, carry)
		hi, carry = tc.Add(hi, wHi, carry)
	}
	return hi, lo
}

func (tc TryteCore[T]) MulLo(a, b T) (lo T) {
	for i := 0; i < tc.n; i++ {
		ai := tc.getTrit(a, i)
		var w T
		for j := 0; j < tc.n; j++ {
			bj := tc.getTrit(b, j)
			w = tc.setTrit(w, j, tritsMul(ai, bj))
		}
		var (
			wLo   = tc.Shl(w, i) // w << i
			carry = 0
		)
		lo, carry = tc.Add(lo, wLo, carry)
	}
	return lo
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) QuoRem(x, y T) (quo, rem T) {
	return tryteQuoRem(tc, x, y)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Rand(r *Rand) T {
	var a T
	for i := 0; i < tc.n; i++ {
		a = tc.setTrit(a, i, randTrit(r))
	}
	return a
}

func (tc TryteCore[T]) RandSh(r *Rand) T {
	a := tc.Rand(r)
	return tc.Shr(a, r.Intn(tc.n))
}

//------------------------------------------------------------------------------

// big.BitLen(), TritLen()

// Len returns the minimum number of trits required to represent x; the result is 0 for x == 0.
func (tc TryteCore[T]) Len(x T) int {
	for i := tc.n; i > 0; i-- {
		t := tc.getTrit(x, i-1)
		if t != 0 {
			return i
		}
	}
	return 0
}

func (tc TryteCore[T]) Bounds() (min, max int) {
	return tryteBounds(tc.TotalTrits())
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) DoUnary(a T, f ternary.UnaryFunc) T {
	var b T
	for i := 0; i < tc.n; i++ {
		var (
			ai = tc.getTrit(a, i)
			t  = f(ai)
		)
		b = tc.setTrit(b, i, t)
	}
	return b
}

func (tc TryteCore[T]) DoBinary(a, b T, f ternary.BinaryFunc) T {
	var c T
	for i := 0; i < tc.n; i++ {
		var (
			ai = tc.getTrit(a, i)
			bi = tc.getTrit(b, i)
			t  = f(ai, bi)
		)
		c = tc.setTrit(c, i, t)
	}
	return c
}

//------------------------------------------------------------------------------
