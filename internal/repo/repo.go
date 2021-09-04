package repo

import (
	"context"

	"github.com/ozonva/ova-account-api/internal/entity"
)

// Repo represents the storage interface for the entity.Account.
type Repo interface {
	AddAccounts(ctx context.Context, entities []entity.Account) error
	ListAccounts(ctx context.Context, limit, offset uint64) ([]entity.Account, error)
	DescribeAccount(ctx context.Context, entityID uint64) (*entity.Account, error)
	RemoveAccount(ctx context.Context, id uint64) error
}
