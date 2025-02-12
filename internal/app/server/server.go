package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/YrWaifu/test_go_back/internal/app/server/api"
	authservice "github.com/YrWaifu/test_go_back/internal/domain/auth/service"
	userstorage "github.com/YrWaifu/test_go_back/internal/domain/user/storage/postgres"
	authusecase "github.com/YrWaifu/test_go_back/internal/usecase/auth"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Addr             string
	SQLConnection    string
	SecretKey        string
	AccessTokenDelay time.Duration
}

type Server struct {
	config Config
}

func New(c Config) *Server {
	return &Server{config: c}
}

func (s *Server) Run(ctx context.Context) error {
	db, err := pgxpool.New(ctx, s.config.SQLConnection)
	if err != nil {
		return fmt.Errorf("database conn: %w", err)
	}

	userStorage := userstorage.New(db)

	authService := authservice.New(authservice.Dependency{UserStorage: userStorage}, s.config.SecretKey, s.config.AccessTokenDelay)

	authUsecase := authusecase.New(authusecase.Dependency{AuthService: authService})

	authAPI := api.NewAuthAPI(api.AuthDependency{AuthUsecase: authUsecase})

	r := chi.NewRouter()
	r.Post("/api/auth", authAPI.SignIn)

	http.ListenAndServe(s.config.Addr, r)
}
