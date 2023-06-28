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
