package transform

import (
	"[[.ModName]]/src/domain/model"
	"[[.ModName]]/src/interface/rpc"
)

func TransformCoordinateModel(rCoordinate *rpc.Coordinate) *model.Coordinate {
	return &model.Coordinate{
		Longitude: rCoordinate.GetLongitude(),
		Latitude:  rCoordinate.GetLatitude(),
	}
}

func TransformCoordinateRpc(mCoordinate *model.Coordinate) *rpc.Coordinate {
	return &rpc.Coordinate{
		Longitude: mCoordinate.Longitude,
		Latitude:  mCoordinate.Latitude,
	}
}
