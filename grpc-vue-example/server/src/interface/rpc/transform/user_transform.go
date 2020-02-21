package transform

import (
	"[[.ModName]]/src/domain/model"
	"[[.ModName]]/src/interface/rpc"
)

func TransformUserModel(gUser *rpc.User) *model.User {
	return &model.User{
		Name:       gUser.GetName(),
		Coordinate: TransformCoordinateModel(gUser.GetCoordinate()),
	}
}

func TransformUserRpc(mUser *model.User) *rpc.User {
	return &rpc.User{
		Name:       mUser.Name,
		Coordinate: TransformCoordinateRpc(mUser.Coordinate),
	}
}
