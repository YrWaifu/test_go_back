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

func (s *Storage) BeginPurchase(ctx context.Context, fn func(context.Context) error) error {
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

func (s *Storage) CreatePurchase(ctx context.Context, userID string, merchID string) error {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("purchase").
		Columns("user_id", "merch_id", "quantity").
		Values(userID, merchID, 1).
		Suffix("ON CONFLICT (user_id, merch_id) DO UPDATE SET quantity = quantity + 1")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("create purchase statement: %w", err)
	}

	tx, err := transaction.ExtractTx(ctx)
	if err != nil {
		return fmt.Errorf("extracting tx: %w", err)
	}

	_, err = tx.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return fmt.Errorf("insert purchase: %w", err)
	}

	return nil
}
