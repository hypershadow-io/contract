package schema

import (
	"github.com/hypershadow-io/contract/operation/schema"
)

// Schema defines a high-level interface for representing a integration schema definition.
type Schema interface {
	// GetTitle returns the human-readable title of the schema.
	GetTitle() string

	// GetVersion returns the version string of the schema (e.g., "v1.0").
	GetVersion() string

	// GetOperations returns a list of operation schemas defined within this schema version.
	GetOperations() schema.List
}
