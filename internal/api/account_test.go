package api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/mocks"
	pb "github.com/ozonva/ova-account-api/pkg/ova-account-api"
	"github.com/rs/zerolog"
)

func TestAccountService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Account Service Suite")
}

var _ = Describe("Account Service", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		service  *AccountService
		ctx      context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		service = NewAccountService(zerolog.Logger{}, mockRepo)
		ctx = context.TODO()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("CreateAccount", func() {
		Context("when adding an valid account", func() {
			It("should successfully store", func() {
				req := &pb.CreateAccountRequest{Value: "user@ozon.ru"}
				acc := entity.Account{ID: 1, UserID: 1, Value: "user@ozon.ru"}

				mockRepo.EXPECT().AddAccounts(ctx, []entity.Account{acc}).Return(nil)

				resp, err := service.CreateAccount(ctx, req)

				checkAccountInResponse(resp.GetAccount(), acc)
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

	// TODO: add more tests
})

func checkAccountInResponse(resp *pb.Account, acc entity.Account) {
	Expect(resp.Id).Should(BeIdenticalTo(acc.ID))
	Expect(resp.Value).Should(BeIdenticalTo(acc.Value))
	Expect(resp.UserId).Should(BeIdenticalTo(acc.UserID))
}
