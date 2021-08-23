package flusher

import (
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/repo"
	"github.com/ozonva/ova-account-api/internal/utils"
)

// Flusher represents an interface for flushing entity.Account to the repository.
type Flusher interface {
	Flush(entities []entity.Account) []entity.Account
}

// NewFlusher creates Flusher with support for batch saving.
func NewFlusher(chunkSize int, accountRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		repo:      accountRepo,
	}
}

type flusher struct {
	chunkSize int
	repo      repo.Repo
}

// Flush flushes the list of entity.Account to the storage.
func (f *flusher) Flush(accounts []entity.Account) []entity.Account {
	chunks, err := utils.ChunkSliceAccount(accounts, f.chunkSize)
	if err != nil {
		return accounts
	}

	var unsaved []entity.Account
	for _, chunk := range chunks {
		err := f.repo.AddAccounts(chunk)
		if err != nil {
			unsaved = append(unsaved, chunk...)
		}
	}

	return unsaved
}
