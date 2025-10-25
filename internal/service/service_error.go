package service

import "errors"

var (
	ErrInternalFailure = errors.New("internal failure")
	ErrNotFound        = errors.New("not found")
	ErrBadRequest      = errors.New("bad request")
)

type ServiceError struct {
	Code     int
	appError error
	svcError error
}

func (e ServiceError) Error() string {
	return e.svcError.Error()
}

func NewError(code int, svcError error, appError error) *ServiceError {
	return &ServiceError{
		Code:     code,
		svcError: svcError,
		appError: appError,
	}
}

func (e ServiceError) AppError() error {
	return e.appError
}

func (e ServiceError) SVCError() error {
	return e.svcError
}

func NewBadRequest(svcError error, appError error) *ServiceError {
	return NewError(400001, svcError, appError)
}

func NewNotFound(svcError error, appError error) *ServiceError {
	return NewError(50001, svcError, appError)
}

func NewInternalError(svcError error, appError error) *ServiceError {
	return NewError(60001, svcError, appError)
}
