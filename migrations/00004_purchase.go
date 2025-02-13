package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upPurchase, downPurchase)
}

func upPurchase(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE purchase (
		    user_id INTEGER NOT NULL,
		    merch_id INTEGER NOT NULL,
		    quantity INTEGER NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			
			PRIMARY KEY (user_id, merch_id),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (merch_id) REFERENCES merch(id)
		);
	`)
	if err != nil {

		return err
	}

	return nil
}

func downPurchase(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
