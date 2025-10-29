package apperrors

import "net/http"

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

// NewBadRequest returns a ServiceError with a status code of http.StatusBadRequest and a message
// It is used to indicate that the request was invalid or cannot be processed.
func NewBadRequest(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusBadRequest, Message: msg}
}

// NewConflict returns a ServiceError with a status code of http.StatusConflict and a message.
// It is used to indicate that the request resulted in a conflict with the current state of the resource.
// For example, if a user tries to create a resource with a name that is already in use, a conflict error should be returned.
func NewConflict(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusConflict, Message: msg}
}

// NewInternalError returns a ServiceError with a status code of http.StatusInternalServerError and a message.
// It is used to indicate that an internal error occurred.
func NewInternalError(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusInternalServerError, Message: msg}
}

// NewUnauthorized returns a ServiceError with a status code of http.StatusUnauthorized and a message.
// It is used to indicate that the request was unauthorized.
func NewUnauthorized(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusUnauthorized, Message: msg}
}

// NewNotFound returns a ServiceError with a status code of http.StatusNotFound and a message.
// It is used to indicate that the requested resource was not found.
func NewNotFound(msg string) *ServiceError {
	return &ServiceError{Code: http.StatusNotFound, Message: msg}
}
