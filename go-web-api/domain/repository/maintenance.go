package repository

import (
	"context"

	"[[.ModName]]/domain/entity"
)

type Maintenance interface {
	GetByName(ctx context.Context, name string) (*entity.Maintenance, error)
}
