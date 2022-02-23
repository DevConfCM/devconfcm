package router

import (
	"sync"

	"github.com/labstack/echo/v4"
)

type (
	route struct {
		Path    string
		Method  string
		Handler echo.HandlerFunc
	}
)

var once sync.Once

func Start() {
	once.Do(start)
}

func start() {
	router := echo.New()

	/// add routes
	for _, route := range routes {
		router.Add(route.Method, route.Path, route.Handler)
	}

	router.Static("static", "website/build/static")
	router.File("/", "website/build/index.html")

	// start server
	router.Logger.Fatal(router.Start(":5000"))
}
