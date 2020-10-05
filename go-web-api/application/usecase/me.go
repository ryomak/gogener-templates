package usecase

import (
	"context"

	"[[.ModName]]/application/usecase/request"
	"[[.ModName]]/application/usecase/response"
	"[[.ModName]]/domain/repository"
)

type Me interface {
	Get(ctx context.Context, r *request.MeGet) (*response.MeGet, error)
}

type me struct {
	meRepository repository.Me
}

func NewMe(
	meRepository repository.Me,
) Me {
	return &me{
		meRepository: meRepository,
	}
}

func (m *me) Get(ctx context.Context, r *request.MeGet) (*response.MeGet, error) {
	me, err := m.meRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	return response.NewMeGet(me), nil
}
