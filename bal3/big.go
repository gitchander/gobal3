package bal3

const tritsPerWord = 8 / bitsPerTrit

type word uint64

// Tryte Big
type Big struct {
	n     int // number of trits
	words []uint64
}

func NewBig() *Big {
	return &Big{}
}

func (b *Big) setTrit(i int, t Trit) {
	wordIndex, tritIndex := quoRemInt(i, tritsPerWord)
	b.words[wordIndex] = setTrit(b.words[wordIndex], tritIndex, t)
}

func (b *Big) getTrit(i int) Trit {
	wordIndex, tritIndex := quoRemInt(i, tritsPerWord)
	return getTrit(b.words[wordIndex], tritIndex)
}

// TritLen returns the length of the absolute value of b in trits. The trit length of 0 is 0.
func (b *Big) TritLen() int {
	return b.n
}

//------------------------------------------------------------------------------

func (b *Big) Sign() int {
	for i := b.n; i > 0; {
		i--
		t := b.getTrit(i)
		switch {
		case t < 0:
			return -1
		case t > 0:
			return +1
		}
	}
	return 0
}

// x == 0
func (b *Big) IsZero() bool {
	return b.Sign() == 0
}

// x < 0
func (b *Big) IsNegative() bool {
	return b.Sign() == -1
}

// x > 0
func (b *Big) IsPositive() bool {
	return b.Sign() == 1
}

//------------------------------------------------------------------------------

func (b *Big) Format() string {
	var (
		n  = b.n
		bs = make([]byte, n)
		j  = n - 1
		k  = j
	)
	for i := 0; i < n; i++ {
		t := b.getTrit(i)
		if t != 0 {
			k = j
		}
		bs[j] = mustTritToChar(t)
		j--
	}
	return string(bs[k:])
}
