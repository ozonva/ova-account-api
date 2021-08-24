package flusher_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/flusher"
	"github.com/ozonva/ova-account-api/internal/mocks"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		accounts []entity.Account
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Flush the list of accounts to the storage", func() {
		BeforeEach(func() {
			accounts = generateAccounts(10)
		})

		Context("when the batch size is 3", func() {
			It("should return nil", func() {
				flusher := flusher.NewFlusher(3, mockRepo)
				gomock.InOrder(
					mockRepo.EXPECT().AddAccounts(accounts[0:3]).Return(nil),
					mockRepo.EXPECT().AddAccounts(accounts[3:6]).Return(nil),
					mockRepo.EXPECT().AddAccounts(accounts[6:9]).Return(nil),
					mockRepo.EXPECT().AddAccounts(accounts[9:10]).Return(nil),
				)

				Expect(flusher.Flush(accounts)).Should(BeNil())
			})
		})

		Context("when the repo returns an error", func() {
			It("should return a part of the list", func() {
				flusher := flusher.NewFlusher(5, mockRepo)
				gomock.InOrder(
					mockRepo.EXPECT().AddAccounts(accounts[0:5]).Return(nil),
					mockRepo.EXPECT().AddAccounts(accounts[5:10]).Return(errors.New("can't store")),
				)

				Expect(flusher.Flush(accounts)).Should(Equal(accounts[5:10]))
			})
		})

		Context("when the batch size is invalid", func() {
			It("should return the entire list", func() {
				flusher := flusher.NewFlusher(-1, mockRepo)
				Expect(flusher.Flush(accounts)).Should(Equal(accounts))
			})
		})
	})
})

func generateAccounts(count int) []entity.Account {
	out := make([]entity.Account, 0, count)

	for i := 0; i < count; i++ {
		account, _ := entity.NewAccount(1, fmt.Sprintf("user%d@ozon.ru", i+1))
		out = append(out, *account)
	}

	return out
}
