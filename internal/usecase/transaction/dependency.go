package transaction

//go:generate mockgen -source dependency.go -destination mock/dependency.go

import (
	"context"
	transactionService "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
)

type TransactionService interface {
	TransferCoins(ctx context.Context, req transactionService.TransferRequest) (transactionService.TransferResponse, error)
}

type Dependency struct {
	TransactionService TransactionService
}
