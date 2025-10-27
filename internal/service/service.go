package service

import (
	"context"
	"database/sql"
	"golang-error/internal/model"
	"golang-error/internal/store"
)

type Service struct {
	store store.Store
}

func NewService(s store.Store) *Service {
	return &Service{store: s}
}

func (u *Service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User
	err := u.store.ExecTx(ctx, func(tx *sql.Tx) error {
		var err error
		user, err = u.store.GetByUserName(ctx, tx, username)
		return err
	})

	if err != nil {
		return nil, NewNotFoundErr(err)
	}
	return user, nil
}
