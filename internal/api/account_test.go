package api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/mocks"
	"github.com/ozonva/ova-account-api/internal/repo"
	pb "github.com/ozonva/ova-account-api/pkg/ova-account-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAccountService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Account Service Suite")
}

var _ = Describe("Account Service", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		producer *mocks.MockProducer
		metrics  *mocks.MockAccountMetrics
		service  *AccountService
		ctx      context.Context
		batchSize = 32
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		producer = mocks.NewMockProducer(ctrl)
		metrics = mocks.NewMockAccountMetrics(ctrl)
		service = NewAccountService(zerolog.Logger{}, mockRepo, producer, metrics, batchSize)
		ctx = context.TODO()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("CreateAccount", func() {
		Context("when adding an valid account", func() {
			It("should successfully store", func() {
				req := &pb.CreateAccountRequest{Value: "user@ozon.ru", UserId: 1}
				acc, _ := entity.NewAccount(1, "user@ozon.ru")

				mockRepo.EXPECT().AddAccounts(ctx, mocks.AccountValueEq([]entity.Account{*acc})).Return(nil)
				producer.EXPECT().Send(ctx, gomock.Any()).Return(nil) // TODO: mb create special matcher
				metrics.EXPECT().IncCreatedCounter().Times(1)

				resp, err := service.CreateAccount(ctx, req)

				checkAccountInResponse(resp.GetAccount(), *acc)
				Expect(err).Should(BeNil())
			})
		})
	})

	Describe("DescribeAccount", func() {
		Context("when requesting an existing account", func() {
			It("should successfully return", func() {
				acc, _ := entity.NewAccount(1, "user@ozon.ru")
				req := &pb.DescribeAccountRequest{Id: acc.ID}

				mockRepo.EXPECT().DescribeAccount(ctx, acc.ID).Return(acc, nil)

				resp, err := service.DescribeAccount(ctx, req)
				checkAccountInResponse(resp.GetAccount(), *acc)
				Expect(err).Should(BeNil())
			})
		})
	})

	Describe("UpdateAccount", func() {
		Context("when updating an existing account", func() {
			It("should successfully update", func() {
				acc, _ := entity.NewAccount(1, "user@ozon.ru")
				req := &pb.UpdateAccountRequest{
					Account: AccountMarshal(*acc),
				}

				mockRepo.EXPECT().UpdateAccount(ctx, *acc).Return(nil)
				producer.EXPECT().Send(ctx, gomock.Any()).Return(nil)
				metrics.EXPECT().IncUpdatedCounter().Times(1)

				resp, err := service.UpdateAccount(ctx, req)
				checkAccountInResponse(resp.GetAccount(), *acc)
				Expect(err).Should(BeNil())
			})
		})

		Context("when updating a non-existent account", func() {
			It("it should return the error record not found", func() {
				acc, _ := entity.NewAccount(1, "user@ozon.ru")
				req := &pb.UpdateAccountRequest{
					Account: AccountMarshal(*acc),
				}

				mockRepo.EXPECT().UpdateAccount(ctx, *acc).Return(repo.ErrRecordNotFound)

				_, err := service.UpdateAccount(ctx, req)
				Expect(err).Should(BeEquivalentTo(status.Errorf(codes.NotFound, "record not found")))
			})
		})
	})

	Describe("MultiCreateAccount", func() {
		Context("when adding an valid accounts", func() {
			It("should successfully store", func() {
				accounts := entity.CreateTestAccounts(25)
				req := createMultiCreateAccountRequest(accounts)

				mockRepo.EXPECT().AddAccounts(gomock.Any(), mocks.AccountValueEq(accounts)).Return(nil)
				producer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				metrics.EXPECT().IncreaseCreatedCounter(25).Times(1)

				_, err := service.MultiCreateAccount(ctx, req)
				Expect(err).Should(BeNil())
			})
		})

		Context("when adding an valid accounts more then chunks size", func() {
			It("should successfully return", func() {
				accounts := entity.CreateTestAccounts(55)
				req := createMultiCreateAccountRequest(accounts)

				gomock.InOrder(
					mockRepo.EXPECT().AddAccounts(gomock.Any(), mocks.AccountValueEq(accounts[:32])).Return(nil),
					mockRepo.EXPECT().AddAccounts(gomock.Any(), mocks.AccountValueEq(accounts[32:])).Return(nil),
				)
				producer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).Times(2)
				metrics.EXPECT().IncreaseCreatedCounter(batchSize).Times(1)
				metrics.EXPECT().IncreaseCreatedCounter(55 - batchSize).Times(1)

				_, err := service.MultiCreateAccount(ctx, req)
				Expect(err).Should(BeNil())
			})
		})
	})

	// TODO: add more tests
})

func checkAccountInResponse(resp *pb.Account, acc entity.Account) {
	// Expect(resp.Id).Should(BeIdenticalTo(acc.ID))
	Expect(resp.Value).Should(BeIdenticalTo(acc.Value))
	Expect(resp.UserId).Should(BeIdenticalTo(acc.UserID))
}

func createMultiCreateAccountRequest(accounts []entity.Account) *pb.MultiCreateAccountRequest {
	pbAccounts := make([]*pb.CreateAccountRequest, 0, len(accounts))
	for _, acc := range accounts {
		pbAccounts = append(pbAccounts, &pb.CreateAccountRequest{
			UserId: acc.UserID,
			Value:  acc.Value,
		})
	}

	return &pb.MultiCreateAccountRequest{Accounts: pbAccounts}
}
