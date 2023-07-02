package main

import (
	"fmt"

	"github.com/gitchander/gobal3/utils/ptab"
)

func main() {
	sss := [][]string{
		{"1", "2", "3", "4"},
		{"A", "B", "C", "D"},
		{"5", "6", "7", "8", "12"},
		{"E", "F", "", ""},
		{"5", "6", "7", "8"},
		{"E", "F", "G", "H"},
	}
	s := ptab.PrintableTable("\t", sss)
	fmt.Println(s)
}
