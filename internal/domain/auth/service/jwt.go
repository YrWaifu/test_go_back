package service

import (
	"context"
	"fmt"
	"time"

	authDomain "github.com/YrWaifu/test_go_back/internal/domain/auth"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey []byte
	expDelay  time.Duration
}

func (s *jwtService) createToken(ctx context.Context, user userDomain.User) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(s.expDelay)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("token sign: %w", err)
	}

	return token, nil
}

func (s *jwtService) verifyToken(ctx context.Context, tokenString string) (string, error) {
	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	}, jwt.WithTimeFunc(func() time.Time { return time.Now().UTC() })) // ZALUPA EBANAYA
	if err != nil {
		return "", authDomain.ErrInvalidToken
	}

	if !token.Valid {
		return "", authDomain.ErrInvalidToken
	}

	return claims.ID, nil
}
