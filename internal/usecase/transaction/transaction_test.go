package transaction

import (
	"context"
	transactionservice "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
	mock_transaction "github.com/YrWaifu/test_go_back/internal/usecase/transaction/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUsecase_Transfer(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := TransferRequest{
			SenderID:     "1",
			ReceiverName: "receiver",
			Amount:       100,
		}

		transactionService := mock_transaction.NewMockTransactionService(ctrl)
		transactionService.EXPECT().
			TransferCoins(gomock.Any(), transactionservice.TransferRequest{
				SenderID:     req.SenderID,
				ReceiverName: req.ReceiverName,
				Amount:       req.Amount,
			}).
			Return(transactionservice.TransferResponse{}, nil)

		usecase := New(Dependency{TransactionService: transactionService})

		resp, err := usecase.Transfer(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, TransferResponse{}, resp)
	})
}
