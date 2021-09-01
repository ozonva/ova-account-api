package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-account-api/internal/entity"
)

type accountRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *accountRepo {
	return &accountRepo{db: db}
}

func (r *accountRepo) DescribeAccount(id uint64) (*entity.Account, error) {
	panic("implement me")
}

func (r *accountRepo) AddAccounts(accounts []entity.Account) error {
	panic("implement me")
}

func (r *accountRepo) ListAccounts(limit, offset uint64) ([]entity.Account, error) {
	panic("implement me")
}
