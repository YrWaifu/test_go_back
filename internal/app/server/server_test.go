package server

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/YrWaifu/test_go_back/internal/app/server/api"
	"github.com/YrWaifu/test_go_back/internal/domain/auth/utils"
	"github.com/YrWaifu/test_go_back/migrations"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func setup() {
	sqlConnetionString := os.Getenv("TEST_SQL_CONNECTION_STRING")
	db, err := sql.Open("pgx", sqlConnetionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
		GRANT ALL ON SCHEMA public TO postgres;
		GRANT ALL ON SCHEMA public TO public;
	`)
	if err != nil {
		log.Fatal(err)
	}

	err = migrations.Up(db)
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown() {

}

func TestMain(m *testing.M) {

	setup()
	m.Run()
}

func PopulateUser(s *Server, testName string) {
	for i := 0; i < 10; i++ {
		req := api.SignInRequest{
			Username: fmt.Sprintf("user-%v-%d", testName, i),
			Password: fmt.Sprintf("user-%v-%d", testName, i),
		}

		reqJSON, _ := json.Marshal(req)

		httpRequest := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
	}
}

func TestAuth_SignIn(t *testing.T) {
	s := New(Config{
		SQLConnection:    os.Getenv("TEST_SQL_CONNECTION_STRING"),
		SecretKey:        "SECRET_KEY",
		AccessTokenDelay: time.Minute * 5,
	})
	err := s.Init(context.Background())
	if err != nil {
		panic(err)
	}

	t.Run("SuccessCreate", func(t *testing.T) {
		req := api.SignInRequest{
			Username: "test",
			Password: "test",
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		httpRequest := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusOK, httpResponse.Code)
	})

	t.Run("SuccessLogin", func(t *testing.T) {
		req := api.SignInRequest{
			Username: "test",
			Password: "test",
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		httpRequest := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusOK, httpResponse.Code)
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		req := api.SignInRequest{
			Username: "test",
			Password: "1",
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		httpRequest := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusUnauthorized, httpResponse.Code)
	})
}

func TestPurchaseMerch(t *testing.T) {
	s := New(Config{
		SQLConnection:    os.Getenv("TEST_SQL_CONNECTION_STRING"),
		SecretKey:        "SECRET_KEY",
		AccessTokenDelay: time.Minute * 5,
	})
	err := s.Init(context.Background())
	if err != nil {
		panic(err)
	}

	ctx := utils.InjectAuth(context.Background(), "1")

	t.Run("SuccessPurchase", func(t *testing.T) {
		httpRequest := httptest.NewRequestWithContext(ctx, http.MethodGet, "/api/buy/pink-hoody", nil)
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusOK, httpResponse.Code)

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusOK, httpResponse.Code)
	})

	t.Run("NotEnoughBalance", func(t *testing.T) {
		httpRequest := httptest.NewRequestWithContext(ctx, http.MethodGet, "/api/buy/pink-hoody", nil)
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusInternalServerError, httpResponse.Code)
	})

	t.Run("NotAuthorized", func(t *testing.T) {
		httpRequest := httptest.NewRequest(http.MethodGet, "/api/buy/pink-hoody", nil)
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusUnauthorized, httpResponse.Code)
	})
}

func TestSendCoin(t *testing.T) {
	s := New(Config{
		SQLConnection:    os.Getenv("TEST_SQL_CONNECTION_STRING"),
		SecretKey:        "SECRET_KEY",
		AccessTokenDelay: time.Minute * 5,
	})
	err := s.Init(context.Background())
	if err != nil {
		panic(err)
	}

	PopulateUser(s, "sendCoin")

	ctx := utils.InjectAuth(context.Background(), "2")

	t.Run("Success", func(t *testing.T) {
		req := api.SendCoinRequest{
			ReceiverName: "user-sendCoin-2",
			Amount:       100,
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		httpRequest := httptest.NewRequestWithContext(ctx, http.MethodPost, "/api/sendCoin", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusOK, httpResponse.Code)
	})

	t.Run("NotEnoughBalance", func(t *testing.T) {
		req := api.SendCoinRequest{
			ReceiverName: "user-sendCoin-2",
			Amount:       1000,
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		httpRequest := httptest.NewRequestWithContext(ctx, http.MethodPost, "/api/sendCoin", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusInternalServerError, httpResponse.Code)
	})

	t.Run("NotAuthorized", func(t *testing.T) {
		req := api.SendCoinRequest{
			ReceiverName: "user-sendCoin-2",
			Amount:       100,
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		httpRequest := httptest.NewRequest(http.MethodPost, "/api/sendCoin", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusUnauthorized, httpResponse.Code)
	})

	t.Run("InvalidReceiver", func(t *testing.T) {
		req := api.SendCoinRequest{
			ReceiverName: "invalid-receiver",
			Amount:       100,
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		httpRequest := httptest.NewRequestWithContext(ctx, http.MethodPost, "/api/sendCoin", bytes.NewBuffer(reqJSON))
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse := httptest.NewRecorder()

		s.router.ServeHTTP(httpResponse, httpRequest)
		assert.Equal(t, http.StatusInternalServerError, httpResponse.Code)
	})
}
