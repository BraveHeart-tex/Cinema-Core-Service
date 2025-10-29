package repositories

import (
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
func (r *BaseRepository) DB() *gorm.DB {
	return r.db
}

// WithTx returns a new repository with the given transaction
func (r *BaseRepository) WithTx(tx *gorm.DB) {
	r.db = tx
}