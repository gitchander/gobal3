package ternary

import (
	"github.com/gitchander/gobal3/utils/ptab"
)

func PrintableBinaryTable(prefix string, f BinaryFunc) string {

	n := len(tritValues)

	sss := make([][]string, n+1)
	for i := range sss {
		sss[i] = make([]string, n+1)
	}

	for i := 0; i < n; i++ {
		var (
			a = tritValues[i]
			b = tritValues[i]
		)
		sss[i+1][0] = string(tritToChar(a))
		sss[0][i+1] = string(tritToChar(b))
	}

	for i := 0; i < n; i++ {
		a := tritValues[i]
		for j := 0; j < n; j++ {
			b := tritValues[j]
			c := f(a, b)
			sss[i+1][j+1] = string(tritToChar(c))
		}
	}

	return ptab.PrintableTable(prefix, sss)
}
