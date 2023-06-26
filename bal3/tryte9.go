package bal3

// 9 trits
type Tryte9 uint32

var t9c = MakeTryteCore[Tryte9](9)
var T9C = t9c

func (a Tryte9) Int() int {
	return t9c.ToInt(a)
}

func (a Tryte9) String() string {
	return t9c.Format(a)
}

func (a Tryte9) Invert() (b Tryte9) {
	return t9c.Invert(a)
}

func (a Tryte9) Add(b Tryte9) (c Tryte9) {
	sum, _ := t9c.Add(a, b, 0)
	return sum
}

func (a Tryte9) Sub(b Tryte9) (c Tryte9) {
	sum, _ := t9c.Sub(a, b, 0)
	return sum
}

func (a Tryte9) Mul(b Tryte9) Tryte9 {
	_, lo := t9c.Mul(a, b)
	return lo
}

func (a Tryte9) Div(b Tryte9) Tryte9 {
	quo, _ := t9c.QuoRem(a, b)
	return quo
}

func (a Tryte9) Compare(b Tryte9) int {
	return t9c.Compare(a, b)
}

func (a Tryte9) Equal(b Tryte9) bool {
	return t9c.Equal(a, b)
}

func (a Tryte9) Less(b Tryte9) bool {
	return t9c.Less(a, b)
}

func (a Tryte9) Shl(i int) Tryte9 {
	return t9c.Shl(a, i)
}

func (a Tryte9) Shr(i int) Tryte9 {
	return t9c.Shr(a, i)
}

func (a Tryte9) IsZero() bool {
	return t9c.IsZero(a)
}
