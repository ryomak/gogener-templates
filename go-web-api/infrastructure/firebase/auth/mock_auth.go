package auth

import (
	"context"

	"[[.ModName]]/domain/entityx"
)

type mockClient struct {
}

func (c *mockClient) GetUserByAuthID(ctx context.Context, authID string) (*entityx.FirebaseUser, error) {
	mail := "mock@mock.com"
	Phone := "+09011111111"
	return &entityx.FirebaseUser{
		Mail:  &mail,
		Phone: &Phone,
	}, nil
}

func (c *mockClient) VerifyIDToken(ctx context.Context, uid string) (*string, error) {
	return &uid, nil
}
