package kafka

import (
	"fmt"

	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/segmentio/kafka-go"
)

// AccountEventType ...
type AccountEventType string

const (
	AccountCreated AccountEventType = "AccountCreated"
	AccountUpdated AccountEventType = "AccountUpdated"
	AccountRemoved AccountEventType = "AccountRemoved"
)

func (t AccountEventType) Bytes() []byte {
	return []byte(t)
}

type AccountEvent struct {
	Type    AccountEventType
	Account entity.Account
}

func NewAccountEvent(t AccountEventType, account entity.Account) AccountEvent {
	return AccountEvent{
		Type:    t,
		Account: account,
	}
}

func NewAccountEvents(t AccountEventType, accounts []entity.Account) []Event {
	out := make([]Event, 0, len(accounts))
	for _, account := range accounts {
		out = append(out, NewAccountEvent(t, account))
	}

	return out
}

func (e AccountEvent) kafkaMessage() kafka.Message {
	// TODO:
	value := fmt.Sprintf(`{"id":"%s", "value":"%s", "user_id":%d}`, e.Account.ID, e.Account.Value, e.Account.UserID)

	return kafka.Message{
		Key:   e.Type.Bytes(),
		Value: []byte(value),
	}
}
