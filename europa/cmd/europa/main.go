package main

import (
	"context"
	"log"

	"github.com/nessus/europa/internal/models"
	"github.com/nessus/europa/internal/server/auth"
	"github.com/nessus/europa/internal/session"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	port = ":50051"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	ctx, err := session.WithDatabase(context.Background(), db)
	if err != nil {
		log.Fatalf("error when initalizing context: %s", err)
	}

	_, err = auth.NewServer(ctx, port)
}
