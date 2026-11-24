package bal3

import (
	"fmt"
)

// https://rosettacode.org/wiki/Balanced_ternary

// Balanced ternary
type BigTryte struct {
	words []word
}

func NewBigTryte() *BigTryte {
	return &BigTryte{}
}

func NewBigTryteInt64(v int64) *BigTryte {

	b := NewBigTryte()

	var t int64

	i := 0
	for v != 0 {
		v, t = quoRemBal3(v)

		b.setLen(i + 1)
		b.setTrit(i, Trit(t))
		i++
	}

	return b
}

func (b *BigTryte) setTrit(i int, t Trit) {
	wordIndex, tritIndex := quoRem(i, tritsPerWord)
	b.words[wordIndex] = setTrit(b.words[wordIndex], tritIndex, t)
}

func (b *BigTryte) getTrit(i int) Trit {
	wordIndex, tritIndex := quoRem(i, tritsPerWord)
	return getTrit(b.words[wordIndex], tritIndex)
}

func (b *BigTryte) setLen(n int) {
	var (
		wn      = ceilDiv(n, tritsPerWord)
		prevLen = len(b.words)
	)
	if cap(b.words) < wn {
		newWords := make([]word, wn)
		copy(newWords, b.words)
		b.words = newWords
	} else if wn < prevLen {
		clear(b.words[wn:prevLen])
	}
	b.words = b.words[:wn]
}

func (b *BigTryte) backward(f TritIterateFunc) {

	ws := b.words

	// words backward
	for wi := len(ws); wi > 0; { // backward iterate
		wi--

		w := ws[wi]

		// trits backward
		for ti := tritsPerWord; ti > 0; { // backward iterate
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
func (b *BigTryte) TritLen() int {

	ws := b.words

	for wi := len(ws); wi > 0; { // backward iterate
		wi--

		w := ws[wi]

		for ti := tritsPerWord; ti > 0; { // backward iterate
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

func (b *BigTryte) Sign() int {
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
func (b *BigTryte) IsNegative() bool {
	return b.Sign() == -1
}

// x == 0
func (b *BigTryte) IsZero() bool {
	return b.Sign() == 0
}

// x > 0
func (b *BigTryte) IsPositive() bool {
	return b.Sign() == +1
}

//------------------------------------------------------------------------------

func (p *BigTryte) Inverse() *BigTryte {
	ws := p.words
	for i, w := range ws {
		var v word
		for j := 0; j < tritsPerWord; j++ {
			t := getTrit(w, j)
			t = tritInverse(t)
			v = setTrit(v, j, t)
		}
		ws[i] = v
	}
	return p
}

//------------------------------------------------------------------------------

func (b *BigTryte) tryteFormat() (string, error) {

	n := b.TritLen()
	if n == 0 {
		return "0", nil
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
		char, err := tritToChar(t)
		if err != nil {
			return "", fmt.Errorf("trit[%d] error: %w", i, err)
		}

		bs[j] = char
		j--
	}
	return string(bs[k:]), nil
}

func (b *BigTryte) Format() string {
	s, err := b.tryteFormat()
	if err != nil {
		panic(err)
	}
	return s
}

func (b *BigTryte) String() string {
	return b.Format()
}

//------------------------------------------------------------------------------

func ParseBigTryte(s string) (*BigTryte, error) {
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

	b := new(BigTryte)
	b.setLen(len(ts))
	for i, t := range ts {
		b.setTrit(i, t)
	}

	return b, nil
}

//------------------------------------------------------------------------------

func bigTryteToBigInt(b *BigTryte, rest *BigInt) *BigInt {

	v := rest
	if v == nil {
		v = newBigInt(0)
	}

	b.backward(func(i int, t Trit) bool {

		index := tritToIndex(t)
		bt := bigAllTrits[index]

		v.Mul(v, bigBase)
		v.Add(v, bt)

		return true
	})

	return v
}

func (b *BigTryte) ToBigInt() *BigInt {
	return bigTryteToBigInt(b, nil)
}

func (b *BigTryte) ToInt64Err() (int64, error) {
	x := b.ToBigInt()
	return bigIntToInt64(x)
}

func (b *BigTryte) ToInt64Bool() (int64, bool) {
	x := b.ToBigInt()
	y, err := bigIntToInt64(x)
	if err != nil {
		return 0, false
	}
	return y, true
}

func (b *BigTryte) ToInt64() int64 {
	y, _ := b.ToInt64Bool()
	return y
}

//------------------------------------------------------------------------------

func (z *BigTryte) Add(x, y *BigTryte) *BigTryte {
	var (
		xn = len(x.words)
		yn = len(y.words)
		n  = max(xn, yn)

		carry Trit
	)

	z.setLen(n)

	for i := 0; i < n; i++ {
		var wx, wy word
		if i < xn {
			wx = x.words[i]
		}
		if i < yn {
			wy = y.words[i]
		}

		z.words[i], carry = trytesAdd(tritsPerWord, wx, wy, carry)
	}

	if carry != 0 {
		z.setLen(n + 1)
		z.words[n] = setTrit[word](0, 0, carry)
	}

	return z
}

func (z *BigTryte) Sub(x, y *BigTryte) *BigTryte {
	var (
		xn = len(x.words)
		yn = len(y.words)
		n  = max(xn, yn)

		borrow Trit
	)

	z.setLen(n)

	for i := 0; i < n; i++ {
		var wx, wy word
		if i < xn {
			wx = x.words[i]
		}
		if i < yn {
			wy = y.words[i]
		}

		z.words[i], borrow = trytesSub(tritsPerWord, wx, wy, borrow)
	}

	if borrow != 0 {
		z.setLen(n + 1)
		z.words[n] = setTrit[word](0, 0, borrow)
	}

	return z
}
