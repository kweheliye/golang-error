package handler

import (
	"encoding/json"
	"errors"
	"golang-error/internal/service"
	"net/http"
)

type ApiError struct {
	Status  int
	Message string
}

func FromError(err error) ApiError {
	var svcError *service.ServiceError
	apiError := ApiError{
		Status:  http.StatusInternalServerError, // default
		Message: "internal server error",        // default
	}
	if errors.As(err, &svcError) {
		svcErr := svcError.SVCError() // the underlying error
		apiError.Message = svcErr.Error()

		switch {
		case errors.Is(svcErr, service.ErrBadRequest):
			apiError.Status = http.StatusBadRequest
		case errors.Is(svcErr, service.ErrInternalFailure):
			apiError.Status = http.StatusInternalServerError
		case errors.Is(svcErr, service.ErrNotFound):
			apiError.Status = http.StatusNotFound
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
