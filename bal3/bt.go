package bal3

// https://rosettacode.org/wiki/Balanced_ternary

const (
	bitsPerWord  = 64
	tritsPerWord = bitsPerWord / bitsPerTrit
)

type word uint64

// Balanced ternary
type Bt struct {
	words []word
}

func NewBt() *Bt {
	return &Bt{}
}

func (b *Bt) setTrit(i int, t Trit) {
	wordIndex, tritIndex := quoRemInt(i, tritsPerWord)
	b.words[wordIndex] = setTrit(b.words[wordIndex], tritIndex, t)
}

func (b *Bt) getTrit(i int) Trit {
	wordIndex, tritIndex := quoRemInt(i, tritsPerWord)
	return getTrit(b.words[wordIndex], tritIndex)
}

func (b *Bt) backward(f func(i int, t Trit) bool) {
	ws := b.words
	for wi := len(ws); wi > 0; {
		wi--

		w := ws[wi]

		for ti := tritsPerWord; ti > 0; {
			ti--

			var (
				i = wi*tritsPerWord + ti
				t = bitsToTrit(w >> (ti * bitsPerTrit))
			)
			if f(i, t) {
				break
			}
		}
	}
}

// TritLen returns the length of the absolute value of b in trits. The trit length of 0 is 0.
func (b *Bt) TritLen() int {
	// todo
	ws := b.words
	for wi := len(ws); wi > 0; {
		wi--

		w := ws[wi]

		for ti := tritsPerWord; ti > 0; {
			ti--

			k := ti * bitsPerTrit
			t := bitsToTrit(w >> k)
			if t != 0 {
				return wi*tritsPerWord + ti + 1
			}
		}
	}
	return 0
}

//------------------------------------------------------------------------------

func (b *Bt) Sign() int {
	var v int
	b.backward(func(i int, t Trit) bool {
		switch {
		case t < 0:
			v = -1
			return false
		case t > 0:
			v = +1
			return false
		}
		return true
	})
	return v
}

// x == 0
func (b *Bt) IsZero() bool {
	return b.Sign() == 0
}

// x < 0
func (b *Bt) IsNegative() bool {
	return b.Sign() == -1
}

// x > 0
func (b *Bt) IsPositive() bool {
	return b.Sign() == 1
}

//------------------------------------------------------------------------------

func (p *Bt) Neg() *Bt {
	ws := p.words
	for i, w := range ws {
		var v word
		for j := 0; j < tritsPerWord; j++ {
			k := j * bitsPerTrit
			t := bitsToTrit(w >> k)
			t = terNeg(t)
			v |= tritToBits[word](t) << k
		}
		ws[i] = v
	}
	return p
}

//------------------------------------------------------------------------------

func (b *Bt) String() string {
	return b.Format()
}

func (b *Bt) Format() string {
	// todo
	var (
		n  = b.TritLen()
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

func (b *Bt) setLen(n int) {
	wn := ceilDiv(n, tritsPerWord)
	if cap(b.words) < wn {
		b.words = make([]word, wn)
	}
	b.words = b.words[:wn]
}

func ParseBt(s string) (*Bt, error) {
	var (
		chars = []byte(s)
		ts    = make([]Trit, 0, len(chars))
	)
	for _, char := range chars {
		if char == '_' {
			continue
		}
		t, err := charToTrit(char)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	reverseSlice(ts)

	b := new(Bt)
	b.setLen(len(ts))
	for i, t := range ts {
		b.setTrit(i, t)
	}

	return b, nil
}
