package utiliter

import "iter"

// Iter2ToSlice converts a 2-value iterator into a slice using the provided mapping function.
// The iterator should yield (K, V) pairs, and cb maps each pair to an output element.
// Returns nil if the result is empty.
func Iter2ToSlice[TOut any, TIn iter.Seq2[K, V], K, V any](
	items TIn,
	cb func(k K, v V) TOut,
) []TOut {
	var result []TOut
	for k, v := range items {
		result = append(result, cb(k, v))
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// Iter2ToIter2Err converts an iterator of (TIn, error) to (TOut, error), casting each item to TOut.
// If the error is non-nil, it is passed through with a zero-value TOut.
// Assumes TIn can be cast to TOut via empty interface.
func Iter2ToIter2Err[TOut any, TIn any](
	seq iter.Seq2[TIn, error],
) iter.Seq2[TOut, error] {
	return func(yield func(TOut, error) bool) {
		var null TOut
		for item, err := range seq {
			if err != nil {
				if !yield(null, err) {
					return
				}
				continue
			}
			if !yield(any(item).(TOut), nil) {
				return
			}
		}
	}
}
