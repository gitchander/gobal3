package ternary

import (
	"fmt"
)

var tritValues = [...]int{-1, 0, +1}

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
