package response

import (
	"codename-rl/internal/model"

	"encoding/json"
	"net/http"
	"time"
)

// Success response helper
func NewResponse[T any](message string, data T) *model.BaseResponse[T] {
	return &model.BaseResponse[T]{
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
	}
}

// Error response helper
func NewErrorResponse(message string, err error) *model.BaseResponse[any] {
	return &model.BaseResponse[any]{
		Message:   message,
		Errors:    err.Error(),
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
	}
}

// Writer helper for direct HTTP output
func WriteJSON[T any](w http.ResponseWriter, status int, resp *model.BaseResponse[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}
