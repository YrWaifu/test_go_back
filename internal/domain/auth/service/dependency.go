package service

//go:generate mockgen -source dependency.go -destination mock/dependency.go

import (
	"context"

	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	userStorage "github.com/YrWaifu/test_go_back/internal/domain/user/storage"
)

type UserStorage interface {
	GetByUsername(ctx context.Context, username string, opts userStorage.GetOptions) (userDomain.User, error)
	Create(ctx context.Context, user userDomain.User) (string, error)
	GetById(ctx context.Context, id string, opts userStorage.GetOptions) (userDomain.User, error)
}

type Dependency struct {
	UserStorage UserStorage
}
