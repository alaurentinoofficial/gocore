package handlers

import (
	"context"
)

type HandlerFunc[T any, U any] func(
	ctx context.Context,
	request T,
) (*U, error)
