package transaction

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

var ErrTxNotFound = errors.New("transaction not found")

type txKey struct{}

func InjectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func ExtractTx(ctx context.Context) (pgx.Tx, error) {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)
	if !ok {
		return nil, ErrTxNotFound
	}

	return tx, nil
}
