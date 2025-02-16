package postgres

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	transactionDomain "github.com/YrWaifu/test_go_back/internal/domain/transaction"
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

func (s *Storage) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
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

func (s *Storage) ListByUserID(ctx context.Context, userID string) ([]transactionDomain.Transaction, []transactionDomain.Transaction, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("t.sender_id", "t.receiver_id", "t.amount", "s.username as sender_name", "r.username as receiver_name").
		From("transactions t").
		Where(sq.Or{
			sq.Eq{"sender_id": userID},
			sq.Eq{"receiver_id": userID},
		}).
		LeftJoin("users r ON t.receiver_id = r.id").
		LeftJoin("users s ON t.sender_id = s.id")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("list transactions: %w", err)
	}

	rows, err := s.db.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("list transactions: %w", err)
	}

	var received, sent []transactionDomain.Transaction
	for rows.Next() {
		var t transactionDomain.Transaction

		if err := rows.Scan(&t.SenderId, &t.ReceiverId, &t.Amount, &t.SenderName, &t.ReceiverName); err != nil {
			return nil, nil, fmt.Errorf("list transactions: %w", err)
		}

		if t.SenderId == userID {
			sent = append(sent, t)
		} else {
			received = append(received, t)
		}
	}

	return sent, received, nil
}
