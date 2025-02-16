package service

import (
	"context"
	merchDomain "github.com/YrWaifu/test_go_back/internal/domain/merch"
	purchaseDomain "github.com/YrWaifu/test_go_back/internal/domain/purchase"
	mock_service "github.com/YrWaifu/test_go_back/internal/domain/purchase/service/mock"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"github.com/YrWaifu/test_go_back/internal/domain/user/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_BuyMerch(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := BuyRequest{
			UserID:    "1",
			MerchName: "test_merch",
		}

		user := userDomain.User{ID: "1", Username: "test_user", Balance: 200}
		merch := merchDomain.Merch{ID: "1", Name: "test_merch", Price: 100}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		purchaseStorage := mock_service.NewMockPurchaseStorage(ctrl)
		merchStorage := mock_service.NewMockMerchStorage(ctrl)

		merchStorage.EXPECT().GetByName(gomock.Any(), req.MerchName).Return(merch, nil)
		userStorage.EXPECT().GetById(gomock.Any(), req.UserID, storage.GetOptions{ForUpdate: true}).Return(user, nil)
		userStorage.EXPECT().IncrementBalance(gomock.Any(), user.Username, -merch.Price).Return(nil)
		purchaseStorage.EXPECT().CreatePurchase(gomock.Any(), user.ID, merch.ID).Return(nil)

		purchaseStorage.EXPECT().BeginPurchase(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(ctx context.Context) error) error {
				return fn(ctx)
			},
		)

		service := New(Dependency{
			MerchStorage:    merchStorage,
			PurchaseStorage: purchaseStorage,
			UserStorage:     userStorage,
		})

		_, err := service.BuyMerch(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("InsufficientBalance", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := BuyRequest{
			UserID:    "1",
			MerchName: "test_merch",
		}

		user := userDomain.User{ID: "1", Username: "test_user", Balance: 50}
		merch := merchDomain.Merch{ID: "1", Name: "test_merch", Price: 100}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		purchaseStorage := mock_service.NewMockPurchaseStorage(ctrl)
		merchStorage := mock_service.NewMockMerchStorage(ctrl)

		merchStorage.EXPECT().GetByName(gomock.Any(), req.MerchName).Return(merch, nil)
		userStorage.EXPECT().GetById(gomock.Any(), req.UserID, storage.GetOptions{ForUpdate: true}).Return(user, nil)

		purchaseStorage.EXPECT().BeginPurchase(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(ctx context.Context) error) error {
				return fn(ctx)
			},
		)

		service := New(Dependency{
			MerchStorage:    merchStorage,
			PurchaseStorage: purchaseStorage,
			UserStorage:     userStorage,
		})

		_, err := service.BuyMerch(context.Background(), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "out of money")
	})
}

func TestService_ListByUserID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := ListByUserIDRequest{
			UserID: "1",
		}

		purchases := []purchaseDomain.Purchase{
			{UserID: "1", MerchID: "1"},
		}

		purchaseStorage := mock_service.NewMockPurchaseStorage(ctrl)
		purchaseStorage.EXPECT().ListByUserID(gomock.Any(), req.UserID).Return(purchases, nil)

		service := New(Dependency{PurchaseStorage: purchaseStorage})

		resp, err := service.ListByUserID(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, purchases, resp.Purchases)
	})
}
