package auth

import (
	"context"
	"errors"
	"fmt"

	authDomain "github.com/YrWaifu/test_go_back/internal/domain/auth"
	authService "github.com/YrWaifu/test_go_back/internal/domain/auth/service"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
)

type Usecase struct {
	d Dependency
}

func New(d Dependency) *Usecase {
	return &Usecase{d: d}
}

type SignInRequest struct {
	Username string
	Password string
}

type SignInResponse struct {
	AccessToken string
}

func (u *Usecase) SignIn(ctx context.Context, req SignInRequest) (SignInResponse, error) {
	resp, err := u.d.AuthService.SignIn(ctx, authService.SignInRequest{Username: req.Username, Password: req.Password})

	if err != nil && errors.Is(err, userDomain.ErrUserNotFound) {
		_, err := u.d.AuthService.SignUp(ctx, authService.SignUpRequest{Username: req.Username, Password: req.Password})
		if err != nil {
			return SignInResponse{}, fmt.Errorf("auth service sign up: %w", err)
		}

		resp, err = u.d.AuthService.SignIn(ctx, authService.SignInRequest{Username: req.Username, Password: req.Password})
		if err != nil {
			return SignInResponse{}, fmt.Errorf("auth service sign in: %w", err)
		}

	} else if err != nil {
		if errors.Is(err, authDomain.ErrInvalidPassword) {
			return SignInResponse{}, userDomain.ErrUserNotFound
		}

		return SignInResponse{}, fmt.Errorf("auth service sign in: %w", err)
	}

	return SignInResponse{AccessToken: resp.AccessToken}, nil
}
