package api

import (
	context "context"

	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/repo"
	"github.com/ozonva/ova-account-api/internal/utils"
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
func NewAccountService(logger zerolog.Logger, repo repo.Repo) *AccountService {
	return &AccountService{
		logger: logger.With().Str("service", "AccountService").Logger(),
		repo:   repo,
	}
}

func (s *AccountService) DescribeAccount(ctx context.Context, req *pb.DescribeAccountRequest) (*pb.DescribeAccountResponse, error) {
	s.logger.Info().Str("id", req.Id).Msg("RPC: DescribeAccount")

	account, err := s.repo.DescribeAccount(ctx, req.Id)
	if err != nil {
		return nil, wrapError(err)
	}

	return &pb.DescribeAccountResponse{Account: AccountMarshal(*account)}, nil
}

func (s *AccountService) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	s.logger.Info().Uint64("user_id", req.UserId).Msg("RPC: ListAccounts")

	accounts, err := s.repo.ListAccounts(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, wrapError(err)
	}

	return &pb.ListAccountsResponse{Accounts: AccountListMarshal(accounts)}, nil
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	s.logger.Info().Str("account", req.Value).Msg("RPC: CreateAccount")

	account, err := entity.NewAccount(req.UserId, req.Value)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = s.repo.AddAccounts(ctx, []entity.Account{*account})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CreateAccountResponse{Account: AccountMarshal(*account)}, nil
}

func (s *AccountService) MultiCreateAccount(ctx context.Context, req *pb.MultiCreateAccountRequest) (*emptypb.Empty, error) {
	s.logger.Info().Msg("RPC: MultiCreateAccount")

	accounts := make([]entity.Account, 0, len(req.Accounts))
	for _, a := range req.Accounts {
		acc, err := entity.NewAccount(a.UserId, a.Value)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		accounts = append(accounts, *acc)
	}

	batchSize := 32
	chunks, _ := utils.ChunkSliceAccount(accounts, batchSize)
	for _, chunk := range chunks {
		if err := s.createChunkAccounts(ctx, chunk); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *AccountService) createChunkAccounts(ctx context.Context, accounts []entity.Account) error {
	if err := s.repo.AddAccounts(ctx, accounts); err != nil {
		return err
	}
	// TODO: span here
	return nil
}

func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	s.logger.Info().Str("account", req.Account.Value).Msg("RPC: UpdateAccount")
	// TODO: add validation
	account := AccountUnmarshal(req.Account)
	err := s.repo.UpdateAccount(ctx, account)
	if err != nil {
		return nil, wrapError(err)
	}

	return &pb.UpdateAccountResponse{Account: AccountMarshal(account)}, nil
}

func (s *AccountService) RemoveAccount(ctx context.Context, req *pb.RemoveAccountRequest) (*emptypb.Empty, error) {
	s.logger.Info().Str("id", req.Id).Msg("RPC: RemoveAccount")

	if err := s.repo.RemoveAccount(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
