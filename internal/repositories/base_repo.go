package repositories

import (
	"context"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/db"
	"gorm.io/gorm"
)

// BaseRepository provides common functionality for all repositories
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository creates a new BaseRepository
func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{db: db}
}

// DB returns the database connection
func (r *BaseRepository) DB(ctx context.Context) *gorm.DB {
	if tx := db.TxFromCtx(ctx); tx != nil {
		return tx
	}
	return r.db
}
