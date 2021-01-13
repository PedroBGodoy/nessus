package models

import (
	"time"

	"gorm.io/gorm"
)

// ModelBase is
type ModelBase struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
