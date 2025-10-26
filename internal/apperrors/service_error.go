package apperrors

import "net/http"

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewBadRequest(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusBadRequest, Message: msg}
}

func NewConflict(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusConflict, Message: msg}
}

func NewInternalError(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusInternalServerError, Message: msg}
}

func NewUnauthorized(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusUnauthorized, Message: msg}
}

func NewNotFound(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusNotFound, Message: msg}
}
