package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User entity
type User struct {
	ID        int
	UUID      uuid.UUID  `gorm:"type:uuid"`
	Email     string     `gorm:"type:varchar(256);unique_index"`
	Password  string     `gorm:"type:varchar(72)"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// New returns an instance of User
func New(e string, p string) *User {
	u := uuid.NewV4()

	return &User{Email: e, Password: p, UUID: u}
}
