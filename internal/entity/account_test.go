package entity_test

import (
	"testing"

	"github.com/ozonva/ova-account-api/internal/entity"
)

func TestAccount(t *testing.T) {
	account, err := entity.NewAccount(1, "vyacheslavv@ozon.ru")
	assertNoError(t, err)
	assertEqualString(t, "ozon.ru", account.Domain())
	assertEqualString(t, "vyacheslavv", account.Username())
	assertTrue(t, account.Exists(), "account.Exists()")
}

func TestAccountInvalidAddress(t *testing.T) {
	_, err := entity.NewAccount(1, "vyacheslavv")
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestAccountNonExistentDomain(t *testing.T) {
	domain := "megaozon.ru"
	account, err := entity.NewAccount(1, "vyacheslavv@"+domain)
	assertNoError(t, err)
	if account.Exists() {
		t.Fatalf("Expected domain %s does not exist.", domain)
	}
}

func assertTrue(t *testing.T, value bool, msg ...string) {
	t.Helper()
	if !value {
		t.Fatal("Should be true", msg)
	}
}

func assertEqualString(t *testing.T, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Received unexpected error:\n%+v", err)
	}
}
