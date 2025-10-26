package handler

import (
	"encoding/json"
	"errors"
	"golang-error/internal/service"
	"net/http"
)

type ApiError struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
}

func FromError(err error) ApiError {
	var serviceError *service.ServiceError
	apiError := ApiError{
		Status:  http.StatusInternalServerError, // default
		Message: "internal server error",        // default
	}

	if errors.As(err, &serviceError) {
		svcErr := serviceError.SvcErr()
		apiError.Message = svcErr.Error()

		switch {
		case errors.Is(svcErr, service.ErrNotFound):
			apiError.Status = http.StatusBadRequest
		}
	}
	return apiError
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	apiErr := FromError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.Status)
	json.NewEncoder(w).Encode(apiErr)
}
