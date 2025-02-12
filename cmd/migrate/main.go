package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/YrWaifu/test_go_back/migrations"
)

func main() {
	sqlConnetionString := os.Getenv("SQL_CONNECTION_STRING")
	db, err := sql.Open("pgx", sqlConnetionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = migrations.Up(db)
	if err != nil {
		log.Fatal(err)
	}
}
