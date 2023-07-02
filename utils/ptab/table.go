package ptab

import (
	"strings"
)

type borderRunes [][]rune

var borderTables = []borderRunes{
	{
		[]rune("+-++"),
		[]rune("| ||"),
		[]rune("+-++"),
		[]rune("+-++"),
	},
	{
		[]rune("┌─┬┐"),
		[]rune("│ ││"),
		[]rune("├─┼┤"),
		[]rune("└─┴┘"),
	},
	{
		[]rune("╓─╥╖"),
		[]rune("║ ║║"),
		[]rune("╟─╫╢"),
		[]rune("╙─╨╜"),
	},
	{
		[]rune("╒═╤╕"),
		[]rune("│ ││"),
		[]rune("╞═╪╡"),
		[]rune("╘═╧╛"),
	},
	{
		[]rune("╔═╦╗"),
		[]rune("║ ║║"),
		[]rune("╠═╬╣"),
		[]rune("╚═╩╝"),
	},
	{
		[]rune("╔═╤╗"),
		[]rune("║ │║"),
		[]rune("╟─┼╢"),
		[]rune("╚═╧╝"),
	},
	{
		[]rune("┌─╥┐"),
		[]rune("│ ║│"),
		[]rune("╞═╬╡"),
		[]rune("└─╨┘"),
	},
}

func PrintableTable(prefix string, sss [][]string) string {

	bn := 0
	for _, ss := range sss {
		bn = maxInt(bn, len(ss))
	}

	const contentLen = 3

	bt := borderTables[0]

	fillRune := bt[1][1]

	contentRunes := repeatRune(fillRune, contentLen)

	var b strings.Builder
	var frameLine string

	formatCell := func(s string) string {
		rs := []rune(s)
		if len(rs) > contentLen {
			rs = rs[:contentLen]
		}
		for i := range contentRunes {
			contentRunes[i] = fillRune
		}
		j := ceilDiv(contentLen-len(rs), 2)
		for i, r := range rs {
			contentRunes[j+i] = r
		}
		return string(contentRunes)
	}

	// make frameLine
	{
		frameLine = makeFrameLine(bn, contentLen, bt[0])

		b.WriteString(prefix)
		b.WriteString(frameLine)
		b.WriteByte('\n')

		frameLine = makeFrameLine(bn, contentLen, bt[2])
	}

	for i, ss := range sss {
		bts := bt[1]

		if i > 0 {
			b.WriteString(prefix)
			b.WriteString(frameLine)
			b.WriteByte('\n')
		}

		b.WriteString(prefix)
		b.WriteRune(bts[0])
		for j := 0; j < bn; j++ {
			if j > 0 {
				b.WriteRune(bts[2])
			}
			var fs string
			if j < len(ss) {
				fs = ss[j]
			}
			b.WriteString(formatCell(fs))
		}
		b.WriteRune(bts[3])
		b.WriteByte('\n')
	}

	{
		frameLine = makeFrameLine(bn, contentLen, bt[3])

		b.WriteString(prefix)
		b.WriteString(frameLine)
		b.WriteByte('\n')
	}

	return b.String()
}

func makeFrameLine(n int, contentLen int, bts []rune) string {
	frs := string(repeatRune(bts[1], contentLen))
	var b strings.Builder
	b.WriteRune(bts[0])
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteRune(bts[2])
		}
		b.WriteString(frs)
	}
	b.WriteRune(bts[3])
	return b.String()
}

func ceilDiv(a, b int) int {
	return (a + (b - 1)) / b
}

func repeatRune(r rune, count int) []rune {
	rs := make([]rune, count)
	for i := range rs {
		rs[i] = r
	}
	return rs
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
