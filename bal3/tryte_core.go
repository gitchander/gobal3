package bal3

import (
	"fmt"

	ivl "github.com/gitchander/gobal3/utils/interval"
)

type TryteCore[T CoreTryte] struct {
	n int
}

// n - number of trits
func MakeTryteCore[T CoreTryte](n int) TryteCore[T] {
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

func (tc TryteCore[T]) GetTrit(a T, i int) Trit {
	return getTrit(a, i)
}

func (tc TryteCore[T]) SetTrit(a T, i int, t Trit) T {
	return setTrit(a, i, t)
}

//------------------------------------------------------------------------------

//            /
//            | -1: a < b
// cmp(a,b) = |  0: a = b
//            | +1: a > b
//            \

// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (tc TryteCore[T]) Compare(a, b T) int {
	return trytesCompare(tc.n, a, b)
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
	return (c == +1) || (c == 0)
}

//------------------------------------------------------------------------------

//           /
//           | -1: x < 0
// sign(x) = |  0: x = 0
//           | +1: x > 0
//           \

func (tc TryteCore[T]) Sign(x T) int {
	return tryteSign(tc.n, x)
}

// x < 0
func (tc TryteCore[T]) IsNegative(x T) bool {
	return tc.Sign(x) == -1
}

// x == 0
func (tc TryteCore[T]) IsZero(x T) bool {
	return tc.Sign(x) == 0
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
	return tryteNeg(tc.n, a)
}

//------------------------------------------------------------------------------

// Shl - shift left
// a << i
func (tc TryteCore[T]) Shl(a T, i int) T {
	return tryteShiftLeft(tc.n, a, i)
}

// Shr - shift right
// a >> i
func (tc TryteCore[T]) Shr(a T, i int) T {
	return tryteShiftRight(tc.n, a, i)
}

//------------------------------------------------------------------------------

// carryIn  - input carry trit
// carryOut - output carry trit

func (tc TryteCore[T]) Add(x, y T, carryIn Trit) (res T, carryOut Trit) {
	return trytesAdd(tc.n, x, y, carryIn)
}

func (tc TryteCore[T]) Sub(x, y T, carryIn Trit) (res T, carryOut Trit) {
	return trytesSub(tc.n, x, y, carryIn)
}

func (tc TryteCore[T]) Mul(a, b T) (hi, lo T) {
	return trytesMul(tc.n, a, b)
}

func (tc TryteCore[T]) MulLo(a, b T) (lo T) {
	return trytesMulLo(tc.n, a, b)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) QuoRem(x, y T) (quo, rem T) {
	return tryteQuoRem(tc, x, y)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Rand(r *Rand) T {
	return randTryte[T](tc.n, r)
}

func (tc TryteCore[T]) RandSh(r *Rand) T {
	return randTryteSh[T](tc.n, r)
}

//------------------------------------------------------------------------------

// big.BitLen(), TritLen()

// Len returns the minimum number of trits required to represent x; the result is 0 for x == 0.
func (tc TryteCore[T]) Len(x T) int {
	for i := tc.n; i > 0; { // backward iterate
		i--

		t := getTrit(x, i)
		if t != 0 {
			return i
		}
	}
	return 0
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

func (tc TryteCore[T]) MinValue() T {
	return setTritsN[T](tc.n, tv_T)
}

func (tc TryteCore[T]) MaxValue() T {
	return setTritsN[T](tc.n, tv_1)
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
