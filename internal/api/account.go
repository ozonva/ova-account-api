package api

import (
	pb "github.com/ozonva/ova-account-api/pkg/ova-account-api"
	"github.com/rs/zerolog"
)

// AccountService ...
type AccountService struct {
	pb.UnimplementedAccountServiceServer
	logger zerolog.Logger
}

// NewAccountService ...
func NewAccountService(logger zerolog.Logger) *AccountService {
	return &AccountService{logger: logger}
}
