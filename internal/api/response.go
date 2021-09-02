package api

import (
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/repo"
	pb "github.com/ozonva/ova-account-api/pkg/ova-account-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func wrapError(err error) error {
	if err == repo.ErrRecordNotFound {
		return status.Errorf(codes.NotFound, err.Error())
	}

	return status.Errorf(codes.Internal, err.Error())
}

func AccountMarshal(account entity.Account) *pb.Account {
	return &pb.Account{
		Id:     account.ID,
		Value:  account.Value,
		UserId: account.UserID,
	}
}

func AccountListMarshal(accounts []entity.Account) []*pb.Account {
	out := make([]*pb.Account, 0, len(accounts))
	for _, account := range accounts {
		out = append(out, AccountMarshal(account))
	}

	return out
}

func AccountUnmarshal(account *pb.Account) entity.Account {
	return entity.Account{
		ID:     account.Id,
		Value:  account.Value,
		UserID: account.UserId,
	}
}

func AccountListUnmarshal(accounts []*pb.Account) []entity.Account {
	out := make([]entity.Account, 0, len(accounts))
	for _, account := range accounts {
		out = append(out, AccountUnmarshal(account))
	}

	return out
}
