package postgres

import (
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-account-api/internal/repo"
)

// Store represents a postgres store
type Store struct {
	db      *sqlx.DB
	account *accountRepo
}

func NewStore(dsn string) (*Store, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// TODO: move to configuration
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Minute)

	return &Store{db: db}, nil
}

func (s *Store) Account() repo.Repo {
	if s.account == nil {
		s.account = NewRepo(s.db)
	}

	return s.account
}

func (s *Store) Close() error {
	return s.db.Close()
}
