package main

import (
	"bytes"
	"fmt"
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"

	"github.com/gitchander/gobal3/draw3"
	"github.com/gitchander/gobal3/geom"
	opt "github.com/gitchander/gobal3/utils/optional"
)

//------------------------------------------------------------------------------

type DrawParams struct {
	colorBG     string
	colorFG     string
	digitDrawer string
	digitSize   string
	digits      string
}

func QueryDrawParams(c echo.Context) (*DrawParams, error) {
	dp := &DrawParams{
		colorBG:     c.QueryParam("color_bg"),
		colorFG:     c.QueryParam("color_fg"),
		digitDrawer: c.QueryParam("digit_drawer"),
		digitSize:   c.QueryParam("digit_size"),
		digits:      c.QueryParam("digits"),
	}
	return dp, nil
}

//------------------------------------------------------------------------------

type DrawConfig struct {
	colorBG color.Color
	colorFG color.Color

	dd        draw3.DigitDrawer
	digitSize geom.Point2f
	factorX   opt.OptFloat64
	digits    []int
}

func errParseParam(paramName string, err error) error {
	return fmt.Errorf("parse %s error: %w", paramName, err)
}

func ParseDrawConfig(dp *DrawParams) (*DrawConfig, error) {

	colorBG, err := parseColorByParam(dp.colorBG, DefaultColorBG)
	if err != nil {
		return nil, errParseParam("colorBG", err)
	}

	colorFG, err := parseColorByParam(dp.colorFG, DefaultColorFG)
	if err != nil {
		return nil, errParseParam("colorFG", err)
	}

	digitDrawer, err := parseInt(dp.digitDrawer)
	if err != nil {
		return nil, errParseParam("digitDrawer", err)
	}

	var (
		dd      draw3.DigitDrawer
		factorX opt.OptFloat64
	)
	switch digitDrawer {
	case 1:
		dd = draw3.DigitDrawer1{}
	case 2:
		dd = draw3.DigitDrawer2{}
	case 3:
		dd = draw3.DigitDrawer3{}
	case 4:
		dd = draw3.DigitDrawer4{}
		factorX.SetValue(0.5)
	case 5:
		dd = draw3.DigitDrawer5{}
		factorX.SetValue(0.5)
	case 6:
		dd = draw3.DigitDrawer6{}
	case 7:
		dd = draw3.DigitDrawer7{}
	default:
		return nil, fmt.Errorf("unknown digitDrawer %d", digitDrawer)
	}

	digitSize, err := parseIntByParam(dp.digitSize, 50)
	if err != nil {
		return nil, errParseParam("digitSize", err)
	}
	digitSizeX := float64(digitSize)

	digits, err := ParseDigits(dp.digits)
	if err != nil {
		return nil, errParseParam("digits", err)
	}

	dc := &DrawConfig{
		colorBG:   colorBG,
		colorFG:   colorFG,
		dd:        dd,
		digitSize: geom.MakePoint2f(digitSizeX, 2*digitSizeX),
		factorX:   factorX,
		digits:    digits,
	}
	return dc, nil
}

//------------------------------------------------------------------------------

func makeImagePNG(dc *DrawConfig) ([]byte, error) {

	var contextSize geom.Point2i
	if dc.factorX.Present {
		var (
			dX      = dc.digitSize.X
			dY      = dc.digitSize.Y
			xn      = len(dc.digits)
			factorX = dc.factorX.Value
		)
		contextSize = geom.Point2i{
			X: int(math.Ceil(dX * (float64(xn) + 1) * factorX)),
			Y: int(math.Ceil(dY)),
		}
	} else {
		var (
			dX = dc.digitSize.X
			dY = dc.digitSize.Y
			xn = len(dc.digits)
		)
		contextSize = geom.Point2i{
			X: int(math.Ceil(dX * float64(xn))),
			Y: int(math.Ceil(dY)),
		}
	}

	c := gg.NewContext(contextSize.X, contextSize.Y)

	c.SetColor(dc.colorBG)
	c.Fill()
	c.Clear()

	c.SetColor(dc.colorFG)

	if dc.factorX.Present {
		draw3.DrawDigitsWithFactor(c, dc.dd, dc.digitSize, dc.factorX.Value, dc.digits)
	} else {
		draw3.DrawDigits(c, dc.dd, dc.digitSize, dc.digits)
	}

	var b bytes.Buffer
	err := c.EncodePNG(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
