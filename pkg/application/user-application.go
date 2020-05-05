package application

import (
	"fmt"
	"hello-go/pkg/domain/repository"
)

// UserApplication class
type UserApplication struct {
	r repository.UserRepository
}

// CreateUserDto for create request
type CreateUserDto struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// SaveUser service
func (u *UserApplication) SaveUser() error {
	fmt.Println("User application called")
	u.r.SaveUser(nil)

	return nil
}
