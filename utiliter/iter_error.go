package utiliter

import "iter"

// ErrorSeq2 creates a 2-value iterator that yields a single (zero-value, error) pair.
// Useful for returning early from a sequence pipeline with an error.
func ErrorSeq2[K any](err error) iter.Seq2[K, error] {
	var null K
	return func(yield func(K, error) bool) { yield(null, err) }
}
