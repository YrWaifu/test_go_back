package auth

import (
	"context"
	authDomain "github.com/YrWaifu/test_go_back/internal/domain/auth"
	authservice "github.com/YrWaifu/test_go_back/internal/domain/auth/service"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	mock_auth "github.com/YrWaifu/test_go_back/internal/usecase/auth/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUsecase_SignIn(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := SignInRequest{
			Username: "test",
			Password: "test",
		}
		token := "test_token"

		authService := mock_auth.NewMockAuthService(ctrl)
		authService.EXPECT().
			SignIn(gomock.Any(), authservice.SignInRequest{Username: req.Username, Password: req.Password}).
			Return(authservice.SignInResponse{AccessToken: token}, nil)

		usecase := New(Dependency{AuthService: authService})

		resp, err := usecase.SignIn(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, token, resp.AccessToken)
	})

	t.Run("UserNotFoundThenSignUpAndSignIn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := SignInRequest{
			Username: "new_user",
			Password: "new_pass",
		}
		token := "new_token"

		authService := mock_auth.NewMockAuthService(ctrl)
		authService.EXPECT().
			SignIn(gomock.Any(), authservice.SignInRequest{Username: req.Username, Password: req.Password}).
			Return(authservice.SignInResponse{}, userDomain.ErrUserNotFound)
		authService.EXPECT().
			SignUp(gomock.Any(), authservice.SignUpRequest{Username: req.Username, Password: req.Password}).
			Return(authservice.SignUpResponse{}, nil)
		authService.EXPECT().
			SignIn(gomock.Any(), authservice.SignInRequest{Username: req.Username, Password: req.Password}).
			Return(authservice.SignInResponse{AccessToken: token}, nil)

		usecase := New(Dependency{AuthService: authService})

		resp, err := usecase.SignIn(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, token, resp.AccessToken)
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := SignInRequest{
			Username: "test_user",
			Password: "wrong_pass",
		}

		authService := mock_auth.NewMockAuthService(ctrl)
		authService.EXPECT().
			SignIn(gomock.Any(), authservice.SignInRequest{Username: req.Username, Password: req.Password}).
			Return(authservice.SignInResponse{}, authDomain.ErrInvalidPassword)

		usecase := New(Dependency{AuthService: authService})

		_, err := usecase.SignIn(context.Background(), req)
		assert.Error(t, err)
		assert.Equal(t, userDomain.ErrUserNotFound, err)
	})
}
