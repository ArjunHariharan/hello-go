package main

import (
	"fmt"
	"hello-go/pkg/application"
	"hello-go/pkg/infrastructure/config"
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

	host := fmt.Sprintf(":%d", config.WebServer.Port)
	restServer.Start(host)
}
