package usecase

import (
	"context"
	"[[.ModName]]/src/domain/model"
	"[[.ModName]]/src/domain/repository"
)

type IRoomUseCase interface {
	GetRoom(context.Context, string, *model.User) (*model.Room, error)
}

type roomUseCase struct {
	transaction repository.ITransactionRepository
	roomRepo    repository.IRoomRepository
	userRepo    repository.IUserRepository
}

func NewIRoomUseCase(
	t repository.ITransactionRepository,
	room repository.IRoomRepository,
	user repository.IUserRepository) IRoomUseCase {
	return &roomUseCase{t, room, user}
}

func (r *roomUseCase) GetRoom(ctx context.Context, roomName string, user *model.User) (*model.Room, error) {

	room, err := r.roomRepo.Save(ctx, roomName)
	if err != nil {
		return nil, err
	}

	//自身のデータ更新
	user.RoomName = roomName
	r.userRepo.Save(ctx, user)

	room.Users, err = r.userRepo.GetListByRoom(ctx, roomName)
	if err != nil {
		return nil, err
	}
	return room, nil
}
