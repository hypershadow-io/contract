package model

import (
	"github.com/hypershadow-io/contract/meta"
)

// Model defines the base interface for integration entity.
type Model interface {
	// GetID returns the unique identifier of the integration.
	GetID() int64

	// GetTitle returns the display name of the integration.
	GetTitle() string

	// GetDefinitionKey returns the key of the definition managing this integration.
	GetDefinitionKey() string

	// GetMeta returns the metadata used during runtime (editable).
	GetMeta() meta.Meta

	// GetSourceMeta returns the metadata loaded from storage (read-only snapshot).
	GetSourceMeta() meta.Meta

	// MergeMeta merges the given metadata into the integration's current metadata and returns the updated model.
	MergeMeta(meta.Meta) Model
}
