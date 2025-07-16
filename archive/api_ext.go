package archive

import (
	"context"
	"iter"

	"github.com/hypershadow-io/contract/db"
	"github.com/hypershadow-io/contract/utiliter"
)

// AddMany wraps Instance.AddMany and casts the result to a typed sequence.
// Returns a sequence of typed results or errors.
func AddMany[T any](
	c context.Context,
	instance Instance,
	errBuilder func() error,
	customerId int64,
	query db.Query,
) iter.Seq2[T, error] {
	return utiliter.Iter2ToIter2Err[T](instance.AddMany(c, errBuilder, customerId, query))
}

// GetMany wraps Instance.GetMany and casts the result to a typed sequence.
// Returns a sequence of typed results or errors for the given entity IDs.
func GetMany[T any](
	c context.Context,
	instance Instance,
	errBuilder func() error,
	entityId []int64,
) iter.Seq2[T, error] {
	return utiliter.Iter2ToIter2Err[T](instance.GetMany(c, errBuilder, entityId))
}

// AddOne wraps Instance.AddOne and casts the result to a typed value.
// Returns the result, a found flag, and an error if occurred.
func AddOne[T any](
	c context.Context,
	instance Instance,
	errBuilder func() error,
	customerId int64,
	query db.Query,
) (res_ T, found_ bool, err_ error) {
	result, found, err := instance.AddOne(c, errBuilder, customerId, query)
	if err != nil || !found {
		return res_, found, err
	}
	return result.(T), true, nil
}

// GetOne wraps Instance.GetOne and casts the result to a typed value.
// Returns the result, a found flag, and an error if occurred.
func GetOne[T any](
	c context.Context,
	instance Instance,
	errBuilder func() error,
	entityId int64,
) (res_ T, found_ bool, err_ error) {
	result, found, err := instance.GetOne(c, errBuilder, entityId)
	if err != nil || !found {
		return res_, found, err
	}
	return result.(T), true, nil
}
