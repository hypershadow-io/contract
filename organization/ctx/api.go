package ctx

import (
	"context"
)

// Client defines the interface for storing/retrieving Organization ID from context.
type Client interface {
	// IDFromContext extracts the Organization ID from the given context.
	IDFromContext(c context.Context) int64

	// IDToContext creates a new context with the given Organization ID embedded in it.
	IDToContext(c context.Context, organizationID int64) context.Context
}
