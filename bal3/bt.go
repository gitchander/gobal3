package bal3

// https://rosettacode.org/wiki/Balanced_ternary

// Balanced ternary
type Bt struct {
	words []word
}

func NewBt() *Bt {
	return &Bt{}
}

func (b *Bt) setTrit(i int, t Trit) {
	wordIndex, tritIndex := quoRem(i, tritsPerWord)
	b.words[wordIndex] = setTrit(b.words[wordIndex], tritIndex, t)
}

func (b *Bt) getTrit(i int) Trit {
	wordIndex, tritIndex := quoRem(i, tritsPerWord)
	return getTrit(b.words[wordIndex], tritIndex)
}

func (b *Bt) backward(f func(i int, t Trit) bool) {

	ws := b.words

	// words backward
	for wi := len(ws); wi > 0; {
		wi--

		w := ws[wi]

		// trits backward
		for ti := tritsPerWord; ti > 0; {
			ti--

			var (
				i = (wi * tritsPerWord) + ti
				t = getTrit(w, ti)
			)

			if !(f(i, t)) {
				return
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

			t := getTrit(w, ti)
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

// x < 0
func (b *Bt) IsNegative() bool {
	return b.Sign() == -1
}

// x == 0
func (b *Bt) IsZero() bool {
	return b.Sign() == 0
}

// x > 0
func (b *Bt) IsPositive() bool {
	return b.Sign() == +1
}

//------------------------------------------------------------------------------

func (p *Bt) Neg() *Bt {
	ws := p.words
	for i, w := range ws {
		var v word
		for j := 0; j < tritsPerWord; j++ {
			t := getTrit(w, j)
			t = tritNeg(t)
			v = setTrit(v, j, t)
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

	n := b.TritLen()
	if n == 0 {
		n = 1
	}

	var (
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
