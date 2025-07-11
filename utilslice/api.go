package utilslice

// SliceToSlice transforms a slice of In items into a slice of Out
// by applying the provided mapping function to each item.
func SliceToSlice[Out, In any](
	items []In,
	cb func(item In) Out,
) []Out {
	if len(items) == 0 {
		return nil
	}
	result := make([]Out, 0, len(items))
	for _, item := range items {
		result = append(result, cb(item))
	}
	return result
}

// SliceToMap transforms a slice of In items into a map[Key]Out
// using the provided callback that returns a key-value pair for each item.
func SliceToMap[Key comparable, Out any, In any](
	items []In,
	cb func(item In) (Key, Out),
) map[Key]Out {
	if len(items) == 0 {
		return nil
	}
	result := make(map[Key]Out, len(items))
	for _, item := range items {
		k, v := cb(item)
		result[k] = v
	}
	return result
}

// SliceWithIndexToMap transforms a slice of In items into a map[Key]Out,
// using both the index and item value when generating keys and values.
func SliceWithIndexToMap[Key comparable, Out any, In any](
	items []In,
	cb func(index int, item In) (Key, Out),
) map[Key]Out {
	if len(items) == 0 {
		return nil
	}
	result := make(map[Key]Out, len(items))
	for i, item := range items {
		k, v := cb(i, item)
		result[k] = v
	}
	return result
}

// SliceToSliceIfOk maps each item in the input slice to an output item
// using the callback, and includes only those for which the second return value is true.
func SliceToSliceIfOk[Out, In any](
	items []In,
	cb func(item In) (Out, bool),
) []Out {
	if len(items) == 0 {
		return nil
	}
	result := make([]Out, 0, len(items))
	for _, item := range items {
		if newItem, ok := cb(item); ok {
			result = append(result, newItem)
		}
	}
	return result
}

// SliceToSliceOrError maps each item in the input slice to an output item.
// If any mapping returns an error, the function aborts and returns it.
func SliceToSliceOrError[Out, In any](
	items []In,
	cb func(item In) (Out, error),
) ([]Out, error) {
	if len(items) == 0 {
		return nil, nil
	}
	result := make([]Out, 0, len(items))
	for _, item := range items {
		newItem, err := cb(item)
		if err != nil {
			return nil, err
		}
		result = append(result, newItem)
	}
	return result, nil
}
