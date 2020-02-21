package repository

import (
	"context"
	"[[.ModName]]/src/domain/model"
)

type IRoomRepository interface {
	Save(ctx context.Context, roomName string) (*model.Room, error)
	Get(ctx context.Context, name string) (*model.Room, error)
}
