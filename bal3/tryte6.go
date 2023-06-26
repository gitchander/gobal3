package bal3

// 6 trits
type Tryte6 uint16

var t6c = MakeTryteCore[Tryte6](6)
var T6C = t6c

func (a Tryte6) Int() int {
	return t6c.ToInt(a)
}

func (a Tryte6) String() string {
	return t6c.Format(a)
}

func (a Tryte6) Invert() (b Tryte6) {
	return t6c.Invert(a)
}

func (a Tryte6) Add(b Tryte6) (c Tryte6) {
	sum, _ := t6c.Add(a, b, 0)
	return sum
}

func (a Tryte6) Sub(b Tryte6) (c Tryte6) {
	sum, _ := t6c.Sub(a, b, 0)
	return sum
}

func (a Tryte6) Mul(b Tryte6) Tryte6 {
	_, lo := t6c.Mul(a, b)
	return lo
}

func (a Tryte6) Div(b Tryte6) Tryte6 {
	quo, _ := t6c.QuoRem(a, b)
	return quo
}

func (a Tryte6) Compare(b Tryte6) int {
	return t6c.Compare(a, b)
}

func (a Tryte6) Equal(b Tryte6) bool {
	return t6c.Equal(a, b)
}

func (a Tryte6) Less(b Tryte6) bool {
	return t6c.Less(a, b)
}

func (a Tryte6) Shl(i int) Tryte6 {
	return t6c.Shl(a, i)
}

func (a Tryte6) Shr(i int) Tryte6 {
	return t6c.Shr(a, i)
}

func (a Tryte6) IsZero() bool {
	return t6c.IsZero(a)
}
