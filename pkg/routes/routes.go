package routes

import (
	"devconfcm/pkg/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/api/user/", handlers.CreateUser)
	e.GET("/api/user/:username/", handlers.GetUser)
}
