package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kidkender/system-design-lab/internal/apperror"
)

type baseResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Success[T any](w http.ResponseWriter, status int, data T) {
	writeJSON(w, status, baseResponse[T]{Success: true, Data: data})
}

func Error(w http.ResponseWriter, err error) {
	var appErr *apperror.AppError
	if errors.As(err, &appErr) {
		writeJSON(w, appErr.Status, baseResponse[any]{
			Success: false,
			Error:   appErr.Message,
		})
		return
	}

	writeJSON(w, http.StatusInternalServerError, baseResponse[any]{
		Success: false,
		Error:   "internal server error",
	})
}
