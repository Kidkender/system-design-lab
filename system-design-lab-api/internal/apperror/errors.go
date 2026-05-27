package apperror

import "net/http"

type AppError struct {
	Status  int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFound(msg string) *AppError {
	return &AppError{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

func InternalServerError(msg string) *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

func BadRequest(msg string) *AppError {
	return &AppError{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}
