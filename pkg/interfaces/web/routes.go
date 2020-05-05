package web

import (
	"hello-go/pkg/interfaces/web/handler"

	"github.com/labstack/echo/v4"
)

func registerRoutes(e *echo.Echo) {
	registerUserRoutes(e)
}

func registerUserRoutes(e *echo.Echo) {
	e.POST("/users", handler.CreateUser)
}
