package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-error/internal/domain"
)

// Specific errors returned by the service layer.
var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{userRepo: r}
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*domain.User, error) {

	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		// If the database returns no rows, we return our specific ErrUserNotFound.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewNotFound(fmt.Errorf("%w: user %s", ErrNotFound, username), err)
		}
		// For any other database error, we return a generic internal error.
		return nil, NewInternalError(ErrInternalFailure, err)
	}

	return user, nil
}
