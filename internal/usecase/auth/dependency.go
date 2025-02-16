package auth

//go:generate mockgen -source dependency.go -destination mock/dependency.go

import (
	"context"

	authService "github.com/YrWaifu/test_go_back/internal/domain/auth/service"
)

type AuthService interface {
	SignIn(ctx context.Context, req authService.SignInRequest) (authService.SignInResponse, error)
	SignUp(ctx context.Context, req authService.SignUpRequest) (authService.SignUpResponse, error)
}

type Dependency struct {
	AuthService AuthService
}
