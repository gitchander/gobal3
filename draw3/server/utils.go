package main

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/gitchander/gobal3/utils/gocolor"
)

func formatInt(a int) string {
	return strconv.Itoa(a)
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

//------------------------------------------------------------------------------

func parseIntByParam(s string, defaultValue int) (int, error) {
	if s == "" {
		return defaultValue, nil
	}
	return parseInt(s)
}

func parseColorByParam(s string, defaultColor color.Color) (color.Color, error) {
	if s == "" {
		return defaultColor, nil
	}
	return gocolor.ParseColor(s)
}

//------------------------------------------------------------------------------

var charToDigitMap = map[byte]int{
	'T': -1,
	'0': 0,
	'1': +1,

	'N': -1,
	'Z': 0,
	'P': +1,
}

func charToDigit(char byte) (digit int, err error) {
	digit, ok := charToDigitMap[char]
	if !ok {
		return 0, fmt.Errorf("invalid char %U", char)
	}
	return digit, nil
}

func ParseDigits(s string) ([]int, error) {
	var (
		chars  = []byte(s)
		digits = make([]int, len(chars))
	)
	for i, char := range chars {
		var digit int
		digit, err := charToDigit(char)
		if err != nil {
			return nil, err
		}
		digits[i] = digit
	}
	return digits, nil
}
