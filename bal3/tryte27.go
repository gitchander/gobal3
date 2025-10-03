package bal3

// 27 trits
type Tryte27 uint64

var (
	tc27 = MakeTryteCore[Tryte27](27)

	TC27 = tc27
)

func (a Tryte27) ToInt64() int64 {
	return tc27.ToInt64(a)
}

func (a Tryte27) String() string {
	return tc27.Format(a)
}

func (a Tryte27) Inverse() (b Tryte27) {
	return tc27.Inverse(a)
}

func (a Tryte27) Add(b Tryte27) (c Tryte27) {
	sum, _ := tc27.Add(a, b, 0)
	return sum
}

func (a Tryte27) Sub(b Tryte27) (c Tryte27) {
	sum, _ := tc27.Sub(a, b, 0)
	return sum
}

func (a Tryte27) Mul(b Tryte27) Tryte27 {
	_, lo := tc27.Mul(a, b)
	return lo
}

func (a Tryte27) Div(b Tryte27) Tryte27 {
	quo, _ := tc27.QuoRem(a, b)
	return quo
}

func (a Tryte27) Sign() int {
	return tc27.Sign(a)
}

func (a Tryte27) Compare(b Tryte27) int {
	return tc27.Compare(a, b)
}

func (a Tryte27) Equal(b Tryte27) bool {
	return tc27.Equal(a, b)
}

func (a Tryte27) Less(b Tryte27) bool {
	return tc27.Less(a, b)
}

func (a Tryte27) Shl(i int) Tryte27 {
	return tc27.Shl(a, i)
}

func (a Tryte27) Shr(i int) Tryte27 {
	return tc27.Shr(a, i)
}

func (a Tryte27) IsZero() bool {
	return tc27.IsZero(a)
}

func (a Tryte27) ToBigInt() *BigInt {
	return tc27.ToBigInt(a)
}
