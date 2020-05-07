package main

import (
	"fmt"
	"hello-go/pkg/application"
	"hello-go/pkg/common/config"
	"hello-go/pkg/common/log"
	"hello-go/pkg/infrastructure/database"
	"hello-go/pkg/interfaces/web"
	"os"
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

	application := application.New(repository)
	fmt.Println(application)

	restServer, err := web.New(application)
	if err != nil {
		os.Exit(1)
	}

	l := log.GetLogger("")

	host := fmt.Sprintf(":%d", config.WebServer.Port)

	l.Info().Msg("Starting server on " + host)
	if err := restServer.Start(host); err != nil {
		l.Panic().Err(err).Msg("Failed to start server")
	}
}
