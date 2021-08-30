package api

import (
	"net"

	"google.golang.org/grpc"
)

// Server represents gRPC server to serve RPC requests.
type Server struct {
	server *grpc.Server
	errors chan error
	port   string
}

// NewServer creates a new grpc server.
func NewServer(port string) *Server {
	server := grpc.NewServer()

	return &Server{
		server: server,
		errors: make(chan error, 1),
		port:   port,
	}
}

// Start starts the gRPC server.
func (s *Server) Start()  {
	go func() {
		listener, err := net.Listen("tcp", net.JoinHostPort("", s.port))
		if err != nil {
			s.errors <- err
			return
		}

		s.errors <- s.server.Serve(listener)
		close(s.errors)
	}()
}

// Stop stops the gRPC server.
func (s *Server) Stop() {
	s.server.GracefulStop()
}

// Notify returns a channel to notify the caller about errors.
// If you receive an error from the channel you should stop the application.
func (s *Server) Notify() <-chan error {
	return s.errors
}
