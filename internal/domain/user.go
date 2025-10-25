package domain

import "context"

type User struct {
	Username string
	Email    string
}

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
}
