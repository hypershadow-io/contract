package ctx

import (
	"context"
)

// Client defines the interface for storing/retrieving Agent ID from context.
type Client interface {
	// IDFromContext extracts the agent ID from the given context.
	IDFromContext(c context.Context) int64

	// IDToContext creates a new context with the given agent ID embedded in it.
	IDToContext(c context.Context, agentID int64) context.Context
}
