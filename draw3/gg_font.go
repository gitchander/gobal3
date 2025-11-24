package draw3

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gitchander/gobal3/geom"
)

type GG_Context struct {
	c *gg.Context
}

func NewGG_Context(c *gg.Context) *GG_Context {
	return &GG_Context{c: c}
}

func (p *GG_Context) SetFontSize(fontSize float64) error {

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	p.c.SetFontFace(face)

	return nil
}

func (p *GG_Context) SetFontSizeFactor(dX, dY float64, fontSizeFactor float64) error {
	var (
		vmin     = geom.Vmin(dX, dY)
		fontSize = vmin * fontSizeFactor
	)
	return p.SetFontSize(fontSize)
}
