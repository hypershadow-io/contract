package impl

import (
	"context"
	"iter"
	"slices"
	"sync"

	"github.com/hypershadow-io/contract/hook"
	"github.com/hypershadow-io/contract/plugin"
)

// NewMutators creates a plugin-aware hook collection for mutation hooks.
func NewMutators[T any](pc plugin.Client) Mutators[T] {
	return NewCollection[hook.MutatorFunc[T], T](pc)
}

// NewEvents creates a plugin-aware hook collection for event hooks.
func NewEvents[T any](pc plugin.Client) Events[T] { return NewCollection[hook.EventFunc[T], T](pc) }

// NewCollection creates a new plugin-aware collection of hooks,
// allowing registration and filtered retrieval of handlers by plugin ID.
func NewCollection[H any, V any](
	pc plugin.Client,
) Collection[H, V] {
	return &collection[H, V]{
		pc:       pc,
		registry: &registry[H, V]{},
	}
}

type (
	// Mutators is a plugin-aware hook collection for value-transforming hooks.
	Mutators[T any] = Collection[hook.MutatorFunc[T], T]

	// Events is a plugin-aware event collection for non-mutating lifecycle hooks.
	Events[T any] = Collection[hook.EventFunc[T], T]

	// Collection combines Registry and Provider interfaces for a given handler type.
	// It supports multi-plugin isolation and filtered access to handlers.
	Collection[H any, V any] interface {
		// Registry returns a plugin-scoped Registry for hook registration.
		Registry(pluginID string) hook.Registry[H, V]

		// Provider retrieves applicable hooks for the given context, kinds, and value.
		hook.Provider[H, V]
	}

	// collection is the internal implementation of a plugin-aware hook collection.
	collection[H any, V any] struct {
		pc       plugin.Client   // plugin client used to resolve plugin IDs
		registry *registry[H, V] // shared registry across all plugins
	}

	// registry is the core hook storage mechanism.
	// It stores hook handlers along with filters and associated plugin IDs.
	registry[H any, V any] struct {
		locker  sync.RWMutex  // protects concurrent access
		storage []entry[H, V] // list of registered hook entries
	}

	// pluginRegistry is a wrapper that provides plugin-specific access to the shared registry.
	pluginRegistry[H any, V any] struct {
		registry *registry[H, V] // underlying registry
		pluginID string          // ID of the owning plugin
	}

	// entry represents a single hook registration.
	entry[H any, V any] struct {
		// pluginID identifies the plugin that registered this hook.
		// Used for filtering and namespacing hook execution.
		pluginID string

		// filter defines the conditions under which this hook should be executed.
		// If the filter returns false, the hook will be skipped.
		filter hook.Filter[V]

		// handler is the actual hook function to be executed.
		// Its type is generic and depends on whether it's a mutator or event hook.
		handler H
	}
)

func (a *collection[H, V]) Registry(pluginID string) hook.Registry[H, V] {
	return pluginRegistry[H, V]{
		registry: a.registry,
		pluginID: pluginID,
	}
}

func (a pluginRegistry[H, V]) Add(
	filter hook.Filter[V],
	hook H,
) hook.Registry[H, V] {
	a.registry.locker.Lock()
	a.registry.storage = append(a.registry.storage, entry[H, V]{
		pluginID: a.pluginID,
		filter:   filter,
		handler:  hook,
	})
	a.registry.locker.Unlock()
	return a
}

func (a *collection[H, V]) Find(c context.Context, kinds hook.Kinds, value V) iter.Seq[H] {
	a.registry.locker.RLock()
	result := make([]H, 0, len(a.registry.storage))
	for _, h := range a.registry.storage {
		if !a.pc.IsActive(c, h.pluginID) {
			continue
		}
		if h.filter == nil || h.filter(c, kinds, value) {
			result = append(result, h.handler)
		}
	}
	a.registry.locker.RUnlock()
	return slices.Values(result)
}
