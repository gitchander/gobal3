package base27

import (
	"fmt"

	"github.com/gitchander/gobal3/utils/zigzag"
)

const tritsPerDigit = 3 // 3^3 = 27

const (
	MinDigit = -13
	MaxDigit = +13
)

// const (
// 	abc   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // 26 letters
// 	chars = "0" + abc                    // 1 + 26 = 27 letters
// )

func DigitToChar(digit int) (char byte, err error) {
	if (digit < MinDigit) || (MaxDigit < digit) {
		return 0, fmt.Errorf("invalid digit %d", digit)
	}
	u := zigzag.EncodeInt(digit)
	if u == 0 {
		char = '0'
	} else {
		char = byte('A' + (u - 1))
	}
	return char, nil
}

func CharToDigit(char byte) (digit int, err error) {
	var u uint
	if char == '0' {
		u = 0
	} else {
		if ('A' <= char) && (char <= 'Z') {
			u = uint(char-'A') + 1
		} else {
			return 0, fmt.Errorf("invalid char %q", char)
		}
	}
	digit = zigzag.DecodeInt(u)
	return digit, nil
}
