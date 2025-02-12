package service

import (
	"context"

	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
)

type UserStorage interface {
	GetByUsername(ctx context.Context, username string) (userDomain.User, error)
	Create(ctx context.Context, user userDomain.User) (string, error)
	GetById(ctx context.Context, id string) (userDomain.User, error)
}

type Dependency struct {
	UserStorage UserStorage
}
