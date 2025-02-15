package utils

import "context"

type authKey struct{}

func InjectAuth(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, authKey{}, id)
}

func ExtractAuth(ctx context.Context) (string, bool) {
	v := ctx.Value(authKey{})
	id, ok := v.(string)
	return id, ok
}
