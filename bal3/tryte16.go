package bal3

// 16 trits
type Tryte16 uint32

var t16c = MakeTryteCore[Tryte16](16)
var T16C = t16c

func (a Tryte16) Int() int {
	return t16c.ToInt(a)
}

func (a Tryte16) String() string {
	return t16c.Format(a)
}

func (a Tryte16) Invert() (b Tryte16) {
	return t16c.Invert(a)
}

func (a Tryte16) Add(b Tryte16) (c Tryte16) {
	sum, _ := t16c.Add(a, b, 0)
	return sum
}

func (a Tryte16) Sub(b Tryte16) (c Tryte16) {
	sum, _ := t16c.Sub(a, b, 0)
	return sum
}

func (a Tryte16) Mul(b Tryte16) Tryte16 {
	_, lo := t16c.Mul(a, b)
	return lo
}

func (a Tryte16) Div(b Tryte16) Tryte16 {
	quo, _ := t16c.QuoRem(a, b)
	return quo
}

func (a Tryte16) Compare(b Tryte16) int {
	return t16c.Compare(a, b)
}

func (a Tryte16) Equal(b Tryte16) bool {
	return t16c.Equal(a, b)
}

func (a Tryte16) Less(b Tryte16) bool {
	return t16c.Less(a, b)
}

func (a Tryte16) Shl(i int) Tryte16 {
	return t16c.Shl(a, i)
}

func (a Tryte16) Shr(i int) Tryte16 {
	return t16c.Shr(a, i)
}

func (a Tryte16) IsZero() bool {
	return t16c.IsZero(a)
}
