package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/repo"
)

type accountRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *accountRepo {
	return &accountRepo{db: db}
}

func (r *accountRepo) DescribeAccount(id uint64) (*entity.Account, error) {
	acc := &entity.Account{}
	err := r.db.GetContext(context.TODO(), acc, "SELECT id, value, user_id FROM accounts where id = $1 LIMIT 1", id)

	return acc, repo.DBError(err)
}

func (r *accountRepo) AddAccounts(accounts []entity.Account) error {
	_, err := r.db.NamedExecContext(context.TODO(), `INSERT INTO accounts (value, user_id) VALUES (:value, :user_id)`, accounts)

	return err
}

func (r *accountRepo) ListAccounts(limit, offset uint64) ([]entity.Account, error) {
	var accounts []entity.Account
	err := r.db.SelectContext(context.TODO(), &accounts, "SELECT id, value, user_id FROM accounts LIMIT $1 OFFSET $2", limit, offset)

	return accounts, err
}

func (r *accountRepo) RemoveAccount(id uint64) error {
	_, err := r.db.ExecContext(context.TODO(), "DELETE FROM accounts WHERE id = $1", id)

	return err
}
