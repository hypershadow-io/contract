package db

import (
	"context"
)

// WithTx wraps a function that requires transactional execution, passing input and returning output.
// Begins a transaction, executes the callback with the transaction context, and commits or rolls back based on the result.
func WithTx[In any, Out any](
	getRepo func(c context.Context) Instance,
	cb func(c context.Context, in In) (Out, error),
) func(c context.Context, in In) (Out, error) {
	return func(c context.Context, in In) (Out, error) {
		var null Out
		repo := getRepo(c)
		c, err := repo.Begin(c)
		if err != nil {
			return null, err
		}
		defer func() { _ = repo.Rollback(c) }()
		out, err := cb(c, in)
		if err != nil {
			return null, err
		}
		return out, repo.Commit(c)
	}
}

// WithTxNoInput wraps a function that requires transactional execution without input.
// Begins a transaction, executes the callback with the transaction context, and commits or rolls back based on the result.
func WithTxNoInput[Out any](
	getRepo func(c context.Context) Instance,
	cb func(c context.Context) (Out, error),
) func(c context.Context) (Out, error) {
	return func(c context.Context) (Out, error) {
		var null Out
		repo := getRepo(c)
		c, err := repo.Begin(c)
		if err != nil {
			return null, err
		}
		defer func() { _ = repo.Rollback(c) }()
		out, err := cb(c)
		if err != nil {
			return null, err
		}
		return out, repo.Commit(c)
	}
}
