package repository

import (
	"context"
	"[[.ModName]]/src/domain/repository"
	"[[.ModName]]/src/infrastructure/datasource/database"

	"[[.ModName]]/src/domain/model"
)

func NewIUserRepository(dbm database.DBM) repository.IUserRepository {
	return &userRepository{dbm}
}

type userRepository struct {
	dbm database.DBM
}

func (u *userRepository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	query := `insert into users (room_name,user_name,latitude,longitude) values
	 (?,?,?,?) on DUPLICATE key update latitude = ? , longitude = ?`
	stmt, err := u.dbm.Prepare(ctx, query)
	_, err = stmt.ExecContext(
		ctx,
		user.RoomName,
		user.Name,
		user.Coordinate.Latitude,
		user.Coordinate.Longitude,
		user.Coordinate.Latitude,
		user.Coordinate.Longitude)
	if err != nil {
		return nil, err
	}
	return u.GetByNames(ctx, user.Name, user.RoomName)
}

func (u *userRepository) GetByNames(ctx context.Context, userName, roomName string) (*model.User, error) {
	user := new(model.User)
	user.Coordinate = new(model.Coordinate)
	query := "select * from users where user_name = ? and room_name =?"
	stmt, err := u.dbm.Prepare(ctx, query)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRowContext(ctx, userName, roomName).Scan(
		&user.Name,
		&user.RoomName,
		&user.Coordinate.Latitude,
		&user.Coordinate.Longitude,
		&user.Created,
		&user.Updated,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetListByRoom(ctx context.Context, roomName string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	query := "select * from users where room_name =?"
	stmt, err := u.dbm.Prepare(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, roomName)
	for rows.Next() {
		user := new(model.User)
		user.Coordinate = new(model.Coordinate)
		err := rows.Scan(
			&user.Name,
			&user.RoomName,
			&user.Coordinate.Latitude,
			&user.Coordinate.Longitude,
			&user.Created,
			&user.Updated,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
 	return users, nil
}
