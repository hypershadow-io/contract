package model

import (
	"time"

	"github.com/hypershadow-io/contract/meta"
)

// Model defines the base interface for API token entities.
type Model interface {
	// GetID returns the unique identifier of the API token.
	GetID() int64

	// GetTitle returns the display title of the API token.
	GetTitle() string

	// GetExpiredAt returns the API token's expiration time.
	GetExpiredAt() time.Time

	// IsValid returns true if the API token is valid and not expired.
	IsValid() bool

	// GetMeta returns the metadata used during runtime (editable).
	GetMeta() meta.Meta

	// GetSourceMeta returns the metadata loaded from storage (read-only snapshot).
	GetSourceMeta() meta.Meta

	// MergeMeta merges the given metadata into the API token's current metadata and returns the updated model.
	MergeMeta(meta.Meta) Model
}
