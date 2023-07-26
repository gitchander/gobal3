package bal3

type tryteGeneral[T Unsigned] interface {
	isTryte()

	Invert() T

	ToInt64() int64
	String() string

	Add(T) T
	Sub(T) T
	Mul(T) T
	Div(T) T

	Shl(i int) T
	Shr(i int) T

	Compare(T) int
	Equal(T) bool
	Less(T) bool
	IsZero() bool
}

func (Tryte4) isTryte()  {}
func (Tryte8) isTryte()  {}
func (Tryte16) isTryte() {}
func (Tryte32) isTryte() {}
func (Tryte6) isTryte()  {}
func (Tryte9) isTryte()  {}

var (
	_ tryteGeneral[Tryte4]  = Tryte4(0)
	_ tryteGeneral[Tryte8]  = Tryte8(0)
	_ tryteGeneral[Tryte16] = Tryte16(0)
	_ tryteGeneral[Tryte32] = Tryte32(0)

	_ tryteGeneral[Tryte6] = Tryte6(0)
	_ tryteGeneral[Tryte9] = Tryte9(0)
)
