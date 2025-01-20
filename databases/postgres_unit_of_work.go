package databases

import (
	"context"
	"database/sql"
	"time"
)

func PostgresTransactionCtxName() string {
	return "postgresTransactionCtx"
}

type PostgresUnitOfWork struct {
	conn *sql.DB
}

func NewPostgreUnitOfWork(conn *sql.DB) *PostgresUnitOfWork {
	return &PostgresUnitOfWork{conn: conn}
}

func (u *PostgresUnitOfWork) Begin(ctx context.Context) (Transaction, error) {
	if tx, ok := ctx.Value(PostgresTransactionCtxName()).(*sql.Tx); ok {
		return &PostgresdbTransaction{tx: tx, ctx: ctx, isNested: true, alreadyCommited: false}, nil
	}

	tx, err := u.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	txCtx := context.WithValue(ctx, PostgresTransactionCtxName(), tx)
	return &PostgresdbTransaction{tx: tx, ctx: txCtx, isNested: false, alreadyCommited: false}, nil
}

type PostgresdbTransaction struct {
	tx              *sql.Tx
	ctx             context.Context
	isNested        bool
	alreadyCommited bool
}

func (t *PostgresdbTransaction) Deadline() (deadline time.Time, ok bool) {
	return t.ctx.Deadline()
}

func (t *PostgresdbTransaction) Done() <-chan struct{} {
	return t.ctx.Done()
}

func (t *PostgresdbTransaction) Err() error {
	return t.ctx.Err()
}

func (t *PostgresdbTransaction) Value(key any) any {
	return t.ctx.Value(key)
}

func (t *PostgresdbTransaction) Commit() error {
	if t.isNested || t.alreadyCommited {
		return nil
	}

	return t.tx.Commit()
}

func (t *PostgresdbTransaction) Rollback() error {
	if t.isNested || t.alreadyCommited {
		return nil
	}

	return t.tx.Rollback()
}

func (t *PostgresdbTransaction) Context() context.Context {
	return t.ctx
}
