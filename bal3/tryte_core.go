package bal3

import (
	"fmt"

	ivl "github.com/gitchander/gobal3/utils/interval"
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

// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (tc TryteCore[T]) Compare(a, b T) int {
	for i := tc.n; i > 0; {
		i--
		var (
			ta = getTrit(a, i)
			tb = getTrit(b, i)
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

//  +-
//  | -1: x < 0
// -+  0: x = 0
//  | +1: x > 0
//  +-

func (tc TryteCore[T]) Sign(x T) int {
	for i := tc.n; i > 0; {
		i--
		t := getTrit(x, i)
		switch {
		case t < 0:
			return -1
		case t > 0:
			return +1
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

func (tc TryteCore[T]) SetAllTrits(t Trit) T {
	var a T
	for i := 0; i < tc.n; i++ {
		a = setTrit(a, i, t)
	}
	return a
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Int64ToTrite(v int64) (a T, rest int64) {
	return int64ToTrite(tc, v)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) TryteToInt64(a T, rest int64) (int64, bool) {
	return tryteToInt64(tc, a, rest)
}

func (tc TryteCore[T]) ToInt64(a T) int64 {
	x, _ := tc.TryteToInt64(a, 0)
	return x
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Format(a T) string {
	var (
		bs = make([]byte, tc.n)
		j  = tc.n - 1
		k  = j
	)
	for i := 0; i < tc.n; i++ {
		t := getTrit(a, i)
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
		t := getTrit(a, i)
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
		t, err := charToTrit(b)
		if err != nil {
			return v, err
		}
		v = tc.Shl(v, 1)     // v = v << 1
		v = setTrit(v, 0, t) // v[0] = t
		count++
	}
	if l := ivl.Ivl(1, tc.n+1); not(l.Contains(count)) {
		return v, fmt.Errorf("invalid number of trits: have %d, want %v", count, l)
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

// Negative, Invert
func (tc TryteCore[T]) Neg(a T) (b T) {
	for i := 0; i < tc.n; i++ {
		t := getTrit(a, i)
		t = terNeg(t)
		b = setTrit(b, i, t)
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

func (tc TryteCore[T]) Add(x, y T, carryIn Trit) (res T, carryOut Trit) {
	var (
		s     Trit
		carry Trit = carryIn
	)
	for i := 0; i < tc.n; i++ {
		s, carry = tritsAdd(getTrit(x, i), getTrit(y, i), carry)
		res = setTrit(res, i, s)
	}
	return res, carry
}

func (tc TryteCore[T]) Sub(x, y T, carryIn Trit) (res T, carryOut Trit) {
	var (
		s     Trit
		carry Trit = carryIn
	)
	for i := 0; i < tc.n; i++ {
		s, carry = tritsSub(getTrit(x, i), getTrit(y, i), carry)
		res = setTrit(res, i, s)
	}
	return res, carry
}

func (tc TryteCore[T]) Mul(a, b T) (hi, lo T) {
	for i := 0; i < tc.n; i++ {
		ai := getTrit(a, i)
		var w T
		for j := 0; j < tc.n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}
		var (
			wLo = tc.Shl(w, i)          // w << i
			wHi = tc.Shr(w, (tc.n - i)) // w >> (n-i)

			carry Trit = 0
		)
		lo, carry = tc.Add(lo, wLo, carry)
		hi, carry = tc.Add(hi, wHi, carry)
	}
	return hi, lo
}

func (tc TryteCore[T]) MulLo(a, b T) (lo T) {
	for i := 0; i < tc.n; i++ {
		ai := getTrit(a, i)
		var w T
		for j := 0; j < tc.n; j++ {
			bj := getTrit(b, j)
			w = setTrit(w, j, tritsMul(ai, bj))
		}
		var (
			wLo = tc.Shl(w, i) // w << i

			carry Trit = 0
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
		a = setTrit(a, i, randTrit(r))
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
		t := getTrit(x, i-1)
		if t != 0 {
			return i
		}
	}
	return 0
}

// Bounds
func (tc TryteCore[T]) Limits() (min, max T) {
	n := tc.TotalTrits()
	min = setTritsN[T](n, tv_T)
	max = setTritsN[T](n, tv_1)
	return min, max
}

func (tc TryteCore[T]) LimitsInt64() (min, max int64) {
	tmin, tmax := tc.Limits()
	min = tc.ToInt64(tmin)
	max = tc.ToInt64(tmax)
	return min, max
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) DoUnary(a T, f UnaryFunc) T {
	var b T
	for i := 0; i < tc.n; i++ {
		var (
			ai = getTrit(a, i)
			bi = f(ai)
		)
		b = setTrit(b, i, bi)
	}
	return b
}

func (tc TryteCore[T]) DoBinary(a, b T, f BinaryFunc) T {
	var c T
	for i := 0; i < tc.n; i++ {
		var (
			ai = getTrit(a, i)
			bi = getTrit(b, i)
			ci = f(ai, bi)
		)
		c = setTrit(c, i, ci)
	}
	return c
}

//------------------------------------------------------------------------------
