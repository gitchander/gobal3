package bal3

// 4 trits
type Tryte4 uint8

var tc4 = MakeTryteCore[Tryte4](4)
var TC4 = tc4

func (a Tryte4) ToInt64() (int64, bool) {
	return tc4.TryteToInt64(a, 0)
}

func (a Tryte4) String() string {
	return tc4.Format(a)
}

func (a Tryte4) Invert() (b Tryte4) {
	return tc4.Invert(a)
}

func (a Tryte4) Add(b Tryte4) (c Tryte4) {
	sum, _ := tc4.Add(a, b, 0)
	return sum
}

func (a Tryte4) Sub(b Tryte4) (c Tryte4) {
	sum, _ := tc4.Sub(a, b, 0)
	return sum
}

func (a Tryte4) Mul(b Tryte4) Tryte4 {
	_, lo := tc4.Mul(a, b)
	return lo
}

func (a Tryte4) Div(b Tryte4) Tryte4 {
	quo, _ := tc4.QuoRem(a, b)
	return quo
}

func (a Tryte4) Compare(b Tryte4) int {
	return tc4.Compare(a, b)
}

func (a Tryte4) Equal(b Tryte4) bool {
	return tc4.Equal(a, b)
}

func (a Tryte4) Less(b Tryte4) bool {
	return tc4.Less(a, b)
}

func (a Tryte4) Shl(i int) Tryte4 {
	return tc4.Shl(a, i)
}

func (a Tryte4) Shr(i int) Tryte4 {
	return tc4.Shr(a, i)
}

func (a Tryte4) IsZero() bool {
	return tc4.IsZero(a)
}
