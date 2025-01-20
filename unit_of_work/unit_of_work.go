package unit_of_work

import (
	"context"
)

type UnitOfWork interface {
	Begin(context context.Context) (Transaction, error)
}

type Transaction interface {
	context.Context
	Commit() error
	Rollback() error
}
