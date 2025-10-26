package store

import (
	"context"
	"database/sql"
	"golang-error/internal/model"
	"log"
	"time"
)

const (
	QueryTimeoutDuration = 2 * time.Second
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) GetByUserName(ctx context.Context, username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &model.User{}

	err := s.db.QueryRowContext(ctx, "SELECT username, email FROM users WHERE username=$1", username).
		Scan(
			&user.Username,
			&user.Email,
		)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SQLStore) ExecTx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	log.Printf("execting code in transaction")

	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
