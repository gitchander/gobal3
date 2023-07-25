package bal3

// 6 trits
type Tryte6 uint16

var tc6 = MakeTryteCore[Tryte6](6)
var TC6 = tc6

func (a Tryte6) ToInt64() (int64, bool) {
	return tc6.TryteToInt64(a, 0)
}

func (a Tryte6) String() string {
	return tc6.Format(a)
}

func (a Tryte6) Invert() (b Tryte6) {
	return tc6.Invert(a)
}

func (a Tryte6) Add(b Tryte6) (c Tryte6) {
	sum, _ := tc6.Add(a, b, 0)
	return sum
}

func (a Tryte6) Sub(b Tryte6) (c Tryte6) {
	sum, _ := tc6.Sub(a, b, 0)
	return sum
}

func (a Tryte6) Mul(b Tryte6) Tryte6 {
	_, lo := tc6.Mul(a, b)
	return lo
}

func (a Tryte6) Div(b Tryte6) Tryte6 {
	quo, _ := tc6.QuoRem(a, b)
	return quo
}

func (a Tryte6) Compare(b Tryte6) int {
	return tc6.Compare(a, b)
}

func (a Tryte6) Equal(b Tryte6) bool {
	return tc6.Equal(a, b)
}

func (a Tryte6) Less(b Tryte6) bool {
	return tc6.Less(a, b)
}

func (a Tryte6) Shl(i int) Tryte6 {
	return tc6.Shl(a, i)
}

func (a Tryte6) Shr(i int) Tryte6 {
	return tc6.Shr(a, i)
}

func (a Tryte6) IsZero() bool {
	return tc6.IsZero(a)
}
