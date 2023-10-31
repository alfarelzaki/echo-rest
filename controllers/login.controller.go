package controllers

import (
	"echo-rest/helpers"
	"echo-rest/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// call function from helpers and create password based on param
func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}
	
	// if the username and password not match
	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "Berhasil login")
}