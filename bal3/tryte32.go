package bal3

// 32 trits
type Tryte32 uint64

var t32c = MakeTryteCore[Tryte32](32)
var T32C = t32c

func (a Tryte32) Int() int {
	return t32c.ToInt(a)
}

func (a Tryte32) String() string {
	return t32c.Format(a)
}

func (a Tryte32) Invert() (b Tryte32) {
	return t32c.Invert(a)
}

func (a Tryte32) Add(b Tryte32) (c Tryte32) {
	sum, _ := t32c.Add(a, b, 0)
	return sum
}

func (a Tryte32) Sub(b Tryte32) (c Tryte32) {
	sum, _ := t32c.Sub(a, b, 0)
	return sum
}

func (a Tryte32) Mul(b Tryte32) Tryte32 {
	_, lo := t32c.Mul(a, b)
	return lo
}

func (a Tryte32) Div(b Tryte32) Tryte32 {
	quo, _ := t32c.QuoRem(a, b)
	return quo
}

func (a Tryte32) Compare(b Tryte32) int {
	return t32c.Compare(a, b)
}

func (a Tryte32) Equal(b Tryte32) bool {
	return t32c.Equal(a, b)
}

func (a Tryte32) Less(b Tryte32) bool {
	return t32c.Less(a, b)
}

func (a Tryte32) Shl(i int) Tryte32 {
	return t32c.Shl(a, i)
}

func (a Tryte32) Shr(i int) Tryte32 {
	return t32c.Shr(a, i)
}

func (a Tryte32) IsZero() bool {
	return t32c.IsZero(a)
}
