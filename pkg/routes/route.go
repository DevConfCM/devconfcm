package routes

import (
	"github.com/labstack/echo/v4"
)

type (
	Route struct {
		Path    string
		Method  string
		Handler echo.HandlerFunc
	}
)

func RegisterRoutes(e *echo.Echo) {
	for _, route := range routes {
		e.Add(route.Method, route.Path, route.Handler)
	}
}
