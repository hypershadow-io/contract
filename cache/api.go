package cache

import (
	"context"
	"time"
)

type (
	// Instance defines the base cache interface for storing and retrieving values.
	Instance interface {
		// Get retrieves a value from the cache by key.
		// Returns the value, a boolean indicating whether it was found, and an error if occurred.
		Get(c context.Context, errBuilder func() error, key any) (res_ any, found_ bool, err_ error)

		// Set adds a value to the cache under the specified key with optional configuration.
		Set(c context.Context, key any, value any, opts ...SetOption) error

		// Delete removes a value from the cache by key.
		Delete(c context.Context, key any) error
	}

	// SetOption defines a function used to configure cache set behavior (e.g., TTL).
	SetOption func(setOption)

	// setOption defines internal configuration methods for cache set options.
	setOption interface {
		// SetTTL sets the time-to-live (TTL) for the cached value.
		SetTTL(time.Duration)
	}
)

// WithSetTTL returns a SetOption that sets the TTL for the cached value.
func WithSetTTL(v time.Duration) SetOption {
	return func(opt setOption) { opt.SetTTL(v) }
}
