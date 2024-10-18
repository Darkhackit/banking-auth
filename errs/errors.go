package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
