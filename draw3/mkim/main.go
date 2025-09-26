package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gitchander/gobal3/draw3"
	"github.com/gitchander/gobal3/geom"
	"github.com/gitchander/gobal3/utils/osutils"
	"github.com/gitchander/gobal3/utils/random"
)

func main() {
	checkError(makeImages())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func makeImages() error {

	var (
		//colorBG = color.White
		colorBG = draw3.DefaultColorBG
		colorFG = draw3.DefaultColorFG

		dirName    = "images"
		filePrefix = "digits"

		//digits = []int{-1}
		//digits = []int{-1, 0, 1, -1, -1, 0, 1}

		r        = random.NewRandNext()
		digits   = draw3.RandDigits(r, 9)
		digits2D = draw3.RandDigits2D(r, 10, 5)

		//digitSize = geom.Point2f{X: 16, Y: 32}
		//digitSize = geom.Point2f{X: 20, Y: 40}
		//digitSize = geom.Point2f{X: 30, Y: 60}
		digitSize = geom.Point2f{X: 50, Y: 100}

		stride1 = digitSize
		stride2 = draw3.StrideHalfX(digitSize)
	)

	cs := []draw3.DigitsImageConfig{
		{

			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: draw3.DigitDrawer1{},
			DigitSize:   digitSize,
			Stride:      stride1,
		},
		{
			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: draw3.DigitDrawer2{},
			DigitSize:   digitSize,
			Stride:      stride1,
		},
		{
			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: draw3.DigitDrawer3{},
			DigitSize:   digitSize,
			Stride:      stride1,
		},
		{
			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: draw3.DigitDrawer4{},
			DigitSize:   digitSize,
			Stride:      stride2,
		},
		{
			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: draw3.DigitDrawer5{},
			DigitSize:   digitSize,
			Stride:      stride2,
		},
		{
			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: draw3.DigitDrawer6{},
			DigitSize:   digitSize,
			Stride:      stride1,
		},
		{
			ColorBG:     colorBG,
			ColorFG:     colorFG,
			DigitDrawer: draw3.DigitDrawer7{},
			DigitSize:   digitSize,
			Stride:      stride1,
		},
	}

	err := osutils.MkdirIfNotExist(dirName)
	if err != nil {
		return err
	}

	for i, c := range cs {
		m, err := draw3.MakeDigitsImage(c, digits)
		if err != nil {
			return err
		}
		var (
			fileNumber = i + 1
			filename   = filepath.Join(dirName, fmt.Sprintf("%s_%03d.png", filePrefix, fileNumber))
		)
		err = draw3.ImageSavePNG(m, filename)
		if err != nil {
			return err
		}
	}

	// matrixes
	for i, c := range cs {
		m, err := draw3.MakeDigitsImage2D(c, digits2D)
		if err != nil {
			return err
		}
		var (
			fileNumber = i + 1
			filename   = filepath.Join(dirName, fmt.Sprintf("%s_matrix_%03d.png", filePrefix, fileNumber))
		)
		err = draw3.ImageSavePNG(m, filename)
		if err != nil {
			return err
		}
	}

	return nil
}
