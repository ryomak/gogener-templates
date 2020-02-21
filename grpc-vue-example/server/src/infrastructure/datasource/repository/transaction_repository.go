package repository

import (
	"context"
	"[[.ModName]]/src/domain/repository"
	"[[.ModName]]/src/infrastructure/datasource/database"
)

// NewTransactionRepository get repository
func NewITransactionRepository(dbm database.DBM) repository.ITransactionRepository {
	return transactionRepository{dbm}
}

// transactionRepository transaction Repository Sub
type transactionRepository struct {
	dbm database.DBM
}

// Begin begin transaction
func (t transactionRepository) Begin(ctx context.Context) (context.Context, error) {
	return (t.dbm).Begin(ctx)
}

// Rollback rollback transaction
func (t transactionRepository) Rollback(ctx context.Context) (context.Context, error) {
	return t.dbm.Rollback(ctx)
}

// Commit commit transaction
func (t transactionRepository) Commit(ctx context.Context) (context.Context, error) {
	return t.dbm.Commit(ctx)
}
