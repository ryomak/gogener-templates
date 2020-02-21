package transform

import (
	"[[.ModName]]/src/domain/model"
	"[[.ModName]]/src/interface/rpc"
)

func TransformRoomModel(rRoom *rpc.RoomRequest) *model.Room {
	return &model.Room{
		Name: rRoom.GetName(),
		Users: []*model.User{
			TransformUserModel(rRoom.GetUser()),
		},
	}
}

func TransformRoomRpc(mRoom *model.Room) *rpc.RoomResponse {
	users := make([]*rpc.User, len(mRoom.Users))
	for i := 0; i < len(mRoom.Users); i++ {
		users[i] = TransformUserRpc(mRoom.Users[i])
	}
	return &rpc.RoomResponse{
		Name:  mRoom.Name,
		Users: users,
	}
}
