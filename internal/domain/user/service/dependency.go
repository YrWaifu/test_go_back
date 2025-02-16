package service

import (
	"context"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"github.com/YrWaifu/test_go_back/internal/domain/user/storage"
)

type UserStorage interface {
	GetById(ctx context.Context, id string, opts storage.GetOptions) (userDomain.User, error)
}

type Dependency struct {
	UserStorage UserStorage
}
