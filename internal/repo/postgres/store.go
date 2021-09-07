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

func NewStore(dsn string, maxOpenConns, maxIdleConns, ConnMaxLifetime int) (*Store, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(ConnMaxLifetime) * time.Second)

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

func (s *Store) Health() error {
	var err error
	for i := 0; i < 5; i++ {
		_, err = s.db.Query("SELECT 1")
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	return err
}
