package saver_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/mocks"
	"github.com/ozonva/ova-account-api/internal/saver"
)

func TestSaver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Saver Suite")
}

var _ = Describe("Saver", func() {
	var (
		ctrl    *gomock.Controller
		flusher *mocks.MockFlusher
		s       saver.Saver
		ctx     context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		flusher = mocks.NewMockFlusher(ctrl)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Save accounts with buffering", func() {

		Context("when only one account", func() {
			It("should successfully save", func() {
				accounts := generateAccounts(1)

				flusher.EXPECT().Flush(ctx, accounts).Times(1).Return(nil)
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()

				Expect(s.Save(accounts[0])).ShouldNot(HaveOccurred())
				time.Sleep(3 * time.Millisecond)

				s.Close()
			})
		})

		Context("when multiple accounts", func() {
			It("should successfully save", func() {
				accounts := generateAccounts(2)
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()

				flusher.EXPECT().Flush(ctx, accounts).Times(1).Return(nil)

				for _, account := range accounts {
					Expect(s.Save(account)).ShouldNot(HaveOccurred())
				}

				time.Sleep(3 * time.Millisecond)
				s.Close()
			})
		})

		Context("when multiple accounts with delay", func() {
			It("should successfully save", func() {
				accounts := generateAccounts(4)
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()

				gomock.InOrder(
					flusher.EXPECT().Flush(ctx, accounts[0:2]).Times(1).Return(nil),
					flusher.EXPECT().Flush(ctx, accounts[2:4]).Times(1).Return(nil),
				)

				Expect(s.Save(accounts[0])).ShouldNot(HaveOccurred())
				Expect(s.Save(accounts[1])).ShouldNot(HaveOccurred())
				time.Sleep(3 * time.Millisecond)
				Expect(s.Save(accounts[2])).ShouldNot(HaveOccurred())
				Expect(s.Save(accounts[3])).ShouldNot(HaveOccurred())

				time.Sleep(3 * time.Millisecond)
				s.Close()
			})
		})

		Context("when closing is called", func() {
			It("should successfully save", func() {
				accounts := generateAccounts(2)
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()

				flusher.EXPECT().Flush(ctx, accounts[0:2]).Times(1).Return(nil)

				for _, account := range accounts {
					Expect(s.Save(account)).ShouldNot(HaveOccurred())
				}
				s.Close()

				time.Sleep(3 * time.Millisecond)
			})
		})

		Context("when couldn't save the first time", func() {
			It("should successfully re-saving", func() {
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()
				accounts := generateAccounts(2)

				gomock.InOrder(
					flusher.EXPECT().Flush(ctx, accounts).Times(1).Return(accounts), // Couldn't save
					flusher.EXPECT().Flush(ctx, accounts).Times(1).Return(nil),
				)

				for _, account := range accounts {
					Expect(s.Save(account)).ShouldNot(HaveOccurred())
				}

				time.Sleep(3 * time.Millisecond)
				s.Close()
			})
		})

		Context("when the buffer size is exceeded", func() {
			It("should successfully save", func() {
				s = saver.NewSaver(2, flusher, time.Millisecond)
				s.Init()
				accounts := generateAccounts(4)

				gomock.InOrder(
					flusher.EXPECT().Flush(ctx, accounts[0:2]).Times(1).Return(nil), // Flush when overflowing
					flusher.EXPECT().Flush(ctx, accounts[2:4]).Times(1).Return(nil),
				)

				for _, account := range accounts {
					Expect(s.Save(account)).ShouldNot(HaveOccurred())
				}

				time.Sleep(3 * time.Millisecond)
				s.Close()
			})
		})

		Context("when closing and re-init of saver", func() {
			It("should successfully work", func() {
				accounts := generateAccounts(3)

				gomock.InOrder(
					flusher.EXPECT().Flush(ctx, accounts[0:1]).Times(1).Return(nil),
					flusher.EXPECT().Flush(ctx, accounts[1:3]).Times(1).Return(nil),
				)

				s = saver.NewSaver(2, flusher, time.Millisecond)
				s.Init()
				Expect(s.Save(accounts[0])).ShouldNot(HaveOccurred())
				time.Sleep(3 * time.Millisecond)
				s.Close()

				s.Init()
				Expect(s.Save(accounts[1])).ShouldNot(HaveOccurred())
				Expect(s.Save(accounts[2])).ShouldNot(HaveOccurred())
				time.Sleep(3 * time.Millisecond)
				s.Close()
			})
		})

		Context("when the buffer is full and cannot be flushed", func() {
			It("should return error", func() {
				accounts := generateAccounts(4)

				gomock.InOrder(
					flusher.EXPECT().Flush(ctx, accounts[0:2]).Times(1).Return(accounts[0:2]),
					flusher.EXPECT().Flush(ctx, accounts[0:2]).Times(1).Return(accounts[0:2]),
				)

				s = saver.NewSaver(2, flusher, time.Millisecond)
				s.Init()
				defer s.Close()

				Expect(s.Save(accounts[0])).ShouldNot(HaveOccurred())
				Expect(s.Save(accounts[1])).ShouldNot(HaveOccurred())
				Expect(s.Save(accounts[2])).Should(Equal(saver.ErrFullBufferFlush))
				Expect(s.Save(accounts[3])).Should(Equal(saver.ErrFullBufferFlush))
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
