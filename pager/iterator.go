package pager

import (
	"errors"
	"iter"
)

// NewIterator creates a new paginated iterator with validation.
// Returns an error if count is non-positive or from is negative.
func NewIterator[T any](page Pager) (Iterator[T], error) {
	if page.GetCount() <= 0 {
		return nil, errors.New("invalid count: must be > 0")
	}
	if page.GetFrom() < 0 {
		return nil, errors.New("invalid from: must be >= 0")
	}
	return &iterate[T]{
		Pager: page,
	}, nil
}

// NewIteratorRelaxed creates a new paginated iterator without validation.
// Useful when input is already trusted or validation is handled elsewhere.
func NewIteratorRelaxed[T any](page Pager) Iterator[T] {
	return &iterate[T]{
		Pager: page,
	}
}

// iterate is a generic implementation of a paginated iterator.
type iterate[T any] struct {
	Pager           // Embedded pager defining offset and count
	countIter int64 // Internal counter to track iteration progress
}

func (a *iterate[T]) GetCount() int64 { return a.Pager.GetCount() + 1 }
func (a *iterate[T]) Scan(seq iter.Seq2[T, error]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for item, err := range seq {
			a.countIter++
			if a.countIter > a.Pager.GetCount() || !yield(item, err) {
				return
			}
		}
	}
}
func (a *iterate[T]) HasNext() bool { return a.countIter > a.Pager.GetCount() }
