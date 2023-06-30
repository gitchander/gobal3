package bal3

// 9 trits
type Tryte9 uint32

var tc9 = MakeTryteCore[Tryte9](9)
var TC9 = tc9

func (a Tryte9) Int() int {
	return tc9.ToInt(a)
}

func (a Tryte9) String() string {
	return tc9.Format(a)
}

func (a Tryte9) Invert() (b Tryte9) {
	return tc9.Invert(a)
}

func (a Tryte9) Add(b Tryte9) (c Tryte9) {
	sum, _ := tc9.Add(a, b, 0)
	return sum
}

func (a Tryte9) Sub(b Tryte9) (c Tryte9) {
	sum, _ := tc9.Sub(a, b, 0)
	return sum
}

func (a Tryte9) Mul(b Tryte9) Tryte9 {
	_, lo := tc9.Mul(a, b)
	return lo
}

func (a Tryte9) Div(b Tryte9) Tryte9 {
	quo, _ := tc9.QuoRem(a, b)
	return quo
}

func (a Tryte9) Compare(b Tryte9) int {
	return tc9.Compare(a, b)
}

func (a Tryte9) Equal(b Tryte9) bool {
	return tc9.Equal(a, b)
}

func (a Tryte9) Less(b Tryte9) bool {
	return tc9.Less(a, b)
}

func (a Tryte9) Shl(i int) Tryte9 {
	return tc9.Shl(a, i)
}

func (a Tryte9) Shr(i int) Tryte9 {
	return tc9.Shr(a, i)
}

func (a Tryte9) IsZero() bool {
	return tc9.IsZero(a)
}
