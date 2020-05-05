package entity

import (
	uuid "github.com/satori/go.uuid"
)

// User entity
type User struct {
	ID       int
	UUID     uuid.UUID
	Email    string
	Password string
}

// New returns an instance of User
func New(e string, p string) *User {
	u := uuid.NewV4()

	return &User{Email: e, Password: p, UUID: u}
}
