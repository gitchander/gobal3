package bal3

// 8 trits
type Tryte8 uint16

var tc8 = MakeTryteCore[Tryte8](8)
var TC8 = tc8

func (a Tryte8) Int() int {
	return tc8.TryteToInt(a)
}

func (a Tryte8) String() string {
	return tc8.Format(a)
}

func (a Tryte8) Invert() (b Tryte8) {
	return tc8.Invert(a)
}

func (a Tryte8) Add(b Tryte8) (c Tryte8) {
	sum, _ := tc8.Add(a, b, 0)
	return sum
}

func (a Tryte8) Sub(b Tryte8) (c Tryte8) {
	sum, _ := tc8.Sub(a, b, 0)
	return sum
}

func (a Tryte8) Mul(b Tryte8) Tryte8 {
	_, lo := tc8.Mul(a, b)
	return lo
}

func (a Tryte8) Div(b Tryte8) Tryte8 {
	quo, _ := tc8.QuoRem(a, b)
	return quo
}

func (a Tryte8) Compare(b Tryte8) int {
	return tc8.Compare(a, b)
}

func (a Tryte8) Equal(b Tryte8) bool {
	return tc8.Equal(a, b)
}

func (a Tryte8) Less(b Tryte8) bool {
	return tc8.Less(a, b)
}

func (a Tryte8) Shl(i int) Tryte8 {
	return tc8.Shl(a, i)
}

func (a Tryte8) Shr(i int) Tryte8 {
	return tc8.Shr(a, i)
}

func (a Tryte8) IsZero() bool {
	return tc8.IsZero(a)
}
