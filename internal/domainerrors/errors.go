package domainerrors

import "errors"

var (
	ErrNotFound = errors.New("record not found")
	ErrConflict = errors.New("conflict")
	ErrInvalid  = errors.New("invalid input")
	ErrInternal = errors.New("internal error")
)
