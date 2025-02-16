package info

import (
	"context"
	purchaseDomain "github.com/YrWaifu/test_go_back/internal/domain/purchase"
	purchaseservice "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
	transactionDomain "github.com/YrWaifu/test_go_back/internal/domain/transaction"
	transactionservice "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	userservice "github.com/YrWaifu/test_go_back/internal/domain/user/service"
	mock_info "github.com/YrWaifu/test_go_back/internal/usecase/info/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUsecase_Info(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := InfoRequest{
			UserId: "1",
		}

		user := userDomain.User{ID: "1", Username: "test_user"}
		sentTransactions := []transactionDomain.Transaction{
			{SenderId: "1", ReceiverId: "2", Amount: 100},
		}
		receivedTransactions := []transactionDomain.Transaction{
			{SenderId: "3", ReceiverId: "1", Amount: 50},
		}
		purchases := []purchaseDomain.Purchase{
			{UserID: "1", MerchID: "1"},
		}

		userService := mock_info.NewMockUserService(ctrl)
		transactionService := mock_info.NewMockTransactionService(ctrl)
		purchaseService := mock_info.NewMockPurchaseService(ctrl)

		userService.EXPECT().GetByID(gomock.Any(), userservice.GetByIDRequest{ID: req.UserId}).
			Return(userservice.GetByIDResponse{User: user}, nil)
		transactionService.EXPECT().ListByUserID(gomock.Any(), transactionservice.ListByUserIDRequest{UserID: req.UserId}).
			Return(transactionservice.ListByUserIDResponse{Sent: sentTransactions, Received: receivedTransactions}, nil)
		purchaseService.EXPECT().ListByUserID(gomock.Any(), purchaseservice.ListByUserIDRequest{UserID: req.UserId}).
			Return(purchaseservice.ListByUserIDResponse{Purchases: purchases}, nil)

		usecase := New(Dependency{
			UserService:        userService,
			TransactionService: transactionService,
			PurchaseService:    purchaseService,
		})

		resp, err := usecase.Info(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, user, resp.User)
		assert.Equal(t, sentTransactions, resp.Sent)
		assert.Equal(t, receivedTransactions, resp.Received)
		assert.Equal(t, purchases, resp.Purchases)
	})
}
