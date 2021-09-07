package mocks

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/ozonva/ova-account-api/internal/entity"
)

// AccountValueEq returns a matcher that matches on equality Account values.
//
// Example usage:
//   AccountValueEq(entity.Account{u}).Matches(5) // returns true
//   AccountValueEq(5).Matches(4) // returns false
func AccountValueEq(x []entity.Account) gomock.Matcher { return accountsMatcher{x} }

type accountsMatcher struct {
	accounts []entity.Account
}

// Matches returns whether x is a match.
func (m accountsMatcher) Matches(x interface{}) bool {
	actual, ok := x.([]entity.Account)
	if !ok {
		return false
	}

	if len(m.accounts) != len(actual) {
		return false
	}

	for i, v := range m.accounts {
		if v.Value != actual[i].Value {
			return false
		}
		if v.UserID != actual[i].UserID {
			return false
		}
	}

	return true
}

// String describes what the matcher matches.
func (m accountsMatcher) String() string {
	return fmt.Sprintf("is equal to %v", m.accounts)
}
