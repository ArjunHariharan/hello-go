package rest

import (
	"hello-go/pkg/user"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes is used to add routes to echo server
func RegisterRoutes(e *echo.Echo) {
	user.RegisterRoutes(e)
}
