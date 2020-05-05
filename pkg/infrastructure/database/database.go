package database

import (
	"fmt"
	"hello-go/pkg/infrastructure/config"

	"github.com/jinzhu/gorm"
	// Import postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Repository class
type Repository struct {
	db             *gorm.DB
	UserRepository *UserRepository
}

const dbDriver = "postgres"

// New creates an instance of gorm
func New(c *config.Config) (*Repository, error) {
	DBURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		c.Database.DBHost,
		c.Database.DBPort,
		c.Database.DBUser,
		c.Database.DBName,
		c.Database.DBPassword)

	db, err := gorm.Open(dbDriver, DBURL)
	if err != nil {
		fmt.Printf("Unable to connect to database, %v", err)
		defer db.Close()
		return nil, err
	}

	db.LogMode(true)

	userRepository := UserRepository{db}

	r := &Repository{db: db, UserRepository: &userRepository}
	return r, nil
}

// Migrate the database
func (r *Repository) Migrate() {
	r.db.AutoMigrate(&userModel{})
}
