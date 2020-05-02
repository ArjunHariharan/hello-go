package user

import "github.com/labstack/echo/v4"

// RegisterRoutes for user module
func RegisterRoutes(e *echo.Echo) {
	e.POST("/users", create)
}
