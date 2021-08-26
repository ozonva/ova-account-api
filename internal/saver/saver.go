package saver

import (
	"sync"
	"time"

	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/flusher"
)

// Saver ...
type Saver interface {
	Save(entity entity.Account)
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
func (s *saver) Save(account entity.Account) {
	if len(s.buffer) == cap(s.buffer) {
		s.flush()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.buffer = append(s.buffer, account)
}

// Init inits ...
func (s *saver) Init() {
	s.done = make(chan struct{}, 1)

	go func() {
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
