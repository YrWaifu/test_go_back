package service

import (
	"context"
	"fmt"
	transactionDomain "github.com/YrWaifu/test_go_back/internal/domain/transaction"
	"github.com/YrWaifu/test_go_back/internal/domain/user/storage"
)

type Service struct {
	d Dependency
}

func New(d Dependency) *Service {
	return &Service{d: d}
}

type TransferRequest struct {
	SenderID     string
	ReceiverName string
	Amount       int
}

type TransferResponse struct {
}

func (s *Service) TransferCoins(ctx context.Context, req TransferRequest) (TransferResponse, error) {
	err := s.d.TransactionStorage.BeginTransaction(ctx, func(ctx context.Context) error {
		sender, err := s.d.UserStorage.GetById(ctx, req.SenderID, storage.GetOptions{ForUpdate: true})
		if err != nil {
			return fmt.Errorf("get user sender failed: %w", err)
		}

		if sender.Balance-req.Amount < 0 {
			return fmt.Errorf("user sender has not enough balance")
		}

		receiver, err := s.d.UserStorage.GetByUsername(ctx, req.ReceiverName, storage.GetOptions{ForUpdate: true})
		if err != nil {
			return fmt.Errorf("get user reciever failed: %w", err)
		}

		err = s.d.UserStorage.IncrementBalance(ctx, req.ReceiverName, req.Amount)
		if err != nil {
			return fmt.Errorf("increment balance failed: %w", err)
		}

		err = s.d.UserStorage.IncrementBalance(ctx, sender.Username, (-1)*req.Amount)
		if err != nil {
			return fmt.Errorf("increment balance failed: %w", err)
		}

		err = s.d.TransactionStorage.CreateTransaction(ctx, sender.ID, receiver.ID, req.Amount)
		if err != nil {
			return fmt.Errorf("create transaction failed: %w", err)
		}

		return nil
	})
	if err != nil {
		return TransferResponse{}, err
	}

	return TransferResponse{}, nil
}

type ListByUserIDRequest struct {
	UserID string
}

type ListByUserIDResponse struct {
	Sent     []transactionDomain.Transaction
	Received []transactionDomain.Transaction
}

func (s *Service) ListByUserID(ctx context.Context, r ListByUserIDRequest) (ListByUserIDResponse, error) {
	sent, received, err := s.d.TransactionStorage.ListByUserID(ctx, r.UserID)
	if err != nil {
		return ListByUserIDResponse{}, fmt.Errorf("list transactions by user id: %w", err)
	}

	return ListByUserIDResponse{
		Sent:     sent,
		Received: received,
	}, nil
}
