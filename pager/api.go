package pager

import "iter"

type (
	// Pager defines the base interface for pagination settings,
	// including the offset (From) and maximum number of results (Count).
	Pager interface {
		// GetFrom returns the offset from which to start retrieving results.
		GetFrom() int64

		// GetCount returns the maximum number of results to retrieve.
		GetCount() int64
	}

	// Iterator extends Pager to provide paginated iteration capabilities
	// with scanning and navigation control.
	Iterator[T any] interface {
		Pager

		// Scan applies pagination to the given input sequence and returns a new sequence.
		Scan(iter.Seq2[T, error]) iter.Seq2[T, error]

		// HasNext indicates whether there are more items available after the current page.
		HasNext() bool
	}
)
