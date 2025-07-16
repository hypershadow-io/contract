package local

import (
	"time"

	"github.com/hypershadow-io/contract/cache"
)

type (
	// Builder defines the interface for creating a new local in-memory cache instance.
	Builder interface {
		// NewInstance creates a new in-memory cache with optional configuration options.
		NewInstance(ttl time.Duration, opts ...NewInstanceOption) (cache.Instance, error)
	}

	// NewInstanceOption represents a functional option used to configure a local cache instance during creation.
	NewInstanceOption func(newInstanceOption) error

	// newInstanceOption defines methods for setting advanced cache behaviors.
	newInstanceOption interface {
		// SetCleanupTTL sets the time interval for automatic cleanup of expired entries.
		SetCleanupTTL(time.Duration) error

		// SetOnDelete registers a callback function that is invoked when an entry is deleted from the cache.
		SetOnDelete(func(any)) error
	}
)

// WithNewInstanceCleanupTTL configures the cache to periodically clean up expired entries.
func WithNewInstanceCleanupTTL(v time.Duration) NewInstanceOption {
	return func(opt newInstanceOption) error { return opt.SetCleanupTTL(v) }
}

// WithNewInstanceOnDelete configures a callback that is triggered on entry deletion.
func WithNewInstanceOnDelete(v func(any)) NewInstanceOption {
	return func(opt newInstanceOption) error { return opt.SetOnDelete(v) }
}
