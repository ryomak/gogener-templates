package auth

import (
	"context"

	"[[.ModName]]/pkg/log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"[[.ModName]]/domain/entityx"
	"[[.ModName]]/infrastructure/env"
	"[[.ModName]]/pkg/merr"
)

type Client interface {
	VerifyIDToken(ctx context.Context, idToken string) (*string, error)
	GetUserByAuthID(ctx context.Context, authID string) (*entityx.FirebaseUser, error)
}

type client struct {
	client *auth.Client
}

func New() (Client, error) {
	ctx := context.Background()
	if env.IsDevelopment() {
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			log.Errorf(ctx, "mock auth client")
			return &mockClient{}, nil
		}
		authClient, err := app.Auth(ctx)
		if err != nil {
			log.Errorf(ctx, "mock auth client")
			return &mockClient{}, nil
		}
		return &client{authClient}, nil
	}
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return &client{authClient}, nil
}

func (c *client) GetUserByAuthID(ctx context.Context, authID string) (*entityx.FirebaseUser, error) {
	record, err := c.client.GetUser(ctx, authID)
	if err != nil {
		return nil, err
	}
	mail := record.Email
	return &entityx.FirebaseUser{
		Mail:  &mail,
		Phone: &record.PhoneNumber,
	}, nil
}

func (c *client) VerifyIDToken(ctx context.Context, idToken string) (*string, error) {
	token, err := c.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, merr.CreateError(err.Error(), merr.CodeAuthFailed)
	}
	user, err := c.client.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}
	if user.PhoneNumber == "" {
		return nil, merr.ErrAuthPhoneUnregisterd
	}
	return &user.UID, nil
}
