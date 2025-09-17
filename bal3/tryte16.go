package bal3

// 16 trits
type Tryte16 uint32

var tc16 = MakeTryteCore[Tryte16](16)
var TC16 = tc16

func (a Tryte16) ToInt64() int64 {
	return tc16.ToInt64(a)
}

func (a Tryte16) String() string {
	return tc16.Format(a)
}

func (a Tryte16) Neg() (b Tryte16) {
	return tc16.Neg(a)
}

func (a Tryte16) Add(b Tryte16) (c Tryte16) {
	sum, _ := tc16.Add(a, b, 0)
	return sum
}

func (a Tryte16) Sub(b Tryte16) (c Tryte16) {
	sum, _ := tc16.Sub(a, b, 0)
	return sum
}

func (a Tryte16) Mul(b Tryte16) Tryte16 {
	_, lo := tc16.Mul(a, b)
	return lo
}

func (a Tryte16) Div(b Tryte16) Tryte16 {
	quo, _ := tc16.QuoRem(a, b)
	return quo
}

func (a Tryte16) Sign() int {
	return tc16.Sign(a)
}

func (a Tryte16) Compare(b Tryte16) int {
	return tc16.Compare(a, b)
}

func (a Tryte16) Equal(b Tryte16) bool {
	return tc16.Equal(a, b)
}

func (a Tryte16) Less(b Tryte16) bool {
	return tc16.Less(a, b)
}

func (a Tryte16) Shl(i int) Tryte16 {
	return tc16.Shl(a, i)
}

func (a Tryte16) Shr(i int) Tryte16 {
	return tc16.Shr(a, i)
}

func (a Tryte16) IsZero() bool {
	return tc16.IsZero(a)
}

func (a Tryte16) ToBigInt() *BigInt {
	return tc16.ToBigInt(a)
}
