package transaction

import (
	"context"
	"fmt"
	transactionService "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
)

type Usecase struct {
	d Dependency
}

func New(d Dependency) *Usecase {
	return &Usecase{d: d}
}

type TransferRequest struct {
	SenderID     string
	ReceiverName string
	Amount       int
}

type TransferResponse struct {
}

func (u *Usecase) Transfer(ctx context.Context, req TransferRequest) (TransferResponse, error) {
	_, err := u.d.TransactionService.TransferCoins(ctx, transactionService.TransferRequest{
		SenderID:     req.SenderID,
		ReceiverName: req.ReceiverName,
		Amount:       req.Amount,
	})
	if err != nil {
		return TransferResponse{}, fmt.Errorf("transfer coins: %w", err)
	}

	return TransferResponse{}, nil
}
