package dbhook

import (
	"context"

	"github.com/hypershadow-io/contract/hook"
	"github.com/hypershadow-io/contract/qb"
)

// Select applies mutation hooks to a SELECT query.
func Select(
	c context.Context, kinds hook.Kinds, provider Provider[qb.SelectQuery], value qb.SelectQuery,
) qb.SelectQuery {
	return run[qb.SelectQuery](c, kinds, provider, value)
}

// Insert applies mutation hooks to an INSERT query.
func Insert(
	c context.Context, kinds hook.Kinds, provider Provider[qb.InsertQuery], value qb.InsertQuery,
) qb.InsertQuery {
	return run[qb.InsertQuery](c, kinds, provider, value)
}

// Replace applies mutation hooks to a REPLACE query (semantically treated like INSERT).
func Replace(
	c context.Context, kinds hook.Kinds, provider Provider[qb.InsertQuery], value qb.InsertQuery,
) qb.InsertQuery {
	return run[qb.InsertQuery](c, kinds, provider, value)
}

// Update applies mutation hooks to an UPDATE query.
func Update(
	c context.Context, kinds hook.Kinds, provider Provider[qb.UpdateQuery], value qb.UpdateQuery,
) qb.UpdateQuery {
	return run[qb.UpdateQuery](c, kinds, provider, value)
}

// Delete applies mutation hooks to a DELETE query.
func Delete(
	c context.Context, kinds hook.Kinds, provider Provider[qb.DeleteQuery], value qb.DeleteQuery,
) qb.DeleteQuery {
	return run[qb.DeleteQuery](c, kinds, provider, value)
}

// run sequentially applies all registered mutators for the given query type and kinds.
// If a mutator returns an error, it is attached to the result using SetError.
func run[T qb.SetError[T]](
	c context.Context,
	kinds hook.Kinds,
	provider Provider[T],
	value T,
) T {
	for handler := range provider.Find(c, kinds, value) {
		patch, err := handler(c, kinds, value)
		if err != nil {
			return patch.SetError(err)
		}
		value = patch
	}
	return value
}

// Provider defines a hook provider for query mutators of type V.
type Provider[V any] = hook.Provider[hook.MutatorFunc[V], V]
