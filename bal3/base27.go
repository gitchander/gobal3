package bal3

import (
	"errors"
	"fmt"
)

//------------------------------------------------------------------------------

// Negative chars arguments:

// English alphabet:
// ABCDEFGHIJKLMNOPQRSTUVWXYZ

// From 'T' in the opposite direction:
// ABCDEFGHIJKLMNOPQRST

// exclude: IOQ
// ABCDEFGHJKLMNPRST

// Only 13 chars from 'T':
// EFGHJKLMNPRST

// Replace EF -> XY (because "EF" use in hex):
// XYGHJKLMNPRST

const (
	negativeChars = "XYGHJKLMNPRST"
	positiveChars = "123456789ABCD"
)

var coderB27 = mustMakeBase27Coder(negativeChars + "0" + positiveChars)

//------------------------------------------------------------------------------

type base27Coder struct {
	enc [27]byte
	dec [256]byte
}

func makeBase27Coder(s string) (*base27Coder, error) {
	const n = 27
	bs := []byte(s)
	if len(bs) != n {
		return nil, fmt.Errorf("invalid table len: have %d, want %d", len(bs), n)
	}
	// check unique
	if hasDuplicate(bs) {
		return nil, errors.New("table has duplicates")
	}
	var et [n]byte
	var dt [256]byte
	for i, b := range bs {
		et[i] = b
		dt[b] = (byte(i) << 1) | 1
	}

	c := &base27Coder{
		enc: et,
		dec: dt,
	}

	return c, nil
}

func mustMakeBase27Coder(s string) *base27Coder {
	t, err := makeBase27Coder(s)
	if err != nil {
		panic(err)
	}
	return t
}

func (p *base27Coder) digitToChar(digit int) (char byte, ok bool) {
	if inInterval(digit, -13, (13 + 1)) {
		char = p.enc[digit+13]
		return char, true
	}
	return 0, false
}

func (p *base27Coder) charToDigit(char byte) (digit int, ok bool) {
	x := p.dec[char]
	if (x & 1) == 1 {
		digit = int(x>>1) - 13
		return digit, true
	}
	return 0, false
}

//------------------------------------------------------------------------------

func FormatBase27[T Unsigned](tc TryteCore[T], a T) string {

	const tritsPerDigit = 3

	var (
		dn = ceilDiv(tc.n, tritsPerDigit) // number of digits
		cs = make([]byte, dn)

		j = dn - 1
		k = j
	)

	writeDigit := func(digit int) {

		if digit != 0 {
			k = j
		}

		char, ok := coderB27.digitToChar(digit)
		if !ok {
			err := fmt.Errorf("invalid digit %d", digit)
			panic(err)
		}
		cs[j] = char
		j--
	}

	var b T
	count := 0 // count of trits in 'b'.

	for i := 0; i < tc.n; i++ {

		t := tc.getTrit(a, i)       // t = a[i]
		b = tc.setTrit(b, count, t) // b[count] = t
		count++

		if count == tritsPerDigit {
			digit := tc.ToInt(b)
			writeDigit(digit)

			// reset all
			b = 0
			count = 0
		}
	}

	if count > 0 {
		digit := tc.ToInt(b)
		writeDigit(digit)
	}

	return string(cs[k:])
}

func ParseBase27[T Unsigned](tc TryteCore[T], s string) (T, error) {

	const tritsPerDigit = 3

	var a T
	count := 0 // count of trits in 'a'.

	bs := []byte(s)
	for _, char := range bs {
		if char == '_' {
			continue
		}
		digit, ok := coderB27.charToDigit(char)
		if !ok {
			return 0, fmt.Errorf("invalid char %c", char)
		}
		b := tc.FromInt(digit)
		for j := tritsPerDigit; j > 0; j-- {
			if count >= tc.n {
				return 0, fmt.Errorf("number of trits more than %d", tc.n)
			}
			t := tc.getTrit(b, j-1)
			if (count > 0) || (t != 0) {
				count++
			}
			a = tc.Shl(a, 1) // a = a << 1
			a = tc.setTrit(a, 0, t)
		}
	}
	return a, nil
}