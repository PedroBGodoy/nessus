package entity

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/nessus/europa/infra/security"
	"github.com/nessus/europa/infra/validation"
	"github.com/nessus/europa/internal/models"
)

// User represents de user model
type User struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string     `gorm:"size:100;not null;" json:"first_name"`
	LastName  string     `gorm:"size:100;not null;" json:"last_name"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Password  string     `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// CreateUser create a new user
func CreateUser(name string, email string, password string, db *gorm.DB) (*models.User, error) {
	if err := validation.UserName(name); err != nil {
		return nil, err
	}
	if err := validation.UserEmail(email); err != nil {
		return nil, err
	}
	if err := validation.UserPassword(password); err != nil {
		return nil, err
	}
	password, err := security.Hash(password)
	if err != nil {
		return nil, err
	}
	id := security.GenerateUUID()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ModelBase: models.ModelBase{
			ID: id,
		},
		Name:     name,
		Email:    email,
		Password: password,
	}

	result := db.Create(user)

	return user, result.Error
}

// LoginUser login user and returns JWT
func LoginUser(email string, password string) (string, error) {
	token, err := security.GenerateToken(email)
	if err != nil {
		log.Fatal("error when generating token: ", err)
	}

	return token, nil
}

// Authenticate validates user token
func Authenticate(token string, db *gorm.DB) (*models.User, error) {
	email, err := security.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	u := &models.User{}
	db.First(u, "email = ?", email)

	log.Printf("User: %s", u.ID)

	if u.ID == "" {
		return nil, errors.New("user not found")
	}

	return u, nil
}
