package hook

import (
	"github.com/hypershadow-io/contract/agenttoken/model"
	"github.com/hypershadow-io/contract/hook"
	"github.com/hypershadow-io/contract/qb"
)

// Client defines the hook-aware client interface for working with AgentToken logic.
// Provides access to model-level and SQL-level mutators and event hooks.
type Client interface {
	// ModelHook returns a mutator registry for agent token models associated with the given plugin ID.
	ModelHook(pluginID string) hook.Mutator[model.Model]

	// ModelEvent returns an event registry for agent token models associated with the given plugin ID.
	ModelEvent(pluginID string) hook.Event[model.Model]

	// SQLSelectHook returns a mutator registry for SELECT SQL queries associated with the given plugin ID.
	SQLSelectHook(pluginID string) hook.Mutator[qb.SelectQuery]

	// SQLInsertHook returns a mutator registry for INSERT SQL queries associated with the given plugin ID.
	SQLInsertHook(pluginID string) hook.Mutator[qb.InsertQuery]

	// SQLDeleteHook returns a mutator registry for DELETE SQL queries associated with the given plugin ID.
	SQLDeleteHook(pluginID string) hook.Mutator[qb.DeleteQuery]
}
