package purchase

import (
	"context"
	"fmt"
	purchaseService "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
)

type UseCase struct {
	d Dependency
}

func New(d Dependency) *UseCase {
	return &UseCase{d: d}
}

type BuyMerchRequest struct {
	UserID    string
	MerchName string
}

type BuyMerchResponse struct{}

func (u *UseCase) BuyMerch(ctx context.Context, req BuyMerchRequest) (BuyMerchResponse, error) {
	_, err := u.d.PurchaseService.BuyMerch(ctx, purchaseService.BuyRequest{
		UserID:    req.UserID,
		MerchName: req.MerchName,
	})
	if err != nil {
		return BuyMerchResponse{}, fmt.Errorf("purchase service buy merch: %w", err)
	}

	return BuyMerchResponse{}, nil
}
