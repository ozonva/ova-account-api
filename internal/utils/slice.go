package utils

import (
	"errors"

	"github.com/ozonva/ova-account-api/internal/entity"
)

// ChunkSliceInt chunks slice on slices of the given size.
func ChunkSliceInt(s []int, size int) ([][]int, error) {
	if size < 1 {
		return nil, errors.New("the slice chunk size less than 1")
	}

	out := make([][]int, 0, (len(s)+size-1)/size)
	l := len(s)
	var i int
	for i = 0; i <= l-size; i += size {
		out = append(out, s[i:i+size])
	}

	if i < l {
		out = append(out, s[i:])
	}

	return out, nil
}

// ChunkSliceAccount chunks slice on slices of the given size.
func ChunkSliceAccount(s []entity.Account, size int) ([][]entity.Account, error) {
	if size < 1 {
		return nil, errors.New("the slice chunk size less than 1")
	}

	out := make([][]entity.Account, 0, (len(s)+size-1)/size)
	l := len(s)
	var i int
	for i = 0; i <= l-size; i += size {
		out = append(out, s[i:i+size])
	}

	if i < l {
		out = append(out, s[i:])
	}

	return out, nil
}

// FilterSliceString filters a slice of strings by the filter predicate.
func FilterSliceString(s []string, filter func(string) bool) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}

	return out
}

// ConvertAccountsToMap converts a slice of entity.Account to a map.
// key - Account.ID, value - entity.Account.
func ConvertAccountsToMap(s []entity.Account) map[string]entity.Account {
	out := make(map[string]entity.Account, len(s))
	for _, account := range s {
		out[account.ID] = account
	}

	return out
}
