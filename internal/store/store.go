package store

import (
	"context"
	"database/sql"
	"golang-error/internal/model"
)

type Store interface {
	GetByUserName(ctx context.Context, username string) (*model.User, error)
	ExecTx(ctx context.Context, fn func(tx *sql.Tx) error) error
}
