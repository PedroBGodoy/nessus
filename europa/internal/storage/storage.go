package storage

import "github.com/nessus/europa/internal/models"

// UserRepository is
type UserRepository interface {
	Save(*models.User) (*models.User, error)
	FindByID(UserID string) (*models.User, error)
}
