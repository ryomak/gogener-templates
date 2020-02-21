package repository

import (
	"context"
	"[[.ModName]]/src/domain/model"
)

type IUserRepository interface {
	Save(ctx context.Context, user *model.User) (*model.User, error)
	GetByNames(ctx context.Context, userName, roomName string) (*model.User, error)
	GetListByRoom(ctx context.Context, roomName string) ([]*model.User, error)
}
