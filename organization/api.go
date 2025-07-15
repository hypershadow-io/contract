package organization

import (
	"context"

	"github.com/hypershadow-io/contract/db"
	"github.com/hypershadow-io/contract/entity"
)

// Client defines the base interface for working with Organization.
type Client interface {
	// IDFromContext extracts the organization ID from the given context.
	IDFromContext(c context.Context) int64

	// IDToContext creates a new context with the given organization ID embedded in it.
	IDToContext(c context.Context, organizationID int64) context.Context

	// DB retrieves or creates a database instance associated with the specified organization ID.
	// Returns the instance, a boolean indicating whether it was found, and an error if occurred.
	DB(c context.Context, organizationID int64) (res_ db.Instance, found_ bool, err_ error)
}

// EntityType is the global identifier for the organization entity type.
const EntityType entity.Type = "organization"
