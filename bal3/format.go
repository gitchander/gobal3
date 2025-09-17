package bal3

import (
	"fmt"

	ivl "github.com/gitchander/gobal3/utils/interval"
)

func tryteFormat[Tryte GenericTryte](n int, a Tryte) (string, error) {
	var (
		bs = make([]byte, n)
		j  = n - 1
		k  = j
	)
	for i := 0; i < n; i++ {
		t := getTrit(a, i)
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

func tryteFormatAllTrits[Tryte GenericTryte](n int, a Tryte) (string, error) {
	var (
		bs = make([]byte, n)
		j  = n - 1
	)
	for i := 0; i < n; i++ {
		t := getTrit(a, i)
		char, err := tritToChar(t)
		if err != nil {
			return "", fmt.Errorf("trit[%d] error: %w", i, err)
		}
		bs[j] = char
		j--
	}
	return string(bs), nil
}

func tryteParse[Tryte GenericTryte](n int, s string) (Tryte, error) {
	var v Tryte
	var count int
	bs := []byte(s)
	for i, b := range bs {
		if b == '_' {
			continue
		}
		t, err := charToTrit(b)
		if err != nil {
			return v, fmt.Errorf("char[%d] error: %w", i, err)
		}
		v = tryteShl(n, v, 1) // v = v << 1
		v = setTrit(v, 0, t)  // v[0] = t
		count++
	}
	if l := ivl.Ivl(1, n+1); not(l.Contains(count)) {
		return v, fmt.Errorf("invalid number of trits: have %d, want %v", count, l)
	}
	return v, nil
}
