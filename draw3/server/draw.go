package main

import (
	"fmt"
	"image/color"

	"github.com/labstack/echo/v4"

	"github.com/gitchander/gobal3/draw3"
	"github.com/gitchander/gobal3/geom"
	"github.com/gitchander/gobal3/utils/random"
)

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

func errParseParam(paramName string, err error) error {
	return fmt.Errorf("parse %s error: %w", paramName, err)
}

type DigitsConfig struct {
	DIC    *draw3.DigitsImageConfig
	Digits []int
}

func ParseDigitsConfig(dp *DrawParams) (*DigitsConfig, error) {

	var (
		DefaultColorBG = color.Transparent
		DefaultColorFG = color.Black
	)

	colorBG, err := parseColorByParam(dp.colorBG, DefaultColorBG)
	if err != nil {
		return nil, errParseParam("colorBG", err)
	}

	colorFG, err := parseColorByParam(dp.colorFG, DefaultColorFG)
	if err != nil {
		return nil, errParseParam("colorFG", err)
	}

	digitSizeX, err := parseIntByParam(dp.digitSize, 50)
	if err != nil {
		return nil, errParseParam("digitSize", err)
	}
	digitSize := geom.Point2f{
		X: float64(digitSizeX),
		Y: float64(digitSizeX) * 2,
	}

	digitDrawer, err := parseInt(dp.digitDrawer)
	if err != nil {
		return nil, errParseParam("digitDrawer", err)
	}

	var (
		dd draw3.DigitDrawer

		stride geom.Point2f

		stride1 = digitSize
		stride2 = draw3.StrideHalfX(digitSize)
	)
	switch digitDrawer {
	case 1:
		dd = draw3.DigitDrawer1{}
		stride = stride1
	case 2:
		dd = draw3.DigitDrawer2{}
		stride = stride1
	case 3:
		dd = draw3.DigitDrawer3{}
		stride = stride1
	case 4:
		dd = draw3.DigitDrawer4{}
		stride = stride2
	case 5:
		dd = draw3.DigitDrawer5{}
		stride = stride2
	case 6:
		dd = draw3.DigitDrawer6{}
		stride = stride1
	case 7:
		dd = draw3.DigitDrawer7{}
		stride = stride1
	default:
		return nil, fmt.Errorf("unknown digitDrawer %d", digitDrawer)
	}

	var digits []int
	switch {
	case dp.digits == "":
		digits = draw3.RandDigits(random.NewRandNext(), 9)
	case dp.digits == "random":
		digits = draw3.RandDigits(random.NewRandNext(), 9)
	default:
		{
			digitsParsed, err := ParseDigits(dp.digits)
			if err != nil {
				return nil, errParseParam("digits", err)
			}
			digits = digitsParsed
		}
	}

	dc := &DigitsConfig{
		DIC: &draw3.DigitsImageConfig{
			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: dd,
			DigitSize:   digitSize,
			Stride:      stride,
		},
		Digits: digits,
	}
	return dc, nil
}

//------------------------------------------------------------------------------
