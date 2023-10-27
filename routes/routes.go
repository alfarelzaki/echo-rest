package routes

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// creates new instances of echo web framework that allows to configure app, define routes, and setup middleware
	e := echo.New()


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	return e
}