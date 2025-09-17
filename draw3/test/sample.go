package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"

	"github.com/gitchander/gobal3/draw3"
	"github.com/gitchander/gobal3/geom"
)

type Sample struct {
	digits    []int
	dd        draw3.DigitDrawer
	dirName   string
	fileName  string
	digitSize geom.Point2f
}

func makeSample(ds Sample) error {

	contextSize := geom.Point2i{
		X: int(math.Ceil(ds.digitSize.X * float64(len(ds.digits)))),
		Y: int(math.Ceil(ds.digitSize.Y)),
	}

	c := gg.NewContext(contextSize.X, contextSize.Y)

	c.SetColor(color.White)
	c.Clear()

	draw3.DrawDigits(c, ds.dd, ds.digitSize, ds.digits)

	filename := ds.fileName

	if ds.dirName != "" {
		err := MkdirIfNotExist(ds.dirName)
		if err != nil {
			return err
		}
		filename = filepath.Join(ds.dirName, filename)
	}

	return c.SavePNG(filename)
}

func MkdirIfNotExist(dirname string) error {
	fi, err := os.Stat(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dirname, os.ModePerm)
		}
		return err
	}
	if !(fi.IsDir()) {
		return fmt.Errorf("name %q is not a directory", dirname)
	}
	return nil
}

func exampleDrawSamples() error {

	digits := []int{-1, 0, +1, -1, -1}

	samples := []Sample{
		{
			digits:    digits,
			dd:        draw3.DigitDrawer1{},
			dirName:   "images",
			fileName:  "digits_d1.png",
			digitSize: geom.Point2f{X: 50, Y: 100},
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer2{},
			dirName:   "images",
			fileName:  "digits_d2.png",
			digitSize: geom.Point2f{X: 50, Y: 100},
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer3{},
			dirName:   "images",
			fileName:  "digits_d3.png",
			digitSize: geom.Point2f{X: 50, Y: 100},
		},
	}
	for _, sample := range samples {
		err := makeSample(sample)
		if err != nil {
			return err
		}
	}
	return nil
}
