package main

import (
	"fmt"
	"image/color"
)

var (
	DefaultColorBG = color.Transparent
	DefaultColorFG = color.Black
)

func ParseColor(s string) (color.Color, error) {

	bs := []byte(s)
	if (len(bs) > 0) && (bs[0] == '#') { // skip first #
		bs = bs[1:]
	}

	ns, err := decodeNibbles(bs)
	if err != nil {
		return nil, fmt.Errorf("invalid color (%s): %w", s, err)
	}

	var c color.Color
	switch k := len(ns); k {
	case 3: // rgb
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[0]),
			G: nibblesToByte(ns[1], ns[1]),
			B: nibblesToByte(ns[2], ns[2]),
			A: 0xff,
		}
	case 4: // rgba
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[0]),
			G: nibblesToByte(ns[1], ns[1]),
			B: nibblesToByte(ns[2], ns[2]),
			A: nibblesToByte(ns[3], ns[3]),
		}
	case 6: // rrggbb
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[1]),
			G: nibblesToByte(ns[2], ns[3]),
			B: nibblesToByte(ns[4], ns[5]),
			A: 0xff,
		}
	case 8: // rrggbbaa
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[1]),
			G: nibblesToByte(ns[2], ns[3]),
			B: nibblesToByte(ns[4], ns[5]),
			A: nibblesToByte(ns[6], ns[7]),
		}
	default:
		return nil, fmt.Errorf("invalid color (%s): invalid number of nibbles", s)
	}
	c = color.RGBAModel.Convert(c)
	return c, nil
}

func MustParseColor(s string) color.Color {
	c, err := ParseColor(s)
	if err != nil {
		panic(err)
	}
	return c
}
