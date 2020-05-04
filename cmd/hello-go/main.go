package main

import (
	"fmt"
	"hello-go/pkg/config"
	"hello-go/pkg/infrastructure/database"
	"hello-go/pkg/rest"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.New()
	if err != nil {
		os.Exit(1)
	}

	repository, err := database.New(config)
	if err != nil {
		os.Exit(1)
	}

	repository.Migrate()

	e := echo.New()

	rest.RegisterRoutes(e)
	rest.RegisterValidator(e)
	rest.RegisterErrorHandler(e)

	host := fmt.Sprintf(":%d", config.Server.Port)
	e.Logger.Fatal(e.Start(host))
}
