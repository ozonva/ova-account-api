package entity

import "fmt"

func CreateTestAccounts(count int) []Account {
	out := make([]Account, 0, count)

	for i := 0; i < count; i++ {
		account, _ := NewAccount(1, fmt.Sprintf("user%d@ozon.ru", i+1))
		out = append(out, *account)
	}

	return out
}
