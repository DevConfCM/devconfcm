package handlers

import (
	"devconfcm/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func GetUser(c echo.Context) error {

	u := &models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Username: c.Param("username"),
	}
	return c.JSON(http.StatusOK, u)
}
