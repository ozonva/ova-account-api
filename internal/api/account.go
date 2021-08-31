package api

import (
	context "context"

	"github.com/ozonva/ova-account-api/internal/repo"
	pb "github.com/ozonva/ova-account-api/pkg/ova-account-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AccountService ...
type AccountService struct {
	pb.UnimplementedAccountServiceServer
	logger zerolog.Logger
	repo   repo.Repo
}

// NewAccountService ...
func NewAccountService(logger zerolog.Logger) *AccountService {
	return &AccountService{
		logger: logger.With().Str("service", "AccountService").Logger(),
	}
}

func (s *AccountService) DescribeAccount(ctx context.Context, req *pb.DescribeAccountRequest) (*pb.DescribeAccountResponse, error) {
	s.logger.Info().Uint64("id", req.Id).Msg("RPC: DescribeAccount")

	return nil, status.Errorf(codes.Unimplemented, "method DescribeAccount not implemented")
}

func (s *AccountService) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	s.logger.Info().Uint64("user_id", req.UserId).Msg("RPC: ListAccounts")

	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	s.logger.Info().Str("account", req.Value).Msg("RPC: CreateAccount")

	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}

func (s *AccountService) RemoveAccount(ctx context.Context, req *pb.RemoveAccountRequest) (*emptypb.Empty, error) {
	s.logger.Info().Uint64("id", req.Id).Msg("RPC: RemoveAccount")

	return nil, status.Errorf(codes.Unimplemented, "method RemoveAccount not implemented")
}
