package session

import (
	"context"

	"gorm.io/gorm"
)

// WithDatabase is
func WithDatabase(ctx context.Context, db *gorm.DB) (context.Context, error) {
	newCtx := context.WithValue(ctx, gorm.DB{}, db)
	db.WithContext(newCtx)
	return newCtx, nil
}

// GetDatabase is
func GetDatabase(ctx context.Context) *gorm.DB {
	return ctx.Value(gorm.DB{}).(*gorm.DB)
}
