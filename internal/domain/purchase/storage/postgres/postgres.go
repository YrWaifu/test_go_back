package postgres

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	purchaseDomain "github.com/YrWaifu/test_go_back/internal/domain/purchase"
	"github.com/YrWaifu/test_go_back/pkg/transaction"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
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

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			slog.Info("rollback error")
		}
	}()

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
		Suffix("ON CONFLICT (user_id, merch_id) DO UPDATE SET quantity = purchase.quantity + 1")

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

func (s *Storage) ListByUserID(ctx context.Context, userID string) ([]purchaseDomain.Purchase, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("p.user_id", "p.merch_id", "p.quantity", "m.name").
		From("purchase p").
		Where(sq.Eq{"p.user_id": userID}).
		LeftJoin("merch m ON p.merch_id = m.id")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("query to sql: %w", err)
	}

	rows, err := s.db.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("query rows: %w", err)
	}

	var purchases []purchaseDomain.Purchase

	for rows.Next() {
		var p purchaseDomain.Purchase

		if err := rows.Scan(&p.UserID, &p.MerchID, &p.Quantity, &p.MerchName); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		purchases = append(purchases, p)
	}

	return purchases, nil
}
