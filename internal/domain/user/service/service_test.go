package service

import (
	"context"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	mock_service "github.com/YrWaifu/test_go_back/internal/domain/user/service/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_GetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "1"
		expectedUser := userDomain.User{
			ID:       id,
			Username: "test",
			Balance:  1000,
		}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		userStorage.EXPECT().
			GetById(gomock.Any(), id, gomock.Any()).
			Return(expectedUser, nil)
		service := New(Dependency{UserStorage: userStorage})

		resp, err := service.GetByID(context.Background(), GetByIDRequest{ID: id})
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, resp.User)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userStorage := mock_service.NewMockUserStorage(ctrl)
		userStorage.EXPECT().
			GetById(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(userDomain.User{}, userDomain.ErrUserNotFound)
		service := New(Dependency{UserStorage: userStorage})

		_, err := service.GetByID(context.Background(), GetByIDRequest{ID: "1"})
		assert.Error(t, err)
		assert.ErrorIs(t, err, userDomain.ErrUserNotFound)
	})
}
