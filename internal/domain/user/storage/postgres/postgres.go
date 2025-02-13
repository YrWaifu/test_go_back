package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/YrWaifu/test_go_back/internal/domain/user/storage"
	"github.com/YrWaifu/test_go_back/pkg/transaction"

	sq "github.com/Masterminds/squirrel"
	userDomain "github.com/YrWaifu/test_go_back/internal/domain/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) GetByUsername(ctx context.Context, username string, opts storage.GetOptions) (userDomain.User, error) {
	return s.get(ctx, "", username, opts)
}

func (s *Storage) Create(ctx context.Context, user userDomain.User) (string, error) {
	query := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("users").
		Columns("username", "password_hash", "balance").
		Values(user.Username, user.PasswordHash, user.Balance).
		Suffix("RETURNING \"id\"")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return "", fmt.Errorf("building query: %w", err)
	}

	var id string
	err = s.db.QueryRow(ctx, sqlQuery, args...).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("raw scan: %w", err)
	}

	return id, nil
}

func (s *Storage) GetById(ctx context.Context, id string, opts storage.GetOptions) (userDomain.User, error) {
	return s.get(ctx, id, "", opts)
}

func (s *Storage) get(ctx context.Context, id, username string, opts storage.GetOptions) (userDomain.User, error) {
	query := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "username", "password_hash", "balance").
		From("users")

	if id != "" {
		query = query.Where(sq.Eq{"id": id})
	} else {
		query = query.Where(sq.Eq{"username": username})
	}

	if opts.ForUpdate {
		query = query.Suffix("FOR UPDATE")
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return userDomain.User{}, fmt.Errorf("building query: %w", err)
	}

	var user userDomain.User
	if tx, txErr := transaction.ExtractTx(ctx); txErr != nil && tx != nil {
		err = tx.QueryRow(ctx, sqlQuery, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Balance)
	} else {
		err = s.db.QueryRow(ctx, sqlQuery, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Balance)
	}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return userDomain.User{}, userDomain.ErrUserNotFound
		}
		return userDomain.User{}, fmt.Errorf("raw scan: %w", err)
	}

	return user, nil
}

func (s *Storage) IncrementBalance(ctx context.Context, username string, inc int) error {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update("users").
		Set("balance", sq.Expr("balance + (?)", inc)).
		Where(sq.Eq{"username": username})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("building query: %w", err)
	}

	tx, err := transaction.ExtractTx(ctx)
	if err != nil {
		return fmt.Errorf("extracting tx: %w", err)
	}

	_, err = tx.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return fmt.Errorf("incrementing balance: %w", err)
	}

	return nil
}
