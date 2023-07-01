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

var globalCharCoder = mustMakeCharCoder(negativeChars + "0" + positiveChars)

//------------------------------------------------------------------------------

type charCoder struct {
	enc [27]byte
	dec [256]byte
}

func makeCharCoder(s string) (*charCoder, error) {
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

	c := &charCoder{
		enc: et,
		dec: dt,
	}

	return c, nil
}

func mustMakeCharCoder(s string) *charCoder {
	t, err := makeCharCoder(s)
	if err != nil {
		panic(err)
	}
	return t
}

func (p *charCoder) digitToChar(digit int) (char byte, ok bool) {
	if inInterval(digit, -13, (13 + 1)) {
		char = p.enc[digit+13]
		return char, true
	}
	return 0, false
}

func (p *charCoder) charToDigit(char byte) (digit int, ok bool) {
	x := p.dec[char]
	if (x & 1) == 1 {
		digit = int(x>>1) - 13
		return digit, true
	}
	return 0, false
}

//------------------------------------------------------------------------------

func FormatBase27[T Unsigned](tc TryteCore[T], a T) string {
	var (
		ps = powersOfThree
		dn = ceilDiv(tc.n, 3) // number of digits
		cs = make([]byte, dn)

		j = dn - 1
		k = j
	)
	for i := 0; i < dn; i++ {

		var (
			t0 = tc.getTrit(a, 0)
			t1 = tc.getTrit(a, 1)
			t2 = tc.getTrit(a, 2)
		)
		a = tc.Shr(a, 3)

		digit := t2*ps[2] + t1*ps[1] + t0*ps[0]
		if digit != 0 {
			k = j
		}

		char, ok := globalCharCoder.digitToChar(digit)
		if !ok {
			err := fmt.Errorf("invalid digit %d", digit)
			panic(err)
		}
		cs[j] = char
		j--
	}

	return string(cs[k:])
}

func ParseBase27[T Unsigned](tc TryteCore[T], s string) (T, error) {

	bs := []byte(s)

	var a T
	countTrits := 0

	digits := make([]int, len(bs))

	for i, char := range bs {
		digit, ok := globalCharCoder.charToDigit(char)
		if !ok {
			return 0, fmt.Errorf("invalid char %c", char)
		}
		digits[i] = digit
	}

	for i := len(digits); i > 0; i-- {
		digit := digits[i-1]
		b := tc.FromInt(digit)
		fmt.Println("digit:", digit)

		for j := 3; j > 0; j-- {
			t := tc.getTrit(b, j-1)

			if (t != 0) && (countTrits >= tc.n) {
				fmt.Println(t, countTrits, tc.n)
				return 0, fmt.Errorf("trits are more than need")
			}
			countTrits++

			a = tc.Shl(a, 1)
			a = tc.setTrit(a, 0, t)
		}
	}

	return a, nil
}
