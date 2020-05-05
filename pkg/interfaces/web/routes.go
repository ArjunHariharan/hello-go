package web

import (
	"hello-go/pkg/interfaces/web/handler"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes is used to add routes to echo server
func registerRoutes(e *echo.Echo) {
	registerUserRoutes(e)
}

func registerUserRoutes(e *echo.Echo) {
	e.POST("/users", handler.CreateUser)
}
