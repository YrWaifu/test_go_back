package purchase

import (
	"context"
	purchaseservice "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
	mock_purchase "github.com/YrWaifu/test_go_back/internal/usecase/purchase/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUseCase_BuyMerch(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := BuyMerchRequest{
			UserID:    "1",
			MerchName: "test_merch",
		}

		purchaseService := mock_purchase.NewMockPurchaseService(ctrl)
		purchaseService.EXPECT().
			BuyMerch(gomock.Any(), purchaseservice.BuyRequest{UserID: req.UserID, MerchName: req.MerchName}).
			Return(purchaseservice.BuyResponse{}, nil)

		usecase := New(Dependency{PurchaseService: purchaseService})

		resp, err := usecase.BuyMerch(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, BuyMerchResponse{}, resp)
	})
}
