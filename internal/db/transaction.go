package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// TxManager defines the interface for transaction management
type TxManager interface {
	// WithTransaction executes the given function within a transaction
	// If the function returns an error, the transaction is rolled back
	// Otherwise, the transaction is committed
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}

// GormTxManager implements TxManager using GORM
type GormTxManager struct {
	db *gorm.DB
}

// NewGormTxManager creates a new GormTxManager
func NewGormTxManager(db *gorm.DB) *GormTxManager {
	return &GormTxManager{db: db}
}

// WithTransaction executes the given function within a transaction
func (m *GormTxManager) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	defer func() {
		// Recover from panic
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}