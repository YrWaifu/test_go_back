package server

import (
	"context"
	"fmt"
	"github.com/YrWaifu/test_go_back/internal/domain/auth/utils"
	"net/http"
	"time"

	"github.com/YrWaifu/test_go_back/internal/app/server/api"
	authservice "github.com/YrWaifu/test_go_back/internal/domain/auth/service"
	merchstorage "github.com/YrWaifu/test_go_back/internal/domain/merch/storage/postgres"
	purchaseservice "github.com/YrWaifu/test_go_back/internal/domain/purchase/service"
	purchasestorage "github.com/YrWaifu/test_go_back/internal/domain/purchase/storage/postgres"
	transactionservice "github.com/YrWaifu/test_go_back/internal/domain/transaction/service"
	transactionstorage "github.com/YrWaifu/test_go_back/internal/domain/transaction/storage/postgres"
	userstorage "github.com/YrWaifu/test_go_back/internal/domain/user/storage/postgres"
	authusecase "github.com/YrWaifu/test_go_back/internal/usecase/auth"
	purchaseusecase "github.com/YrWaifu/test_go_back/internal/usecase/purchase"
	transactionusecase "github.com/YrWaifu/test_go_back/internal/usecase/transaction"
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
	purchaseStorage := purchasestorage.New(db)
	merchStorage := merchstorage.New(db)
	transactionStorage := transactionstorage.New(db)

	authService := authservice.New(authservice.Dependency{UserStorage: userStorage}, s.config.SecretKey, s.config.AccessTokenDelay)
	purchaseService := purchaseservice.New(purchaseservice.Dependency{
		PurchaseStorage: purchaseStorage,
		UserStorage:     userStorage,
		MerchStorage:    merchStorage,
	})
	transactionService := transactionservice.New(transactionservice.Dependency{
		TransactionStorage: transactionStorage,
		UserStorage:        userStorage,
	})

	authUsecase := authusecase.New(authusecase.Dependency{AuthService: authService})
	purchaseUsecase := purchaseusecase.New(purchaseusecase.Dependency{PurchaseService: purchaseService})
	transactionUsecase := transactionusecase.New(transactionusecase.Dependency{TransactionService: transactionService})

	authAPI := api.NewAuthAPI(api.AuthDependency{AuthUsecase: authUsecase})
	purchaseAPI := api.NewPurchaseAPI(api.PurchaseDependency{PurchaseUsecase: purchaseUsecase})
	transactionAPI := api.NewTransactionAPI(api.TransactionDependency{TransactionUsecase: transactionUsecase})

	r := chi.NewRouter()
	r.Use(utils.AuthMiddleware(authService.Authenticate))
	r.Post("/api/auth", authAPI.SignIn)
	r.Route("/", func(r chi.Router) {
		r.Use(utils.AuthRequiredMiddleware())
		r.Get("/api/buy/{merchName}", purchaseAPI.PurchaseMerch)
		r.Post("/api/sendCoin", transactionAPI.SendCoin)
	})

	http.ListenAndServe(s.config.Addr, r)

	return nil
}
