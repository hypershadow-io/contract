package hook

import (
	"context"
	"iter"
)

type (
	// Mutator defines a registry of mutation hooks for type T.
	// Mutators can transform a value before or after a specific action.
	Mutator[T any] = Registry[MutatorFunc[T], T]

	// MutatorFunc represents a hook function that can modify a value of type T
	// before it proceeds through the system. It receives context and hook kinds.
	MutatorFunc[T any] func(c context.Context, kinds Kinds, value T) (T, error)

	// Event defines a registry of event hooks for type T.
	// Event hooks are used to observe or react to lifecycle events without modifying data.
	Event[T any] = Registry[EventFunc[T], T]

	// EventFunc represents a non-mutating hook function for a value of type T.
	// It can be used for logging, auditing, notifications, or other side effects.
	EventFunc[T any] func(c context.Context, kinds Kinds, value T) error

	// Registry defines a hook registry that allows adding hook handlers with filters.
	// The registry may contain multiple handlers that match different kinds or values.
	Registry[H any, V any] interface {
		// Add registers a new hook handler with an associated filter.
		Add(filter Filter[V], hook H) Registry[H, V]
	}

	// Provider defines an interface for retrieving applicable hook handlers
	// based on the current context, kinds, and value.
	Provider[H any, V any] interface {
		// Find returns a sequence of handlers matching the given context, kinds, and value.
		Find(c context.Context, kinds Kinds, value V) iter.Seq[H]
	}
)
