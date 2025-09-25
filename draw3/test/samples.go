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
	opt "github.com/gitchander/gobal3/utils/optional"
	"github.com/gitchander/gobal3/utils/random"
)

type Sample struct {
	digits    []int
	dd        draw3.DigitDrawer
	dirName   string
	fileName  string
	digitSize geom.Point2f
	factorX   opt.OptFloat64
}

func makeSample(ds Sample) error {

	var contextSize geom.Point2i
	if ds.factorX.Present {
		var (
			dX      = ds.digitSize.X
			dY      = ds.digitSize.Y
			xn      = len(ds.digits)
			factorX = ds.factorX.Value
		)
		contextSize = geom.Point2i{
			X: int(math.Ceil(dX * (float64(xn) + 1) * factorX)),
			Y: int(math.Ceil(dY)),
		}
	} else {
		var (
			dX = ds.digitSize.X
			dY = ds.digitSize.Y
			xn = len(ds.digits)
		)
		contextSize = geom.Point2i{
			X: int(math.Ceil(dX * float64(xn))),
			Y: int(math.Ceil(dY)),
		}
	}

	c := gg.NewContext(contextSize.X, contextSize.Y)

	c.SetColor(color.White)
	c.Fill()
	c.Clear()

	c.SetColor(color.Black)

	if ds.factorX.Present {
		draw3.DrawDigitsWithFactor(c, ds.dd, ds.digitSize, ds.factorX.Value, ds.digits)
	} else {
		draw3.DrawDigits(c, ds.dd, ds.digitSize, ds.digits)
	}

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

func randDigits(r *random.Rand, n int) []int {
	digits := make([]int, n)
	for i := range digits {
		digits[i] = random.RandIntIn(r, -1, 2) // [-1,0,+1]
	}
	return digits
}

func exampleDrawSamples() error {

	var (
		dirName    = "images"
		filePrefix = "digits"

		//digits = []int{-1}
		//digits = []int{-1, 0, 1, -1, -1, 0, 1}
		digits = randDigits(random.NewRandNext(), 9)

		//digitSize = geom.Point2f{X: 16, Y: 32}
		//digitSize = geom.Point2f{X: 20, Y: 40}
		//digitSize = geom.Point2f{X: 30, Y: 60}
		digitSize = geom.Point2f{X: 50, Y: 100}
	)

	samples := []Sample{
		{
			digits:    digits,
			dd:        draw3.DigitDrawer1{},
			dirName:   dirName,
			fileName:  filePrefix + "_d1.png",
			digitSize: digitSize,
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer2{},
			dirName:   dirName,
			fileName:  filePrefix + "_d2.png",
			digitSize: digitSize,
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer3{},
			dirName:   dirName,
			fileName:  filePrefix + "_d3.png",
			digitSize: digitSize,
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer4{},
			dirName:   dirName,
			fileName:  filePrefix + "_d4.png",
			digitSize: digitSize,
			factorX:   opt.MakePresentFloat64(0.5),
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer5{},
			dirName:   dirName,
			fileName:  filePrefix + "_d5.png",
			digitSize: digitSize,
			factorX:   opt.MakePresentFloat64(0.5),
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer6{},
			dirName:   dirName,
			fileName:  filePrefix + "_d6.png",
			digitSize: digitSize,
		},
		{
			digits:    digits,
			dd:        draw3.DigitDrawer7{},
			dirName:   dirName,
			fileName:  filePrefix + "_d7.png",
			digitSize: digitSize,
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
