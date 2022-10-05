package utils

import "net/http"

type AppError struct {
	Code    int `json:",omitempty"`
	Message string
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}
func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func (err AppError) GetMessage() *AppError {
	return &AppError{
		Message: err.Message,
	}
}
