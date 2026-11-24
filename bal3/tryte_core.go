package bal3

type TryteCore[T GenericTryte] struct {
	n int // Trits per tryte
}

func MakeTryteCore[T GenericTryte](n int) TryteCore[T] {
	bits := bitsPerUnsigned[T]()
	maxTrits := bits / 2
	if (n < 0) || (maxTrits < n) {
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

// a == b
func (tc TryteCore[T]) Equal(a, b T) bool {
	return tc.Compare(a, b) == 0
}

// a < b
func (tc TryteCore[T]) Less(a, b T) bool {
	return tc.Compare(a, b) == -1
}

// a > b
func (tc TryteCore[T]) Greater(a, b T) bool {
	return tc.Compare(a, b) == +1
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
	return tc.Sign(x) == +1
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) SetAllTrits(t Trit) T {
	return tryteSetAllTrits[T](tc.n, t)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Int64ToTriteRest(v int64) (a T, rest int64) {
	return int64ToTriteRest(tc, v)
}

func (tc TryteCore[T]) TryteRestToInt64(a T, rest int64) (int64, bool) {
	return tryteRestToInt64(tc, a, rest)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Int64ToTrite(v int64) T {
	t, _ := tc.Int64ToTriteRest(v)
	return t
}

func (tc TryteCore[T]) TryteToInt64(a T) int64 {
	v, ok := tc.TryteRestToInt64(a, 0)
	if !ok {
		panic("TryteToInt64")
	}
	return v
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) formatV1(a T) string {
	s, err := tryteFormat(tc.n, a)
	if err != nil {
		panic(err)
	}
	return s
}

func (tc TryteCore[T]) FormatAllTrits(a T) string {
	s, err := tryteFormatAllTrits(tc.n, a)
	if err != nil {
		panic(err)
	}
	return s
}

func (tc TryteCore[T]) Format(a T) string {
	// return tc.formatV1(a)
	return tc.FormatAllTrits(a)
}

func (tc TryteCore[T]) Parse(s string) (T, error) {
	return tryteParse[T](tc.n, s)
}

func (tc TryteCore[T]) MustParse(s string) T {
	a, err := tc.Parse(s)
	if err != nil {
		panic(err)
	}
	return a
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Inverse(a T) (b T) {
	return tryteInverse(tc.n, a)
}

//------------------------------------------------------------------------------

// Shl - shift left
// a << i
func (tc TryteCore[T]) Shl(a T, i int) T {
	return tryteShl(tc.n, a, i)
}

// Shr - shift right
// a >> i
func (tc TryteCore[T]) Shr(a T, i int) T {
	return tryteShr(tc.n, a, i)
}

//------------------------------------------------------------------------------

// sum + (3^n) * c1 = x + y + c0

func (tc TryteCore[T]) Add(x, y T, c0 Trit) (sum T, c1 Trit) {
	return trytesAdd(tc.n, x, y, c0)
}

//------------------------------------------------------------------------------

// diff - (3^n) * b1 = x - y - b0

func (tc TryteCore[T]) Sub(x, y T, b0 Trit) (diff T, b1 Trit) {
	return trytesSub(tc.n, x, y, b0)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Mul(a, b T) (hi, lo T) {
	return trytesMul(tc.n, a, b)
}

func (tc TryteCore[T]) MulLo(a, b T) (lo T, overflow bool) {
	return trytesMulLo(tc.n, a, b)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) QuoRem(hi, lo, y T) (quo, rem T) {
	return tryteQuoRem(tc, hi, lo, y)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) Rand(r *Rand) T {
	return randTryte[T](tc.n, r)
}

func (tc TryteCore[T]) RandSh(r *Rand) T {
	return randTryteSh[T](tc.n, r)
}

func (tc TryteCore[T]) randTryteForTest(r *Rand) T {
	return randTryteForTest[T](tc.n, r)
}

//------------------------------------------------------------------------------

// big.BitLen(), TritLen()

// Len returns the minimum number of trits required to represent x; the result is 0 for x == 0.
func (tc TryteCore[T]) Len(x T) int {
	return tryteLen(tc.n, x)
}

//------------------------------------------------------------------------------

func (tc TryteCore[T]) DoUnary(a T, f UnaryFunc) T {
	return tryteDoUnary(tc.n, a, f)
}

func (tc TryteCore[T]) DoBinary(a, b T, f BinaryFunc) T {
	return tryteDoBinary(tc.n, a, b, f)
}

//------------------------------------------------------------------------------

// Bounds
func (tc TryteCore[T]) Limits() (minValue, maxValue T) {
	return tryteLimits[T](tc.n)
}

func (tc TryteCore[T]) LimitsInt64() (minValue, maxValue int64) {
	tmin, tmax := tryteLimits[T](tc.n)
	minValue = tc.TryteToInt64(tmin)
	maxValue = tc.TryteToInt64(tmax)
	return minValue, maxValue
}

func (tc TryteCore[T]) ToBigInt(a T) *BigInt {
	return tryteToBigInt(tc.n, a, nil)
}

//------------------------------------------------------------------------------
