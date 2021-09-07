package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Event interface {
	kafkaMessage() kafka.Message
}

// Producer ...
type Producer interface {
	Send(ctx context.Context, messages ...Event) error
	Close() error
	Health() error
}

type producer struct {
	w *kafka.Writer
}

func NewProducer(addr, topic string) Producer {
	return &producer{
		w: &kafka.Writer{
			Addr:  kafka.TCP(addr),
			Topic: topic,
			Async: true,
		},
	}
}

func (p *producer) Send(ctx context.Context, messages ...Event) error {
	err := p.w.WriteMessages(ctx, convertEvents(messages)...)
	if err != nil {
		return fmt.Errorf("failed to send messages to kafka: %w", err)
	}

	return nil
}

func (p *producer) Close() error {
	return p.w.Close()
}

func (p *producer) Health() error {
	// TODO: Add a health check
	return nil
}

func convertEvents(events []Event) []kafka.Message {
	out := make([]kafka.Message, 0, len(events))
	for _, event := range events {
		out = append(out, event.kafkaMessage())
	}

	return out
}
