package web

import (
	"hello-go/pkg/application"

	"github.com/labstack/echo/v4"
)

// New creates an instance of echo server
func New(a *application.Application) (*echo.Echo, error) {
	e := echo.New()

	registerValidator(e)
	registerErrorHandler(e)
	registerMiddleware(e)
	registerRoutes(e, a)

	return e, nil
}
