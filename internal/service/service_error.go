package service

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
)

type ServiceError struct {
	svcErr error
	appErr error
	code   int
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("svcErr: %v,  code: %d", e.svcErr, e.code)
}

func (e *ServiceError) AppErr() error {
	return e.appErr
}

func (e *ServiceError) SvcErr() error {
	return e.svcErr
}

func NewServiceError(svcErr error, appErr error, code int) *ServiceError {
	return &ServiceError{
		svcErr: svcErr,
		appErr: appErr,
		code:   code,
	}
}

func NewNotFoundErr(appErr error) *ServiceError {
	return NewServiceError(ErrNotFound, appErr, http.StatusBadRequest)
}
