package web

import (
	"hello-go/pkg/application"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

// New creates an instance of echo server
func New(a *application.Application) (*echo.Echo, error) {
	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)

	registerValidator(e)
	registerErrorHandler(e)
	registerMiddleware(e)
	registerRoutes(e, a)

	return e, nil
}
