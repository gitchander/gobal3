package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// http://localhost:1323/draw_digits?digit_drawer=3&digits=11T01T1
// http://localhost:1323/draw_digits?color_bg=ff0&digit_drawer=3&digits=11T01
// http://localhost:1323/draw_digits?digit_drawer=3&digit_size=100&digits=T01T01
// http://localhost:1323/draw_digits?color_bg=fff&digit_drawer=7&digit_size=70&digits=T01T0NZ
// http://localhost:1323/draw_digits?color_bg=00f&color_fg=fff&digit_drawer=4&digit_size=70&digits=T01T0NZ

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
	e.GET("/draw_digits", core.HandleDraw2)
	return e.Start(":1323")
}

type Core struct{}

func (p *Core) HandleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "HandleIndex")
}

func (p *Core) HandleDraw1(c echo.Context) error {
	fmt.Println("id:", c.Param("id"))
	return c.String(http.StatusOK, "HandleDraw")
}

func (p *Core) HandleDraw2(c echo.Context) error {

	dp, err := QueryDrawParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//fmt.Println(dp)

	dc, err := ParseDrawConfig(dp)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//fmt.Println(dc)

	dataPNG, err := makeImagePNG(dc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", dataPNG)
}
