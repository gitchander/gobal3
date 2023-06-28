package ternary

import (
	"fmt"
	"strings"
)

func PrintableLogicTable(prefix string, f BinaryFunc) string {

	const n = 3

	var (
		vs = [n]int{-1, 0, 1}
		//chars = []byte("T01")
		chars = []byte("-0+")
	)
	tritChar := func(t int) byte {
		return chars[t+1]
	}

	var br strings.Builder
	var frameLine string

	// make frameLine
	{
		fmt.Fprintf(&br, "%s+---+", prefix)
		for j := 0; j < n; j++ {
			fmt.Fprintf(&br, "---+")
		}
		br.WriteByte('\n')
		frameLine = br.String()
		br.Reset()
	}

	// write b values
	{
		br.WriteString(frameLine)
		fmt.Fprintf(&br, "%s| %c |", prefix, ' ')
		for j := 0; j < n; j++ {
			b := vs[j]
			fmt.Fprintf(&br, " %c |", tritChar(b))
		}
		br.WriteByte('\n')
		br.WriteString(frameLine)
	}

	for i := 0; i < n; i++ {
		a := vs[i]
		fmt.Fprintf(&br, "%s| %c |", prefix, tritChar(a))
		for j := 0; j < n; j++ {
			b := vs[j]
			c := f(a, b)
			fmt.Fprintf(&br, " %c |", tritChar(c))
		}
		br.WriteByte('\n')
		br.WriteString(frameLine)
	}

	return br.String()
}
