package api

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/ozonva/ova-account-api/internal/app"
	pb "github.com/ozonva/ova-account-api/pkg/ova-account-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server represents gRPC server to serve RPC requests.
type Server struct {
	server *grpc.Server
	errors chan error
	port   string
}

// NewServer creates a new grpc server.
func NewServer(application *app.App) *Server {
	recoveryHandler := func(p interface{}) (err error) {
		log.Error().Msgf("panic triggered: %v", p)
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryHandler)),
		),
	)

	accountService := NewAccountService(
		application.Store.Account(),
		application.Producer,
		application.Metrics,
		application.Conf.DB.BatchSize,
	)
	pb.RegisterAccountServiceServer(server, accountService)

	return &Server{
		server: server,
		errors: make(chan error, 1),
		port:   application.Conf.GrpcPort,
	}
}

// Start starts the gRPC server.
func (s *Server) Start() {
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
