package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// User define user structure
type User struct {
	gorm.Model
	UserID   string `gorm:"primaryKey"`
	Name     string `gorm:"index"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"index"`
}

// CreateUser create a new user
func CreateUser(name string, email string, password string, db *gorm.DB) (*User, error) {
	user := &User{
		UserID:   "789",
		Name:     name,
		Email:    email,
		Password: password,
	}

	result := db.Create(user)

	return user, result.Error
}

// LoginUser login user and returns JWT
func LoginUser(email string, password string) (string, error) {
	token, err := generateToken(email)
	if err != nil {
		log.Fatal("error when generating token: ", err)
	}

	return token, nil
}

// Authenticate validates user token
func Authenticate(token string, db *gorm.DB) (*User, error) {
	email, err := validatesToken(token)
	if err != nil {
		return nil, err
	}

	u := &User{}
	db.First(u, "email = ?", email)

	log.Printf("User: %s", u.UserID)

	if u.UserID == "" {
		return nil, errors.New("user not found")
	}

	return u, nil
}
