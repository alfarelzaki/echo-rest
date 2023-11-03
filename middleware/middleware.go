package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
)

var isAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
