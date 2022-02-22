package handlers

import (
	"devconfcm/pkg/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	models.CreateUser(u)
	return c.JSONPretty(http.StatusOK, u, " ")
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	u := models.GetUser(id)
	if u == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No user with given ID found")
	}
	return c.JSONPretty(http.StatusOK, u, " ")
}

func GetAllUsers(c echo.Context) error {
	users := models.GetAllUsers()
	return c.JSONPretty(http.StatusOK, users, " ")
}

func UpdateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := models.UpdateUser(id, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSONPretty(http.StatusOK, user, " ")
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := models.DeleteUser(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSONPretty(http.StatusNoContent, nil, " ")
}
