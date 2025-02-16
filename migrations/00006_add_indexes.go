package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddIndexes, downAddIndexes)
}

func upAddIndexes(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE INDEX ON transactions(sender_id);
		CREATE INDEX ON transactions(receiver_id);
		CREATE INDEX ON transactions(sender_id, receiver_id);
		CREATE INDEX ON purchase(user_id);
	`)
	if err != nil {

		return err
	}

	return nil
}

func downAddIndexes(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
