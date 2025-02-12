package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInit, downInit)
}

func upInit(ctx context.Context, tx *sql.Tx) error {
	_, error := tx.Exec(`
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(256) NOT NULL,
		password_hash TEXT NOT NULL,
		balance INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	`)
	return nil
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
