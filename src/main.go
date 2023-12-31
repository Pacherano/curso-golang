package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	var i int
}
