package repo

import (
	"database/sql"
	"errors"
)

// ErrRecordNotFound ...
var ErrRecordNotFound = errors.New("record not found")

// Store ...
type Store interface {
	Account() Repo
	Close() error
}

// DBError wraps sql.ErrNoRows.
func DBError(err error) error {
	if err == sql.ErrNoRows {
		return ErrRecordNotFound
	}

	return err
}
