package purchase

//go:generate mockgen -source dependency.go -destination mock/dependency.go

import (
	"context"
	purchaseService "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
)

type PurchaseService interface {
	BuyMerch(ctx context.Context, req purchaseService.BuyRequest) (purchaseService.BuyResponse, error)
}

type Dependency struct {
	PurchaseService PurchaseService
}
