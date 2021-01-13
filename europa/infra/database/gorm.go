package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/hashicorp/go-hclog"
	"github.com/nessus/europa/internal/models"
)

// InitSQLite is
func InitSQLite(logger hclog.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	if db.Error != nil || err != nil {
		logger.Error("error when trying to connect to database")
		os.Exit(1)
	}

	return db
}
