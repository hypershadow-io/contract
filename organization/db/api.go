package db

import (
	"context"

	"github.com/hypershadow-io/contract/db"
)

// Client defines the interface for working with Organization DB.
type Client interface {
	// DB retrieves or creates a database instance associated with the specified organization ID.
	// Returns the instance, a boolean indicating whether it was found, and an error if occurred.
	DB(c context.Context, organizationID int64) (res_ db.Instance, found_ bool, err_ error)
}
