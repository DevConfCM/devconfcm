package devconfcm

import (
	"devconfcm/pkg/routes"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	routes.RegisterRoutes(e)

	e.Static("static", "website/build/static")
	e.File("/", "website/build/index.html")
	e.Logger.Fatal(e.Start(":5000"))
}