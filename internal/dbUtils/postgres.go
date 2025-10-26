package dbutils

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// IsUniqueConstraintViolation returns true if the error is a Postgres unique constraint violation
func IsUniqueConstraintViolationError(err error) bool {
	if err == nil {
		return false
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// 23505 is the Postgres error code for unique_violation
		return pgErr.Code == "23505"
	}

	// Fall back to GORM's ErrDuplicatedKey for compatibility
	return errors.Is(err, gorm.ErrDuplicatedKey)
}
