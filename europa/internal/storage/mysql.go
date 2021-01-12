package storage

import "github.com/nessus/europa/internal/models"

// MySQLUserRepo is
type MySQLUserRepo struct {
	Host string
	Port int16
}

func (*MySQLUserRepo) Connect() error {
	return nil
}

// Save is
func (*MySQLUserRepo) Save(*models.User) (*models.User, error) {
	return &models.User{}, nil
}
