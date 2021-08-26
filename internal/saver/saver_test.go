package saver_test

import (
	"fmt"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-account-api/internal/entity"
	"github.com/ozonva/ova-account-api/internal/mocks"
	"github.com/ozonva/ova-account-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		ctrl    *gomock.Controller
		flusher *mocks.MockFlusher
		s       saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		flusher = mocks.NewMockFlusher(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Save accounts with buffering", func() {

		Context("when only one account", func() {
			It("should successfully save", func() {
				accounts := generateAccounts(1)

				flusher.EXPECT().Flush(accounts).Times(1).Return(nil)
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()

				s.Save(accounts[0])
				time.Sleep(3 * time.Millisecond)

				s.Close()
			})
		})

		Context("when multiple accounts", func() {
			It("should successfully save", func() {
				accounts := generateAccounts(2)
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()

				flusher.EXPECT().Flush(accounts).Times(1).Return(nil)

				for _, account := range accounts {
					s.Save(account)
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
					flusher.EXPECT().Flush(accounts[0:2]).Times(1).Return(nil),
					flusher.EXPECT().Flush(accounts[2:4]).Times(1).Return(nil),
				)

				s.Save(accounts[0])
				s.Save(accounts[1])
				time.Sleep(2 * time.Millisecond)
				s.Save(accounts[2])
				s.Save(accounts[3])

				time.Sleep(3 * time.Millisecond)
				s.Close()
			})
		})

		Context("when closing is called", func() {
			It("should successfully save", func() {
				accounts := generateAccounts(2)
				s = saver.NewSaver(10, flusher, time.Millisecond)
				s.Init()

				flusher.EXPECT().Flush(accounts[0:2]).Times(1).Return(nil)

				for _, account := range accounts {
					s.Save(account)
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
					flusher.EXPECT().Flush(accounts).Times(1).Return(accounts), // Couldn't save
					flusher.EXPECT().Flush(accounts).Times(1).Return(nil),
				)

				for _, account := range accounts {
					s.Save(account)
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
					flusher.EXPECT().Flush(accounts[0:2]).Times(1).Return(nil), // Flush when overflowing
					flusher.EXPECT().Flush(accounts[2:4]).Times(1).Return(nil),
				)

				for _, account := range accounts {
					s.Save(account)
				}

				time.Sleep(3 * time.Millisecond)
				s.Close()
			})
		})

		Context("when closing and re-init of saver", func() {
			It("should successfully work", func() {
				accounts := generateAccounts(3)

				gomock.InOrder(
					flusher.EXPECT().Flush(accounts[0:1]).Times(1).Return(nil),
					flusher.EXPECT().Flush(accounts[1:3]).Times(1).Return(nil),
				)

				s = saver.NewSaver(2, flusher, time.Millisecond)
				s.Init()
				s.Save(accounts[0])
				time.Sleep(3 * time.Millisecond)
				s.Close()

				s.Init()
				s.Save(accounts[1])
				s.Save(accounts[2])
				time.Sleep(3 * time.Millisecond)
				s.Close()
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
