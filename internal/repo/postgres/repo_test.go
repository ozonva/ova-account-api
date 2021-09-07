package postgres

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func newMock(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	return sqlx.NewDb(db, "pgx"), mock
}

func generateAccounts(count int) []entity.Account {
	out := make([]entity.Account, 0, count)

	for i := 0; i < count; i++ {
		account, _ := entity.NewAccount(1, fmt.Sprintf("user%d@ozon.ru", i+1))
		out = append(out, *account)
	}

	return out
}

func Test_accountRepo_ListAccounts(t *testing.T) {
	ctx := context.Background()
	db, mock := newMock(t)
	defer db.Close()
	repo := NewRepo(db)

	rows := sqlmock.NewRows([]string{"id", "value", "user_id"})
	accounts := generateAccounts(2)
	for _, account := range accounts {
		rows.AddRow(account.ID, account.Value, account.UserID)
	}

	mock.ExpectQuery("SELECT id, value, user_id FROM accounts LIMIT \\$1 OFFSET \\$2").
		WithArgs(11, 64).
		WillReturnRows(rows)

	result, err := repo.ListAccounts(ctx, 11, 64)
	assert.NoError(t, err)
	assert.Equal(t, accounts, result)
}

func Test_accountRepo_DescribeAccount(t *testing.T) {
	ctx := context.Background()
	db, mock := newMock(t)
	defer db.Close()
	repo := NewRepo(db)

	account := generateAccounts(1)[0]
	rows := sqlmock.NewRows([]string{"id", "value", "user_id"}).
		AddRow(account.ID, account.Value, account.UserID)

	mock.ExpectQuery("SELECT id, value, user_id FROM accounts where id = \\$1 LIMIT 1").
		WithArgs(account.ID).
		WillReturnRows(rows)

	result, err := repo.DescribeAccount(ctx, account.ID)
	assert.NoError(t, err)
	assert.Equal(t, account, *result)
}

func Test_accountRepo_AddAccounts(t *testing.T) {
	ctx := context.Background()
	db, mock := newMock(t)
	defer db.Close()
	repo := NewRepo(db)

	accounts := generateAccounts(4)
	var args []driver.Value
	for _, acc := range accounts {
		args = append(args, acc.ID, acc.Value, acc.UserID)
	}

	mock.ExpectExec("INSERT INTO accounts \\(id, value, user_id\\) VALUES").
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(4, 4))

	err := repo.AddAccounts(ctx, accounts)
	assert.NoError(t, err)
}

func Test_accountRepo_RemoveAccount(t *testing.T) {
	ctx := context.Background()
	db, mock := newMock(t)
	defer db.Close()
	repo := NewRepo(db)

	account := generateAccounts(1)[0]

	mock.ExpectExec("DELETE FROM accounts WHERE id = \\$1").
		WithArgs(account.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.RemoveAccount(ctx, account.ID)
	assert.NoError(t, err)
}

func Test_accountRepo_RemoveAccountWithError(t *testing.T) {
	ctx := context.Background()
	db, mock := newMock(t)
	defer db.Close()
	repo := NewRepo(db)

	account := generateAccounts(1)[0]

	err := errors.New("something went wrong")
	mock.ExpectExec("DELETE FROM accounts WHERE id = \\$1").
		WithArgs(account.ID).
		WillReturnError(err)

	resultErr := repo.RemoveAccount(ctx, account.ID)
	assert.ErrorIs(t, err, resultErr)
}
