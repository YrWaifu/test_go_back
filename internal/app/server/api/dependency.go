package api

//go:generate mockgen -source dependency.go -destination mock/dependency.go

import (
	"context"
	authUsecase "github.com/YrWaifu/test_go_back/internal/usecase/auth"
	infoUsecase "github.com/YrWaifu/test_go_back/internal/usecase/info"
	purchaseUsecase "github.com/YrWaifu/test_go_back/internal/usecase/purchase"
)

import (
	transfactionUsecase "github.com/YrWaifu/test_go_back/internal/usecase/transaction"
)

type AuthUsecase interface {
	SignIn(ctx context.Context, req authUsecase.SignInRequest) (authUsecase.SignInResponse, error)
}

type AuthDependency struct {
	AuthUsecase AuthUsecase
}

type PurchaseUsecase interface {
	BuyMerch(ctx context.Context, req purchaseUsecase.BuyMerchRequest) (purchaseUsecase.BuyMerchResponse, error)
}

type PurchaseDependency struct {
	PurchaseUsecase PurchaseUsecase
}

type TransactionUsecase interface {
	Transfer(ctx context.Context, req transfactionUsecase.TransferRequest) (transfactionUsecase.TransferResponse, error)
}

type TransactionDependency struct {
	TransactionUsecase TransactionUsecase
}

type InfoUsecase interface {
	Info(ctx context.Context, req infoUsecase.InfoRequest) (infoUsecase.InfoResponse, error)
}

type InfoDependency struct {
	InfoUsecase InfoUsecase
}
