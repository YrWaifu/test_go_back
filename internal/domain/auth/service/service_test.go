package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	authDomain "github.com/YrWaifu/test_go_back/internal/domain/auth"
	mock_service "github.com/YrWaifu/test_go_back/internal/domain/auth/service/mock"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
	"time"
)

func TestService_SignUp(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "1"
		req := SignUpRequest{
			Username: "test1",
			Password: "test1",
		}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		userStorage.EXPECT().
			Create(gomock.Any(), gomock.Cond(func(user userDomain.User) bool {
				err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
				return err == nil && user.Username == req.Username && user.Balance == 1000
			})).
			Return(id, nil)
		service := New(Dependency{UserStorage: userStorage}, "test", time.Minute*5)

		_, err := service.SignUp(context.Background(), req)
		assert.NoError(t, err)
	})
}

func TestService_Authenticate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwibmFtZSI6InRlc3QiLCJpYXQiOjU1MTYyMzkwMjJ9.PyVvW0U4rWxX70Nv8wvIijFtlBnIKYniMOnvJe2jlxg"

		service := New(Dependency{}, "test", time.Minute*5)

		id, err := service.Authenticate(context.Background(), token)
		assert.NoError(t, err)
		assert.Equal(t, id, "1")
	})

	t.Run("TokenTimedOut", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwibmFtZSI6InRlc3QiLCJleHAiOjE1MTYyMzkwMjJ9.u_BTPSmnSO48F9iwslCFWCFk2Rz6fwIbPn_sIJUL6uM"

		service := New(Dependency{}, "test", time.Minute*5)

		_, err := service.Authenticate(context.Background(), token)
		assert.Error(t, err)
		assert.ErrorIs(t, err, authDomain.ErrInvalidToken)
	})
}

func TestService_SignIn(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "1"
		req := SignInRequest{
			Username: "test1",
			Password: "test1",
		}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		userStorage.EXPECT().
			GetByUsername(gomock.Any(), req.Username, gomock.Any()).
			Return(userDomain.User{Username: req.Username, PasswordHash: mustHash(req.Password), ID: id}, nil)
		service := New(Dependency{UserStorage: userStorage}, "test", time.Minute*5)

		resp, err := service.SignIn(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, id, mustDecodeToken(resp.AccessToken))
	})

	t.Run("UserNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		req := SignInRequest{
			Username: "test1",
			Password: "test1",
		}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		userStorage.EXPECT().
			GetByUsername(gomock.Any(), req.Username, gomock.Any()).
			Return(userDomain.User{}, userDomain.ErrUserNotFound)
		service := New(Dependency{UserStorage: userStorage}, "test", time.Minute*5)

		_, err := service.SignIn(context.Background(), req)
		assert.Error(t, err)
		assert.ErrorIs(t, err, userDomain.ErrUserNotFound)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "1"
		req := SignInRequest{
			Username: "test1",
			Password: "test1",
		}

		userStorage := mock_service.NewMockUserStorage(ctrl)
		userStorage.EXPECT().
			GetByUsername(gomock.Any(), req.Username, gomock.Any()).
			Return(userDomain.User{Username: req.Username, PasswordHash: mustHash("test0"), ID: id}, nil)
		service := New(Dependency{UserStorage: userStorage}, "test", time.Minute*5)

		_, err := service.SignIn(context.Background(), req)
		assert.Error(t, err)
		assert.ErrorIs(t, err, authDomain.ErrInvalidPassword)
	})
}

func mustHash(p string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func mustDecodeToken(token string) string {
	bytes, err := base64.RawURLEncoding.DecodeString(strings.Split(token, ".")[1])
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}

	rawID, ok := data["sub"]
	if !ok {
		panic("id not found")
	}

	id, ok := rawID.(string)
	if !ok {
		panic("id invalid type")
	}

	return id
}
