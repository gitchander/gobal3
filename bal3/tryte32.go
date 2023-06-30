package bal3

// 32 trits
type Tryte32 uint64

var tc32 = MakeTryteCore[Tryte32](32)
var TC32 = tc32

func (a Tryte32) Int() int {
	return tc32.ToInt(a)
}

func (a Tryte32) String() string {
	return tc32.Format(a)
}

func (a Tryte32) Invert() (b Tryte32) {
	return tc32.Invert(a)
}

func (a Tryte32) Add(b Tryte32) (c Tryte32) {
	sum, _ := tc32.Add(a, b, 0)
	return sum
}

func (a Tryte32) Sub(b Tryte32) (c Tryte32) {
	sum, _ := tc32.Sub(a, b, 0)
	return sum
}

func (a Tryte32) Mul(b Tryte32) Tryte32 {
	_, lo := tc32.Mul(a, b)
	return lo
}

func (a Tryte32) Div(b Tryte32) Tryte32 {
	quo, _ := tc32.QuoRem(a, b)
	return quo
}

func (a Tryte32) Compare(b Tryte32) int {
	return tc32.Compare(a, b)
}

func (a Tryte32) Equal(b Tryte32) bool {
	return tc32.Equal(a, b)
}

func (a Tryte32) Less(b Tryte32) bool {
	return tc32.Less(a, b)
}

func (a Tryte32) Shl(i int) Tryte32 {
	return tc32.Shl(a, i)
}

func (a Tryte32) Shr(i int) Tryte32 {
	return tc32.Shr(a, i)
}

func (a Tryte32) IsZero() bool {
	return tc32.IsZero(a)
}
