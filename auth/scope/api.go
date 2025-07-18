package scope

import (
	"context"

	"github.com/hypershadow-io/contract/entity"
)

type (
	// Maker is responsible for constructing an access scope from the request context.
	Maker interface {
		// Scope extracts or builds a Scope from the given context.
		Scope(c context.Context) (Scope, error)
	}

	// Scope represents a security scope tied to a specific entity (e.g., "agent", "dispatcher").
	Scope interface {
		// EntityType returns the type of the scoped entity (e.g. "agent", "dispatcher").
		EntityType() entity.Type

		// GetEntityID returns the ID of the scoped entity.
		GetEntityID() entity.ID
	}
)
