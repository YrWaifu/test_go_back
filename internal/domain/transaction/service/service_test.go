package service

import (
	"context"
	transactionDomain "github.com/YrWaifu/test_go_back/internal/domain/transaction"
	mock_service "github.com/YrWaifu/test_go_back/internal/domain/transaction/service/mock"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"github.com/YrWaifu/test_go_back/internal/domain/user/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_TransferCoins(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := TransferRequest{
			SenderID:     "1",
			ReceiverName: "test_receiver",
			Amount:       100,
		}

		sender := userDomain.User{ID: "1", Username: "test_sender", Balance: 200}
		receiver := userDomain.User{ID: "2", Username: "test_receiver", Balance: 300}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		transactionStorage := mock_service.NewMockTransactionStorage(ctrl)

		userStorage.EXPECT().GetById(gomock.Any(), req.SenderID, storage.GetOptions{ForUpdate: true}).Return(sender, nil)
		userStorage.EXPECT().GetByUsername(gomock.Any(), req.ReceiverName, storage.GetOptions{ForUpdate: true}).Return(receiver, nil)
		userStorage.EXPECT().IncrementBalance(gomock.Any(), req.ReceiverName, req.Amount).Return(nil)
		userStorage.EXPECT().IncrementBalance(gomock.Any(), sender.Username, -req.Amount).Return(nil)
		transactionStorage.EXPECT().CreateTransaction(gomock.Any(), sender.ID, receiver.ID, req.Amount).Return(nil)

		transactionStorage.EXPECT().BeginTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(ctx context.Context) error) error {
				return fn(ctx)
			},
		)
		service := New(Dependency{
			TransactionStorage: transactionStorage,
			UserStorage:        userStorage,
		})

		_, err := service.TransferCoins(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("InsufficientBalance", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := TransferRequest{
			SenderID:     "1",
			ReceiverName: "test_receiver",
			Amount:       300,
		}

		sender := userDomain.User{ID: "1", Username: "test_sender", Balance: 200}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		transactionStorage := mock_service.NewMockTransactionStorage(ctrl)

		userStorage.EXPECT().GetById(gomock.Any(), req.SenderID, storage.GetOptions{ForUpdate: true}).Return(sender, nil)

		transactionStorage.EXPECT().BeginTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(ctx context.Context) error) error {
				return fn(ctx)
			},
		)

		service := New(Dependency{
			TransactionStorage: transactionStorage,
			UserStorage:        userStorage,
		})

		_, err := service.TransferCoins(context.Background(), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not enough balance")
	})
}

func TestService_Transfer(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := ListByUserIDRequest{UserID: "1"}

		sentTransactions := []transactionDomain.Transaction{
			{SenderId: "1", ReceiverId: "2", Amount: 100},
		}
		receivedTransactions := []transactionDomain.Transaction{
			{SenderId: "2", ReceiverId: "3", Amount: 50},
		}

		transactionStorage := mock_service.NewMockTransactionStorage(ctrl)
		transactionStorage.EXPECT().
			ListByUserID(gomock.Any(), req.UserID).
			Return(sentTransactions, receivedTransactions, nil)
		service := New(Dependency{TransactionStorage: transactionStorage})

		resp, err := service.ListByUserID(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, sentTransactions, resp.Sent)
		assert.Equal(t, receivedTransactions, resp.Received)
	})
}
