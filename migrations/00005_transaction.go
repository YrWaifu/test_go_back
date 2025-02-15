package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upTransaction, downTransaction)
}

func upTransaction(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE transactions (
		    id SERIAL PRIMARY KEY,
		    sender_id INTEGER NOT NULL,
		    receiver_id INTEGER NOT NULL,
		    amount INTEGER NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			
			FOREIGN KEY (sender_id) REFERENCES users(id),
			FOREIGN KEY (receiver_id) REFERENCES users(id)
		);
	`)
	if err != nil {

		return err
	}

	return nil
}

func downTransaction(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
