package service

import (
	"context"
	"fmt"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	userStorage "github.com/YrWaifu/test_go_back/internal/domain/user/storage"
)

type Service struct {
	d Dependency
}

func New(d Dependency) *Service {
	return &Service{d: d}
}

type GetByIDRequest struct {
	ID string
}

type GetByIDResponse struct {
	User userDomain.User
}

func (s *Service) GetByID(ctx context.Context, req GetByIDRequest) (GetByIDResponse, error) {
	user, err := s.d.UserStorage.GetById(ctx, req.ID, userStorage.GetOptions{})
	if err != nil {
		return GetByIDResponse{}, fmt.Errorf("get user by id: %w", err)
	}

	return GetByIDResponse{
		User: user,
	}, nil
}
