package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/repo"
)

// Check implementation
var _ repo.Repo = &accountRepo{}

type accountRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *accountRepo {
	return &accountRepo{db: db}
}

func (r *accountRepo) DescribeAccount(ctx context.Context, id string) (*entity.Account, error) {
	acc := &entity.Account{}
	err := r.db.GetContext(ctx, acc, "SELECT id, value, user_id FROM accounts where id = $1 LIMIT 1", id)

	return acc, repo.DBError(err)
}

func (r *accountRepo) AddAccounts(ctx context.Context, accounts []entity.Account) error {
	_, err := r.db.NamedExecContext(ctx, `INSERT INTO accounts (id, value, user_id) VALUES (:id, :value, :user_id)`, accounts)

	return err
}

func (r *accountRepo) ListAccounts(ctx context.Context, limit, offset uint64) ([]entity.Account, error) {
	var accounts []entity.Account
	err := r.db.SelectContext(ctx, &accounts, "SELECT id, value, user_id FROM accounts LIMIT $1 OFFSET $2", limit, offset)

	return accounts, err
}

func (r *accountRepo) UpdateAccount(ctx context.Context, account entity.Account) error {
	result, err := r.db.ExecContext(
		ctx,
		"UPDATE accounts SET value = $2, user_id = $3, updated_at = $4 where id = $1",
		account.ID,
		account.Value,
		account.UserID,
		time.Now())

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repo.ErrRecordNotFound
	}

	return nil
}

func (r *accountRepo) RemoveAccount(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM accounts WHERE id = $1", id)

	return err
}
