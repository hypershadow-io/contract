package fromparams

import (
	"context"
	"time"

	"github.com/hypershadow-io/contract/agenttoken/model"
	"github.com/hypershadow-io/contract/meta"
)

// Client is a specialized interface for creating new agent tokens.
type Client interface {
	// CreateFromParams creates a new agent token for the specified customer.
	// Takes individual parameters.
	// Returns the created agent token model and an error, if any.
	CreateFromParams(
		c context.Context,
		customerID int64,
		lookupKey string,
		expiredTTL time.Duration,
		m meta.Meta,
	) (res_ model.Model, err_ error)
}
