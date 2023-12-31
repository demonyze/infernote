package main

import (
	"net/http"

	data "github.com/demonyze/infernote/internal/data/chords"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data.Chords)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
