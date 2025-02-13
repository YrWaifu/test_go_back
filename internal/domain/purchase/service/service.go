package service

import (
	"context"
	"fmt"
	userStorage "github.com/YrWaifu/test_go_back/internal/domain/user/storage"
)

type Service struct {
	d Dependency
}

func New(d Dependency) *Service {
	return &Service{d: d}
}

type BuyRequest struct {
	UserID    string
	MerchName string
}

type BuyResponse struct {
}

func (s *Service) BuyMerch(ctx context.Context, req BuyRequest) (BuyResponse, error) {
	merch, err := s.d.MerchStorage.GetByName(ctx, req.MerchName)
	if err != nil {
		return BuyResponse{}, fmt.Errorf("get merch by name: %w", err)
	}

	// transaction
	err = s.d.PurchaseStorage.BeginPurchase(ctx, func(ctx context.Context) error {
		user, err := s.d.UserStorage.GetById(ctx, req.UserID, userStorage.GetOptions{ForUpdate: true})
		if err != nil {
			return fmt.Errorf("get user by id: %w", err)
		}

		err = s.d.UserStorage.IncrementBalance(ctx, user.Username, merch.Price*-1)
		if err != nil {
			return fmt.Errorf("increment balance: %w", err)
		}

		err = s.d.PurchaseStorage.CreatePurchase(ctx, user.ID, merch.ID)
		if err != nil {
			return fmt.Errorf("create purchase: %w", err)
		}

		return nil
	})
	if err != nil {
		return BuyResponse{}, err
	}

	return BuyResponse{}, nil
}
