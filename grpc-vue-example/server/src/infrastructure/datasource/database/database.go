package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dm DBM

// init initialize
func init() {
	dm = &dbm{txMap: map[string]*sql.Tx{}}
}

// GetDBM get database manager
func GetDBM() DBM {
	return dm
}

// ContextKey key for transaction context
type ContextKey string

const (
	txIsKey = "db.tx"
	txKey   = "db.tx.key"
)

type DBM interface {
	InitDB(path, env string) error
	CloseDB() error
	Begin(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	Prepare(ctx context.Context, query string) (*sql.Stmt, error)
}

// dbm database manager
type dbm struct {
	db    *sql.DB
	txMap map[string]*sql.Tx
	stmt  *sql.Stmt
}

// closeDB close database
func (m *dbm) closeDB() error {
	m.stmt.Close()
	return m.db.Close()
}

// Begin begin transaction
func (m *dbm) Begin(ctx context.Context) (context.Context, error) {
	t, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	ctx = m.addTx(ctx, t)

	return ctx, nil
}

// Rollback rollback transaction
func (m *dbm) Rollback(ctx context.Context) (context.Context, error) {
	m.getTx(ctx).Rollback()

	// Transaction関連削除
	ctx = m.deleteTx(ctx)

	return ctx, nil
}

// Commit commit transaction
func (m *dbm) Commit(ctx context.Context) (context.Context, error) {
	err := m.getTx(ctx).Commit()

	// Transaction関連削除
	ctx = m.deleteTx(ctx)

	return ctx, err
}

// Prepare prepare statement
func (m *dbm) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	var err error
	if m.isTx(ctx) {
		m.stmt, err = m.getTx(ctx).PrepareContext(ctx, query)
	} else {
		m.stmt, err = m.db.PrepareContext(ctx, query)
	}
	if err != nil {
		return m.stmt, err
	}

	return m.stmt, nil
}

// isTx is in transaction or not
func (m *dbm) isTx(ctx context.Context) bool {
	if txn, ok := ctx.Value(ContextKey(txIsKey)).(bool); ok {
		return txn
	}
	return false
}

// getTx
func (m *dbm) getTx(ctx context.Context) *sql.Tx {
	key := m.getTxKey(ctx)
	return m.txMap[key]
}

// addTx
func (m *dbm) addTx(ctx context.Context, t *sql.Tx) context.Context {
	key := m.generateNewKey()
	m.txMap[key] = t
	ctx = m.setTxKey(ctx, key)

	// Transaction開始フラグ
	ctx = context.WithValue(ctx, ContextKey(txIsKey), true)

	return ctx
}

// deleteTx
func (m *dbm) deleteTx(ctx context.Context) context.Context {
	txKey := m.getTxKey(ctx)
	if _, ok := m.txMap[txKey]; ok {
		delete(m.txMap, txKey)
	}

	// Transaction開始フラグ
	return context.WithValue(ctx, ContextKey(txIsKey), false)
}

// setTxKey
func (m *dbm) setTxKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextKey(txKey), value)
}

// getTxKey
func (m *dbm) getTxKey(ctx context.Context) string {
	return m.getKey(ctx, ContextKey(txKey))
}

// getKey get key
func (m *dbm) getKey(ctx context.Context, ctxKey ContextKey) string {
	key, _ := ctx.Value(ctxKey).(string)
	return key
}

// generateNewKey generate key
func (m *dbm) generateNewKey() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Int())
}

// InitDB DB接続
func (m *dbm) InitDB(path, env string) error {
	cs, err := NewConfigsFromFile(path)
	if err != nil {
		log.Printf("cannot open database configuration. -> %s", err)
		return err
	}
	m.db, err = cs.OpenDB(env)
	if err != nil {
		log.Printf("cannot open database. -> %s", err)
		return err
	}
	return nil
}

// CloseDB DB切断
func (m *dbm) CloseDB() error {
	return m.closeDB()
}
