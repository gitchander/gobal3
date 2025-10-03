package main

import (
	"fmt"
	"log"

	"github.com/gitchander/gobal3/bal3/base27/v2/base27"
)

func main() {
	checkError(encodeDigits())
	checkError(decodeDigits())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func encodeDigits() error {
	for digit := base27.MinDigit; digit <= base27.MaxDigit; digit++ {
		char, err := base27.DigitToChar(digit)
		if err != nil {
			return err
		}
		fmt.Printf("%+3d => %q\n", digit, char)
	}
	return nil
}

func decodeDigits() error {
	chars := []byte("0ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for _, char := range chars {
		digit, err := base27.CharToDigit(char)
		if err != nil {
			return err
		}
		fmt.Printf("%q => %+3d\n", char, digit)
	}
	return nil
}
