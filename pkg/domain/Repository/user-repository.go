package repository

import "hello-go/pkg/domain/entity"

// UserRepository for user entity
type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, error)
}
