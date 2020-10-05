package repository

import (
	"context"

	"[[.ModName]]/domain/entityx"
)

type Me interface {
	GetByAuthID(ctx context.Context, authID string) (*entityx.Me, error)
}
