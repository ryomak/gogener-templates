package repository

import (
	"context"

	"[[.ModName]]/domain/entity"
	"[[.ModName]]/domain/entityx"
	"[[.ModName]]/domain/repository"
	"[[.ModName]]/infrastructure/db"
	"[[.ModName]]/infrastructure/firebase/auth"
	"github.com/volatiletech/sqlboiler/boil"
)

func NewMe(
	conn *db.Conn,
	authClient auth.Client,
) repository.Me {
	return &me{
		conn:       conn,
		authClient: authClient,
	}
}

type me struct {
	conn       *db.Conn
	authClient auth.Client
}

func (m *me) GetByAuthID(ctx context.Context, authID string) (*entityx.Me, error) {
	user, err := entity.Users(
		entity.UserWhere.AuthID.EQ(authID),
	).One(ctx, m.conn)
	if err != nil {
		return nil, err
	}
	firebaseUser, err := m.authClient.GetUserByAuthID(ctx, user.AuthID)
	if err != nil {
		return nil, err
	}
	return &entityx.Me{
		User:         user,
		FirebaseUser: firebaseUser,
	}, nil
}

func (m *me) CreateUser(ctx context.Context, ent *entity.User) error {
	return ent.Insert(ctx, m.conn.DB, boil.Infer())
}
