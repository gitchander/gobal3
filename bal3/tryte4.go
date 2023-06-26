package bal3

// 4 trits
type Tryte4 uint8

var t4c = MakeTryteCore[Tryte4](4)
var T4C = t4c

func (a Tryte4) Int() int {
	return t4c.ToInt(a)
}

func (a Tryte4) String() string {
	return t4c.Format(a)
}

func (a Tryte4) Invert() (b Tryte4) {
	return t4c.Invert(a)
}

func (a Tryte4) Add(b Tryte4) (c Tryte4) {
	sum, _ := t4c.Add(a, b, 0)
	return sum
}

func (a Tryte4) Sub(b Tryte4) (c Tryte4) {
	sum, _ := t4c.Sub(a, b, 0)
	return sum
}

func (a Tryte4) Mul(b Tryte4) Tryte4 {
	_, lo := t4c.Mul(a, b)
	return lo
}

func (a Tryte4) Div(b Tryte4) Tryte4 {
	quo, _ := t4c.QuoRem(a, b)
	return quo
}

func (a Tryte4) Compare(b Tryte4) int {
	return t4c.Compare(a, b)
}

func (a Tryte4) Equal(b Tryte4) bool {
	return t4c.Equal(a, b)
}

func (a Tryte4) Less(b Tryte4) bool {
	return t4c.Less(a, b)
}

func (a Tryte4) Shl(i int) Tryte4 {
	return t4c.Shl(a, i)
}

func (a Tryte4) Shr(i int) Tryte4 {
	return t4c.Shr(a, i)
}

func (a Tryte4) IsZero() bool {
	return t4c.IsZero(a)
}
