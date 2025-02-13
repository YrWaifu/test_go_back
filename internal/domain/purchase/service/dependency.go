package service

import (
	"context"
	merchDomain "github.com/YrWaifu/test_go_back/internal/domain/merch"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	userStorage "github.com/YrWaifu/test_go_back/internal/domain/user/storage"
)

type PurchaseStorage interface {
	BeginPurchase(ctx context.Context, fn func(context.Context) error) error
	CreatePurchase(ctx context.Context, userID string, merchID string) error
}

type UserStorage interface {
	GetById(ctx context.Context, id string, opts userStorage.GetOptions) (userDomain.User, error)
	IncrementBalance(ctx context.Context, username string, inc int) error
}

type MerchStorage interface {
	GetByName(ctx context.Context, name string) (merchDomain.Merch, error)
}

type Dependency struct {
	PurchaseStorage PurchaseStorage
	UserStorage     UserStorage
	MerchStorage    MerchStorage
}
