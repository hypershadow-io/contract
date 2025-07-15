package hook

import (
	"github.com/hypershadow-io/contract/hook"
	"github.com/hypershadow-io/contract/operation/model"
	"github.com/hypershadow-io/contract/qb"
)

// Client defines the hook-aware client interface for working with Operation logic.
// Provides access to model-level and SQL-level mutators and event hooks.
type Client interface {
	// ModelHook returns a mutator registry for operation models associated with the given plugin ID.
	ModelHook(pluginID string) hook.Mutator[model.Model]

	// ModelEvent returns an event registry for operation models associated with the given plugin ID.
	ModelEvent(pluginID string) hook.Event[model.Model]

	// SQLSelectHook returns a mutator registry for SELECT SQL queries associated with the given plugin ID.
	SQLSelectHook(pluginID string) hook.Mutator[qb.SelectQuery]

	// SQLInsertHook returns a mutator registry for INSERT SQL queries associated with the given plugin ID.
	SQLInsertHook(pluginID string) hook.Mutator[qb.InsertQuery]

	// SQLUpdateHook returns a mutator registry for UPDATE SQL queries associated with the given plugin ID.
	SQLUpdateHook(pluginID string) hook.Mutator[qb.UpdateQuery]

	// SQLDeleteHook returns a mutator registry for DELETE SQL queries associated with the given plugin ID.
	SQLDeleteHook(pluginID string) hook.Mutator[qb.DeleteQuery]
}
