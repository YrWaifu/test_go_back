package info

//go:generate mockgen -source dependency.go -destination mock/dependency.go

import (
	"context"
	purchaseService "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
	transactionService "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
	userService "github.com/YrWaifu/test_go_back/internal/domain/user/service"
)

type PurchaseService interface {
	ListByUserID(ctx context.Context, r purchaseService.ListByUserIDRequest) (purchaseService.ListByUserIDResponse, error)
}

type TransactionService interface {
	ListByUserID(ctx context.Context, r transactionService.ListByUserIDRequest) (transactionService.ListByUserIDResponse, error)
}

type UserService interface {
	GetByID(ctx context.Context, req userService.GetByIDRequest) (userService.GetByIDResponse, error)
}

type Dependency struct {
	PurchaseService    PurchaseService
	TransactionService TransactionService
	UserService        UserService
}
