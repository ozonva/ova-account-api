package saver

import (
	"errors"
	"sync"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/flusher"
)

var (
	// ErrFullBufferFlush is type of error returned when the buffer is full and cannot be flushed.
	ErrFullBufferFlush = errors.New("failed to flush the full buffer")
)

// Saver ...
type Saver interface {
	Save(entity entity.Account) error
	Init()
	Close()
}

// NewSaver creates Saver with support for periodic saving.
func NewSaver(capacity uint, flusher flusher.Flusher, timeout time.Duration) Saver {
	return &saver{
		flusher: flusher,
		buffer:  make([]entity.Account, 0, capacity),
		timeout: timeout,
		done:    make(chan struct{}, 1),
	}
}

type saver struct {
	mu      sync.Mutex
	flusher flusher.Flusher
	buffer  []entity.Account
	timeout time.Duration
	done    chan struct{}
}

// Save adds the account to the list for saving.
func (s *saver) Save(account entity.Account) error {
	if len(s.buffer) == cap(s.buffer) {
		s.flush()
	}

	if len(s.buffer) == cap(s.buffer) {
		return ErrFullBufferFlush
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.buffer = append(s.buffer, account)
	return nil
}

// Init inits ...
func (s *saver) Init() {
	s.done = make(chan struct{}, 1)

	go func() {
		defer ginkgo.GinkgoRecover()
		ticker := time.NewTicker(s.timeout)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.flush()
			case <-s.done:
				s.flush()
				return
			}
		}
	}()
}

// Close ...
func (s *saver) Close() {
	s.done <- struct{}{}
}

func (s *saver) flush() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.buffer) == 0 {
		return
	}

	unsaved := s.flusher.Flush(s.buffer)
	s.buffer = s.buffer[:0]
	s.buffer = append(s.buffer, unsaved...)
}
