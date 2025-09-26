package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gitchander/gobal3/draw3"
)

// http://localhost:1323/draw_digits?digit_drawer=3&digits=11T01T1
// http://localhost:1323/draw_digits?color_bg=ff0&digit_drawer=3&digits=11T01
// http://localhost:1323/draw_digits?digit_drawer=3&digit_size=100&digits=T01T01
// http://localhost:1323/draw_digits?color_bg=fff&digit_drawer=7&digit_size=70&digits=T01T0NZ
// http://localhost:1323/draw_digits?color_bg=00f&color_fg=fff&digit_drawer=4&digit_size=70&digits=T01T0NZ
// http://localhost:1323/draw_digits?digit_drawer=3&digit_size=70

func main() {
	checkError(runServer())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	core := &Core{}
	e := echo.New()
	//e.GET("/", core.HandleIndex)
	e.GET("/draw_digits", core.HandleDraw)
	return e.Start(":1323")
}

type Core struct{}

func (p *Core) HandleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "HandleIndex")
}

func (p *Core) HandleDraw(c echo.Context) error {

	dp, err := QueryDrawParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dc, err := ParseDigitsConfig(dp)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	m, err := draw3.MakeDigitsImage(*(dc.DIC), dc.Digits)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dataPNG, err := draw3.ImageEncodePNG(m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", dataPNG)
}
