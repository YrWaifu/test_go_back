package info

import (
	"context"
	"fmt"
	purchaseDomain "github.com/YrWaifu/test_go_back/internal/domain/purchase"
	purchaseService "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
	transactionDomain "github.com/YrWaifu/test_go_back/internal/domain/transaction"
	transactionService "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	userService "github.com/YrWaifu/test_go_back/internal/domain/user/service"
)

type Usecase struct {
	d Dependency
}

func New(d Dependency) *Usecase {
	return &Usecase{d: d}
}

type InfoRequest struct {
	UserId string
}

type InfoResponse struct {
	User      userDomain.User
	Sent      []transactionDomain.Transaction
	Received  []transactionDomain.Transaction
	Purchases []purchaseDomain.Purchase
}

func (u *Usecase) Info(ctx context.Context, req InfoRequest) (InfoResponse, error) {
	userResponse, err := u.d.UserService.GetByID(ctx, userService.GetByIDRequest{ID: req.UserId})
	if err != nil {
		return InfoResponse{}, fmt.Errorf("user service get by id: %w", err)
	}

	transactionResponse, err := u.d.TransactionService.ListByUserID(ctx, transactionService.ListByUserIDRequest{UserID: req.UserId})
	if err != nil {
		return InfoResponse{}, fmt.Errorf("transaction service list by id: %w", err)
	}

	purchaseResponse, err := u.d.PurchaseService.ListByUserID(ctx, purchaseService.ListByUserIDRequest{UserID: req.UserId})
	if err != nil {
		return InfoResponse{}, fmt.Errorf("purchase service list by id: %w", err)
	}

	return InfoResponse{
		User:      userResponse.User,
		Sent:      transactionResponse.Sent,
		Received:  transactionResponse.Received,
		Purchases: purchaseResponse.Purchases,
	}, nil
}
