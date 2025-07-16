package cache

import (
	"context"

	"github.com/hypershadow-io/contract/eb"
)

// Get retrieves a value of type T from the cache using the given key and error builder.
// Returns the typed result, a boolean indicating whether it was found, and an error if occurred.
func Get[T any](
	c context.Context,
	instance Instance,
	errBuilder eb.LazyBuilder,
	key any,
) (res_ T, found_ bool, err_ error) {
	result, found, err := instance.Get(c, errBuilder, key)
	if err != nil {
		return res_, found, err
	}
	if !found {
		return res_, found, nil
	}
	return result.(T), found, nil
}
