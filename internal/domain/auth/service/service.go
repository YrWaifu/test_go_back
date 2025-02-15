package service

import (
	"context"
	"fmt"
	userStorage "github.com/YrWaifu/test_go_back/internal/domain/user/storage"
	"time"

	authDomain "github.com/YrWaifu/test_go_back/internal/domain/auth"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	d   Dependency
	jwt *jwtService
}

func New(d Dependency, secretKey string, expDelay time.Duration) *Service {
	return &Service{
		d:   d,
		jwt: &jwtService{expDelay: expDelay, secretKey: []byte(secretKey)},
	}
}

type SignInRequest struct {
	Username string
	Password string
}

type SignInResponse struct {
	AccessToken string
}

func (s *Service) SignIn(ctx context.Context, req SignInRequest) (SignInResponse, error) {
	user, err := s.d.UserStorage.GetByUsername(ctx, req.Username, userStorage.GetOptions{})
	if err != nil {
		return SignInResponse{}, fmt.Errorf("user storage get by username: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return SignInResponse{}, authDomain.ErrInvalidPassword
	}

	token, err := s.jwt.createToken(ctx, user)
	if err != nil {
		return SignInResponse{}, fmt.Errorf("create token: %w", err)
	}

	return SignInResponse{AccessToken: token}, nil
}

type SignUpRequest struct {
	Username string
	Password string
}

type SignUpResponse struct {
}

func (s *Service) SignUp(ctx context.Context, req SignUpRequest) (SignUpResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return SignUpResponse{}, fmt.Errorf("generate hash: %w", err)
	}

	_, err = s.d.UserStorage.Create(ctx, userDomain.User{
		Username:     req.Username,
		PasswordHash: string(hash),
		Balance:      userDomain.DEFAULT_BALANCE,
	})
	if err != nil {
		return SignUpResponse{}, fmt.Errorf("user storage create: %w", err)
	}

	return SignUpResponse{}, nil
}

func (s *Service) Authenticate(ctx context.Context, token string) (string, error) {
	id, err := s.jwt.verifyToken(ctx, token)
	if err != nil {
		return "", fmt.Errorf("verify token: %w", err)
	}

	return id, nil
}
