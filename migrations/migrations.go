package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func Up(db *sql.DB) error {
	return goose.Up(db, "/")
}
