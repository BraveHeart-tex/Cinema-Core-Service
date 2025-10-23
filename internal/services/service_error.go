package services

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
