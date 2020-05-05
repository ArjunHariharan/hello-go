package entity

// UserProfile entity
type UserProfile struct {
	ID         int
	FirstName  string `gorm:"type:varchar(64)"`
	MiddleName string `gorm:"type:varchar(64)"`
	LastName   string `gorm:"type:varchar(64)"`
	Email      string `gorm:"type:varchar(256);unique_index"`
}
