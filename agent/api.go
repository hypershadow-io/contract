package agent

import (
	"context"

	"github.com/hypershadow-io/contract/agent/model"
	"github.com/hypershadow-io/contract/entity"
	"github.com/hypershadow-io/contract/hook"
	"github.com/hypershadow-io/contract/qb"
)

// Client defines the base interface for working with Agent.
type Client interface {
	// ModelHook returns a mutator registry for agent models associated with the given plugin ID.
	ModelHook(pluginID string) hook.Mutator[model.Model]

	// ModelEvent returns an event registry for agent models associated with the given plugin ID.
	ModelEvent(pluginID string) hook.Event[model.Model]

	// SQLSelectHook returns a mutator registry for SELECT SQL queries related to agents.
	SQLSelectHook(pluginID string) hook.Mutator[qb.SelectQuery]

	// SQLInsertHook returns a mutator registry for INSERT SQL queries related to agents.
	SQLInsertHook(pluginID string) hook.Mutator[qb.InsertQuery]

	// SQLUpdateHook returns a mutator registry for UPDATE SQL queries related to agents.
	SQLUpdateHook(pluginID string) hook.Mutator[qb.UpdateQuery]

	// SQLDeleteHook returns a mutator registry for DELETE SQL queries related to agents.
	SQLDeleteHook(pluginID string) hook.Mutator[qb.DeleteQuery]

	// IDFromContext extracts the agent ID from the given context.
	IDFromContext(c context.Context) int64

	// IDToContext creates a new context with the given agent ID embedded in it.
	IDToContext(c context.Context, agentID int64) context.Context
}

// EntityType is the global identifier for the agent entity type.
const EntityType entity.Type = "agent"
