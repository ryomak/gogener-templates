package repository

import (
	"context"

	"[[.ModName]]/domain/entity"
	"[[.ModName]]/domain/repository"
	"[[.ModName]]/infrastructure/db"
)

func NewMaintenance(
	conn *db.Conn,
) repository.Maintenance {
	return &maintenance{
		conn: conn,
	}
}

type maintenance struct {
	conn *db.Conn
}

func (m *maintenance) GetByName(ctx context.Context, name string) (*entity.Maintenance, error) {
	return entity.Maintenances(entity.MaintenanceWhere.Name.EQ(name)).One(ctx, m.conn)
}
