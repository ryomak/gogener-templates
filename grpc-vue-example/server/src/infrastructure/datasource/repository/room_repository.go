package repository

import (
	"context"
	"[[.ModName]]/src/domain/repository"
	"[[.ModName]]/src/infrastructure/datasource/database"
	"time"

	"[[.ModName]]/src/domain/model"
)

func NewIRoomRepository(dbm database.DBM) repository.IRoomRepository {
	return &roomRepository{dbm}
}

type roomRepository struct {
	dbm database.DBM
}

func (r *roomRepository) Save(ctx context.Context, roomName string) (*model.Room, error) {
	query := `insert into rooms (name) values (?) on DUPLICATE key update updated = ?`
	stmt, err := r.dbm.Prepare(ctx, query)
	_, err = stmt.ExecContext(ctx, roomName, time.Now())
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, roomName)
}

func (r *roomRepository) Get(ctx context.Context, name string) (*model.Room, error) {
	room := &model.Room{}
	query := "select * from rooms where name = ?"
	stmt, err := r.dbm.Prepare(ctx, query)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRowContext(ctx, name).Scan(&room.Name, &room.Created, &room.Updated)
	if err != nil {
		return nil, err
	}
	return room, nil
}
