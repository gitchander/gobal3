package draw3

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/geom"
)

var (
	ColorTransparent = color.Transparent
	ColorWhite       = color.White
	ColorBlack       = color.Black

	DefaultColorBG = color.Transparent
	DefaultColorFG = color.Black
)

type DigitsImageConfig struct {
	ColorBG     color.Color
	ColorFG     color.Color
	DigitDrawer DigitDrawer
	DigitSize   geom.Point2f
	Stride      geom.Point2f
}

func StrideHalfX(digitSize geom.Point2f) geom.Point2f {
	return geom.Point2f{
		X: digitSize.X / 2,
		Y: digitSize.Y,
	}
}

func MakeDigitsImage(dic DigitsImageConfig, digits []int) (image.Image, error) {

	var (
		xn = len(digits)
		yn = 1

		contextSize = geom.Point2i{
			X: int(math.Ceil(float64(xn-1)*dic.Stride.X + dic.DigitSize.X)),
			Y: int(math.Ceil(float64(yn-1)*dic.Stride.Y + dic.DigitSize.Y)),
		}
	)

	c := gg.NewContext(contextSize.X, contextSize.Y)

	c.SetColor(dic.ColorBG)
	c.Fill()
	c.Clear()

	c.SetColor(dic.ColorFG)

	DrawDigits(c, dic.DigitDrawer, dic.DigitSize, dic.Stride, digits)

	return c.Image(), nil
}

func MakeDigitsImage2D(dic DigitsImageConfig, d2d *Digits2D) (image.Image, error) {

	var (
		matrixSize = d2d.Size()

		xn = matrixSize.X
		yn = matrixSize.Y

		contextSize = geom.Point2i{
			X: int(math.Ceil(float64(xn-1)*dic.Stride.X + dic.DigitSize.X)),
			Y: int(math.Ceil(float64(yn-1)*dic.Stride.Y + dic.DigitSize.Y)),
		}
	)

	c := gg.NewContext(contextSize.X, contextSize.Y)

	c.SetColor(dic.ColorBG)
	c.Fill()
	c.Clear()

	c.SetColor(dic.ColorFG)

	DrawDigits2D(c, dic.DigitDrawer, dic.DigitSize, dic.Stride, d2d)

	return c.Image(), nil
}

func ImageEncodePNG(m image.Image) ([]byte, error) {
	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func ImageSavePNG(m image.Image, filename string) error {

	data, err := ImageEncodePNG(m)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
