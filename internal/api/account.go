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

func (s *AccountService) DescribeAccount(ctx context.Context, id *pb.ID) (*pb.Account, error) {
	s.logger.Info().Uint64("id", id.Id).Msg("RPC: DescribeAccount")

	return nil, status.Errorf(codes.Unimplemented, "method DescribeAccount not implemented")
}

func (s *AccountService) ListAccounts(ctx context.Context, request *pb.ListAccountsRequest) (*pb.AccountsList, error) {
	s.logger.Info().Uint64("user_id", request.UserId).Msg("RPC: ListAccounts")

	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}

func (s *AccountService) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.Account, error) {
	s.logger.Info().Str("account", request.Value).Msg("RPC: CreateAccount")

	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}

func (s *AccountService) RemoveAccount(ctx context.Context, id *pb.ID) (*emptypb.Empty, error) {
	s.logger.Info().Uint64("id", id.Id).Msg("RPC: RemoveAccount")

	return nil, status.Errorf(codes.Unimplemented, "method RemoveAccount not implemented")
}
