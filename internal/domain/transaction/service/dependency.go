package service

import (
	"context"
	transactionDomain "github.com/YrWaifu/test_go_back/internal/domain/transaction"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"github.com/YrWaifu/test_go_back/internal/domain/user/storage"
)

type TransactionStorage interface {
	CreateTransaction(ctx context.Context, senderID string, receiverID string, amount int) error
	BeginTransaction(ctx context.Context, fn func(context.Context) error) error
	ListByUserID(ctx context.Context, userID string) ([]transactionDomain.Transaction, []transactionDomain.Transaction, error)
}

type UserStorage interface {
	GetByUsername(ctx context.Context, username string, opts storage.GetOptions) (userDomain.User, error)
	GetById(ctx context.Context, id string, opts storage.GetOptions) (userDomain.User, error)
	IncrementBalance(ctx context.Context, username string, inc int) error
}

type Dependency struct {
	TransactionStorage TransactionStorage
	UserStorage        UserStorage
}
