// +build wireinject

package di

import (
	"[[.ModName]]/application/usecase"
	"[[.ModName]]/domain/service"
	"[[.ModName]]/infrastructure/db"
	"[[.ModName]]/infrastructure/firebase/auth"
	"[[.ModName]]/infrastructure/repository"
	middle "[[.ModName]]/presentation/middleware"
	"[[.ModName]]/presentation/v1/handler"

	"github.com/google/wire"
)

func MeHandler(db *db.Conn) (handler.Me, error) {
	wire.Build(
		auth.New,
		repository.NewMe,
		usecase.NewMe,
		handler.NewMe,
	)
	return nil, nil
}

func AuthMiddleware(db *db.Conn) (*middle.AuthMiddleware, error) {
	wire.Build(
		auth.New,
		middle.NewAuth,
	)
	return nil, nil
}

func MaintenanceMiddleware(db *db.Conn) (*middle.MaintenanceMiddleware, error) {
	wire.Build(
		repository.NewMaintenance,
		service.NewMaintenance,
		middle.NewMaintenance,
	)
	return nil, nil
}
