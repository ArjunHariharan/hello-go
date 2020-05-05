package user

import repository "hello-go/pkg/domain/Repository"

// UserApplication class
type UserApplication struct {
	r repository.UserRepository
}

// SaveUser service
func (u *UserApplication) SaveUser() error {
	return nil
}
