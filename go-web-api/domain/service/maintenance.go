package service

import (
	"context"

	"[[.ModName]]/domain/repository"
	"[[.ModName]]/pkg/log"
	"[[.ModName]]/pkg/merr"
)

type Maintenance interface {
	CheckByName(ctx context.Context, name string) error
}

type maintenance struct {
	maintenanceRepository repository.Maintenance
}

func NewMaintenance(
	maintenanceRepository repository.Maintenance,
) Maintenance {
	return maintenance{
		maintenanceRepository: maintenanceRepository,
	}
}

func (m maintenance) CheckByName(ctx context.Context, name string) error {
	ent, err := m.maintenanceRepository.GetByName(ctx, name)
	switch {
	case merr.IsErrNoRows(err):
	case err != nil:
		log.Errorf(ctx, "Maintenance.CheckByName.Service: %v", err)
		return merr.CreateError(err.Error(), merr.CodeDBMaintenance)
	case ent.IsActive:
		return merr.ErrMaintenance
	}
	return nil
}
