package schema

import (
	"github.com/hypershadow-io/contract/operation/schema"
)

// Schema defines an extended schema for REST-specific usage,
// building on top of the base schema.Schema interface.
type Schema interface {
	schema.Schema

	// GetEndpoint returns the URL path used for the REST request.
	GetEndpoint() string

	// GetContentType returns the expected content type of the request payload (e.g., "application/json").
	GetContentType() string
}
