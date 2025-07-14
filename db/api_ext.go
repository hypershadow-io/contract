package db

import (
	"context"
	"iter"

	"github.com/hypershadow-io/contract/eb"
	"github.com/hypershadow-io/contract/utiliter"
)

// FindOne executes the given query using the provided database instance and attempts to decode a single result into type T.
// Returns the decoded result, a flag indicating if a result was found, and an error if one occurred.
func FindOne[T any](
	c context.Context,
	instance Instance,
	errBuilder eb.LazyBuilder,
	proto T,
	query Query,
) (res_ T, found_ bool, err_ error) {
	result, found, err := instance.FindOne(c, errBuilder, proto, query)
	if err != nil || !found {
		return res_, found, err
	}
	return result.(T), found, nil
}

// FindIterator executes the given query and returns an iterator over results decoded into type T.
// Errors are wrapped in the iterator and surfaced during iteration.
func FindIterator[T any](
	c context.Context,
	instance Instance,
	errBuilder eb.LazyBuilder,
	proto T,
	query Query,
) iter.Seq2[T, error] {
	return utiliter.Iter2ToIter2Err[T](instance.FindIterator(c, errBuilder, proto, query))
}
