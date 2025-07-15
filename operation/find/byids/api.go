package byids

import (
	"context"
	"iter"

	"github.com/hypershadow-io/contract/operation/model"
)

// Client is a specialized interface for retrieving an operation models by their IDs.
type Client interface {
	// FindByIDs fetches multiple operation models by their IDs.
	// Returns an iterator of models paired with potential errors.
	FindByIDs(c context.Context, ids []int64) iter.Seq2[model.Model, error]
}
