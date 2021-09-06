package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server is a metrics server.
type Server struct {
	server http.Server
	errors chan error
}

func NewServer() *Server {
	http.Handle("/metrics", promhttp.Handler())

	return &Server{
		server: http.Server{
			Addr:    "localhost:9002",
			Handler: nil,
		},
		errors: make(chan error, 1),
	}
}

func (s *Server) Start() {
	go func() {
		s.errors <- s.server.ListenAndServe()
		close(s.errors)
	}()
}

// Stop metrics server.
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// Notify returns a channel to notify the caller about errors.
// If you receive an error from the channel diagnostic you should stop the application.
func (s *Server) Notify() <-chan error {
	return s.errors
}
