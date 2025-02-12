package postgres

import (
	"context"
	"errors"
	"fmt"

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

func (s *Storage) GetByUsername(ctx context.Context, username string) (userDomain.User, error) {
	return s.get(ctx, "", username)
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
	err = s.db.QueryRow(ctx, sqlQuery, args).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("raw scan: %w", err)
	}

	return id, nil
}

func (s *Storage) GetById(ctx context.Context, id string) (userDomain.User, error) {
	return s.get(ctx, id, "")
}

func (s *Storage) get(ctx context.Context, id, username string) (userDomain.User, error) {
	query := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "username", "password_hash", "balance").
		From("users")

	if id != "" {
		query = query.Where(sq.Eq{"id": id})
	} else {
		query = query.Where(sq.Eq{"username": username})
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return userDomain.User{}, fmt.Errorf("building query: %w", err)
	}

	var user userDomain.User
	err = s.db.QueryRow(ctx, sqlQuery, args).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Balance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return userDomain.User{}, userDomain.ErrUserNotFound
		}
		return userDomain.User{}, fmt.Errorf("raw scan: %w", err)
	}

	return user, nil
}
