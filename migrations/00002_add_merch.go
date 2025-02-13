package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddMerch, downAddMerch)
}

func upAddMerch(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE merch (
		id SERIAL PRIMARY KEY,
		name text NOT NULL,
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE UNIQUE INDEX merch_name_idx ON merch(name);
	`)
	if err != nil {
		return err
	}

	return nil
}

func downAddMerch(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
