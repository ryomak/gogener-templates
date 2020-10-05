package context

import (
	"context"
)

var (
	authIDKey = "AUTH_ID"
)

func WithAuthID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, &authIDKey, userID)
}

func AuthID(ctx context.Context) string {
	if v := ctx.Value(&authIDKey); v != nil {
		return v.(string)
	}
	return ""
}
