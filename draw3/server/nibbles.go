package main

import (
	"fmt"
)

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/Nibble
func byteToNibbles(b byte) (hi, lo byte) {
	hi = b >> 4
	lo = b & 0xf
	return
}

func nibblesToByte(hi, lo byte) (b byte) {
	b |= hi << 4
	b |= lo & 0xf
	return
}

//------------------------------------------------------------------------------

func decodeNibble(b byte) (byte, bool) {
	if ('0' <= b) && (b <= '9') {
		return b - '0', true
	}
	if ('a' <= b) && (b <= 'f') {
		return b - 'a' + 10, true
	}
	if ('A' <= b) && (b <= 'F') {
		return b - 'A' + 10, true
	}
	return 0, false
}

func decodeNibbles(bs []byte) ([]byte, error) {
	ns := make([]byte, len(bs))
	for i, b := range bs {
		n, ok := decodeNibble(b)
		if !ok {
			return nil, fmt.Errorf("invalid nibble %#U", b)
		}
		ns[i] = n
	}
	return ns, nil
}

//------------------------------------------------------------------------------
