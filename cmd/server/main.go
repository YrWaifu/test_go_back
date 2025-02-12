package main

import (
	"context"
	"os"
	"time"

	"github.com/YrWaifu/test_go_back/internal/app/server"
)

func main() {

	s := server.New(server.Config{
		Addr:             os.Getenv("LISTEN_ADDR"),
		SQLConnection:    os.Getenv("SQL_CONNECTION_STRING"),
		SecretKey:        os.Getenv("SECRET_KEY"),
		AccessTokenDelay: mustParseDuration(os.Getenv("ACCESS_TOKEN_DELAY")),
	})
	s.Run(context.Background())
}

func mustParseDuration(s string) time.Duration {
	t, err := time.ParseDuration(s)
	if err != nil {
		t = time.Minute * 10
	}

	return t
}
