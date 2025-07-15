package byid

import (
	"context"

	"github.com/hypershadow-io/contract/agenttoken/model"
)

// Client is a specialized interface for retrieving an agent token model by its ID.
type Client interface {
	// FindByID fetches an agent token model by ID.
	// Returns the model, a boolean indicating whether it was found, and an error if occurred.
	FindByID(c context.Context, id int64) (res_ model.Model, found_ bool, err_ error)
}
