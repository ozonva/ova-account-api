package health

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// Server is a health server.
type Server struct {
	server http.Server
	errors chan error
}

// NewServer ...
func NewServer(port string, check Check) *Server {
	http.HandleFunc("/health", healthHandler(check))

	return &Server{
		server: http.Server{
			Addr:    net.JoinHostPort("", port),
			Handler: nil,
		},
		errors: make(chan error, 1),
	}
}

// Start ...
func (s *Server) Start() {
	go func() {
		s.errors <- s.server.ListenAndServe()
		close(s.errors)
	}()
}

// Stop ...
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// Notify ...
func (s *Server) Notify() <-chan error {
	return s.errors
}

func healthHandler(check Check) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resources := check()

		err := json.NewEncoder(w).Encode(resources)
		if err != nil {
			log.Error().Err(err).Msgf("Couldn't encode resources: %v", resources)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
