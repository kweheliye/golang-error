package repository

import (
	"context"
	"database/sql"
	"golang-error/internal/domain"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, "SELECT  username, email FROM users WHERE username=$1", username).
		Scan(
			&user.Username,
			&user.Email,
		)
	if err != nil {
		return nil, err
	}
	return user, nil
}
