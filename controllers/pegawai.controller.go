package controllers

import (
	"echo-rest/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// func in controller are using function which is defined previously in models folder
func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai (c echo.Context) error {
	// store the parameter inputed from externall app
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	// passing data from form to models func
	result, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON((http.StatusInternalServerError), err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai (c echo.Context) error {
	// store value from external form
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	// converting inserted string id to int64 (as stated in models)
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// passing stored data from input to models
	result, err := models.UpdatePegawai(conv_id, nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePegawai (c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}