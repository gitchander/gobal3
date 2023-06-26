package bal3

// 8 trits
type Tryte8 uint16

var t8c = MakeTryteCore[Tryte8](8)
var T8C = t8c

func (a Tryte8) Int() int {
	return t8c.ToInt(a)
}

func (a Tryte8) String() string {
	return t8c.Format(a)
}

func (a Tryte8) Invert() (b Tryte8) {
	return t8c.Invert(a)
}

func (a Tryte8) Add(b Tryte8) (c Tryte8) {
	sum, _ := t8c.Add(a, b, 0)
	return sum
}

func (a Tryte8) Sub(b Tryte8) (c Tryte8) {
	sum, _ := t8c.Sub(a, b, 0)
	return sum
}

func (a Tryte8) Mul(b Tryte8) Tryte8 {
	_, lo := t8c.Mul(a, b)
	return lo
}

func (a Tryte8) Div(b Tryte8) Tryte8 {
	quo, _ := t8c.QuoRem(a, b)
	return quo
}

func (a Tryte8) Compare(b Tryte8) int {
	return t8c.Compare(a, b)
}

func (a Tryte8) Equal(b Tryte8) bool {
	return t8c.Equal(a, b)
}

func (a Tryte8) Less(b Tryte8) bool {
	return t8c.Less(a, b)
}

func (a Tryte8) Shl(i int) Tryte8 {
	return t8c.Shl(a, i)
}

func (a Tryte8) Shr(i int) Tryte8 {
	return t8c.Shr(a, i)
}

func (a Tryte8) IsZero() bool {
	return t8c.IsZero(a)
}
