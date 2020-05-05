package database

import (
	"fmt"
	"hello-go/pkg/domain/entity"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type userModel struct {
	ID        int
	UUID      uuid.UUID  `gorm:"type:uuid"`
	Email     string     `gorm:"type:varchar(256);unique_index"`
	Password  string     `gorm:"type:varchar(72)"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UserRepository implements the interface
type UserRepository struct {
	db *gorm.DB
}

//SaveUser saves the user
func (*UserRepository) SaveUser(*entity.User) (*entity.User, error) {
	fmt.Println("save user called from db repository")
	return nil, nil
}

// GetUser fetches the user from DB
func (*UserRepository) GetUser(uint64) (*entity.User, error) {
	return nil, nil
}

// GetUsers fetches all the users
func (*UserRepository) GetUsers() ([]entity.User, error) {
	return nil, nil
}

// GetUserByEmailAndPassword fetches the user details
func (*UserRepository) GetUserByEmailAndPassword(*entity.User) (*entity.User, error) {
	return nil, nil
}
