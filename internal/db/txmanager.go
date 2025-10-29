package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type txKey struct{}

func WithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func TxFromCtx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

// TxManager defines the interface for transaction management
type TxManager interface {
	// WithTransaction executes the given function within a transaction
	// If the function returns an error, the transaction is rolled back
	// Otherwise, the transaction is committed
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// GormTxManager implements TxManager using GORM
type GormTxManager struct {
	db *gorm.DB
}

// NewGormTxManager creates a new GormTxManager
func NewGormTxManager(db *gorm.DB) *GormTxManager {
	return &GormTxManager{db: db}
}

// WithTransaction ensures a transaction exists. Reuses context transaction if present.
func (m *GormTxManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if tx := TxFromCtx(ctx); tx != nil {
		return fn(ctx)
	}

	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // propagate panic
		}
	}()

	ctxWithTx := WithTx(ctx, tx)

	if err := fn(ctxWithTx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
