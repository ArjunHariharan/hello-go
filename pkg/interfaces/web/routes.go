package web

import (
	"hello-go/pkg/application"
	"hello-go/pkg/interfaces/web/handler"

	"github.com/labstack/echo/v4"
)

func registerRoutes(e *echo.Echo, a *application.Application) {
	registerUserRoutes(e, a)
}

func registerUserRoutes(e *echo.Echo, a *application.Application) {
	h := handler.NewUserHandler(a.UserApplication)

	e.POST("/users", h.CreateUser)
}
