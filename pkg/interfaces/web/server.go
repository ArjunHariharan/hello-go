package web

import "github.com/labstack/echo/v4"

// NewServer creates an instance of echo server
func NewServer() (*echo.Echo, error) {
	e := echo.New()

	registerValidator(e)
	registerErrorHandler(e)
	registerRoutes(e)

	return e, nil
}
