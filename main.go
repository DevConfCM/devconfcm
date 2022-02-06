package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("static", "frontend/build/static")
	e.File("/", "frontend/build/index.html")
	e.Logger.Fatal(e.Start(":5000"))
}
