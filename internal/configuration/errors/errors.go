package errors

import "net/http"

type KasError struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

func (k *KasError) Error() string {
	return k.Message
}

func NewBadRequestError(message string) *KasError {
	return &KasError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *KasError {
	return &KasError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string) *KasError {
	return &KasError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *KasError {
	return &KasError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *KasError {
	return &KasError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}