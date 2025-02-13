package postgres

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	merchDomain "github.com/YrWaifu/test_go_back/internal/domain/merch"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) GetByName(ctx context.Context, name string) (merchDomain.Merch, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id", "name", "price").
		From("merch").
		Where(sq.Eq{"name": name})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return merchDomain.Merch{}, fmt.Errorf("query building: %w", err)
	}

	var merch merchDomain.Merch
	err = s.db.QueryRow(ctx, sqlQuery, args...).Scan(&merch.ID, &merch.Name, &merch.Price)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return merchDomain.Merch{}, merchDomain.ErrMerchNotFound
		}
		return merchDomain.Merch{}, fmt.Errorf("query row: %w", err)
	}

	return merch, nil
}
