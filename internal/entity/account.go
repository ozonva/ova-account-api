package entity

import (
	"net"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

// Account represents the note's category.
type Account struct {
	ID     string
	UserID uint64 `db:"user_id"`
	Value  string
}

// NewAccount creates a new entity Account.
func NewAccount(userID uint64, value string) (*Account, error) {
	address, err := mail.ParseAddress(value)
	if err != nil {
		return nil, err
	}

	return &Account{ID: uuid.New().String(), UserID: userID, Value: address.Address}, nil
}

// // Value returns the Value of the Account.
// func (a Account) Value() string {
// 	return a.Value
// }

// Domain returns the domain of the Account.
func (a Account) Domain() string {
	at := strings.LastIndex(a.Value, "@")
	if at == -1 {
		return ""
	}

	return a.Value[at+1:]
}

// Username returns the username of the Account.
func (a Account) Username() string {
	at := strings.LastIndex(a.Value, "@")
	if at == -1 {
		return ""
	}

	return a.Value[:at]
}

// Exists checks if account exists on the internet.
func (a Account) Exists() bool {
	// TODO: add more checks.
	if _, err := net.LookupMX(a.Domain()); err != nil {
		return false
	}

	return true
}
