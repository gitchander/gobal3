package ternary

import (
	"fmt"
)

var (
	tritValues = [...]int{-1, 0, +1}

	tritCharsV1 = [...]byte{'T', '0', '1'}
	tritCharsV2 = [...]byte{'-', '0', '+'}
	tritCharsV3 = [...]byte{'N', '0', 'P'}

	tritChars = tritCharsV1
)

func tritToChar(t int) byte {
	return tritChars[t+1]
}

func errInvalidTrit(t int) error {
	return fmt.Errorf("invalid trit value %d", t)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------------------------------------------------------------

func quoRemMinMax(x int, min, max int) (q, r int) {
	base := max - min + 1
	if x < min {
		q, r = quoRem(x-max, base)
		r += max
	} else {
		q, r = quoRem(x-min, base)
		r += min
	}
	return q, r
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

// ... |-7|-6|-5|-4|-3|-2|-1| 0| 1| 2| 3| 4| 5| 6| 7| ...
// ... |-1| 0| 1|-1| 0| 1|-1| 0| 1|-1| 0| 1|-1| 0| 1| ...

func modBal3_v1(x int) int {
	_, r := quoRemMinMax(x, -1, 1)
	return r
}

func modBal3_v2(x int) int {
	const (
		min  = -1
		max  = +1
		base = max - min + 1
	)
	var r int
	if x < min {
		r = (x - max) % base
		r += max
	} else {
		r = (x - min) % base
		r += min
	}
	return r
}

func modBal3_v3(x int) int {
	const (
		min  = -1
		max  = +1
		base = max - min + 1
	)
	r := x % base
	if r > max {
		r -= base
	}
	if r < min {
		r += base
	}
	return r
}

var (
	// modBal3 = modBal3_v1
	// modBal3 = modBal3_v2
	modBal3 = modBal3_v3
)
