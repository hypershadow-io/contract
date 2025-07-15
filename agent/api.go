package agent

import (
	"context"

	"github.com/hypershadow-io/contract/entity"
)

// Client defines the base interface for working with Agent.
type Client interface {
	// IDFromContext extracts the agent ID from the given context.
	IDFromContext(c context.Context) int64

	// IDToContext creates a new context with the given agent ID embedded in it.
	IDToContext(c context.Context, agentID int64) context.Context
}

// EntityType is the global identifier for the agent entity type.
const EntityType entity.Type = "agent"
