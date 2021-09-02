package repo

import "github.com/ozonva/ova-account-api/internal/entity"

// Repo represents the storage interface for the entity.Account.
type Repo interface {
	AddAccounts(entities []entity.Account) error
	ListAccounts(limit, offset uint64) ([]entity.Account, error)
	DescribeAccount(entityID uint64) (*entity.Account, error)
	RemoveAccount(id uint64) error
}
