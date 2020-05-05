package application

import "hello-go/pkg/infrastructure/database"

// Application wraps all the services
type Application struct {
	UserApplication *UserApplication
}

// New creates an application instance
func New(r *database.Repository) *Application {
	userApplication := &UserApplication{r: r.UserRepository}

	application := Application{UserApplication: userApplication}

	return &application
}
