package db

import (
	"context"
	"database/sql"

	"[[.ModName]]/infrastructure/env"

	_ "github.com/go-sql-driver/mysql"
)

var (
	global *Conn
)

type Conn struct {
	*sql.DB
}

func New() *Conn {
	db, err := sql.Open("mysql", env.GoogleCloudSqlSource())
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(8)
	conn := &Conn{db}
	global = conn
	return conn
}

type Tx struct {
	*sql.Tx
}

func (c *Conn) RunInTx(ctx context.Context, f func(tx *Tx) error) error {
	tx, err := c.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := f(&Tx{tx}); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func RunInTx(ctx context.Context, f func(tx *Tx) error) error {
	return global.RunInTx(ctx, f)
}
