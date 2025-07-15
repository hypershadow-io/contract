package bylookupkey

import (
	"context"

	"github.com/hypershadow-io/contract/agenttoken/model"
)

// Client is a specialized interface for retrieving an agent token model by its LookupKey.
type Client interface {
	// FindByLookupKey fetches an agent token model by LookupKey.
	// Returns the model, a boolean indicating whether it was found, and an error if occurred.
	FindByLookupKey(c context.Context, id int64) (res_ model.Model, found_ bool, err_ error)
}
