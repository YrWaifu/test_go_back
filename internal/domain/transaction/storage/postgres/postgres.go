package postgres

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/YrWaifu/test_go_back/pkg/transaction"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	ctx = transaction.InjectTx(ctx, tx)

	err = fn(ctx)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateTransaction(ctx context.Context, senderID string, receiverID string, amount int) error {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("transactions").
		Columns("sender_id", "receiver_id", "amount").
		Values(senderID, receiverID, amount)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("create transaction statement: %w", err)
	}

	tx, err := transaction.ExtractTx(ctx)
	if err != nil {
		return fmt.Errorf("extracting tx: %w", err)
	}

	_, err = tx.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return fmt.Errorf("insert transaction: %w", err)
	}

	return nil
}
