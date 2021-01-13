package models

// User define user structure
type User struct {
	ModelBase
	Name     string `gorm:"index"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"index"`
}
