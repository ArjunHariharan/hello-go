package database

import (
	"fmt"
	"hello-go/pkg/config"
	"hello-go/pkg/domain/entity"

	"github.com/jinzhu/gorm"
	// Import postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Repository class
type Repository struct {
	db *gorm.DB
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

	r := &Repository{db: db}
	return r, nil
}

// Migrate the database
func (r *Repository) Migrate() {
	r.db.AutoMigrate(&entity.User{})
}
