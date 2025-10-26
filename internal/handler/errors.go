package handler

import (
	"fmt"
	"net/http"
)

// APIError is the standard format for API error responses.
type APIError struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

// Error satisfies the error interface.
func (e APIError) Error() string {
	return fmt.Sprintf("status %d: %v", e.StatusCode, e.Message)
}

// NewAPIError creates a new APIError.
func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

// InvalidRequestData creates a 400 Bad Request error with a map of validation errors.
func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Message:    errors,
	}
}

// APIFunc is a custom handler function type that returns an error.
type APIFunc func(w http.ResponseWriter, r *http.Request) error

// Make converts an APIFunc into a standard http.HandlerFunc.
func Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				WriteJson(w, apiErr.StatusCode, apiErr)
			} else {
				// For non-APIError types, return a generic 500.
				errResp := NewAPIError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
				WriteJson(w, http.StatusInternalServerError, errResp)
			}
		}
	}
}
