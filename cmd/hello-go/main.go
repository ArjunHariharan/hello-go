package main

import (
	"hello-go/pkg/rest"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	rest.RegisterRoutes(e)
	rest.RegisterValidator(e)
	rest.RegisterErrorHandler(e)

	e.Logger.Fatal(e.Start(":1323"))
}
