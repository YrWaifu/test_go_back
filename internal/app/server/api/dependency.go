package api

import (
	"context"

	usecase "github.com/YrWaifu/test_go_back/internal/usecase/auth"
)

type AuthUsecase interface {
	SignIn(ctx context.Context, req usecase.SignInRequest) (usecase.SignInResponse, error)
}

type AuthDependency struct {
	AuthUsecase AuthUsecase
}
