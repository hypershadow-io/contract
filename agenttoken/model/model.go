package model

import (
	"time"

	"github.com/hypershadow-io/contract/meta"
)

// Model defines the base interface for AgentToken entity.
type Model interface {
	// GetID returns the unique identifier of the agent token.
	GetID() int64

	// GetLookupKey returns a short key that can be used to locate the token.
	GetLookupKey() string

	// GetExpiredAt returns the expiration time of the token.
	GetExpiredAt() time.Time

	// IsValid checks whether the token is still valid based on expiration and other conditions.
	IsValid() bool

	// GetMeta returns the metadata used during runtime (editable).
	GetMeta() meta.Meta

	// GetSourceMeta returns the metadata loaded from storage (read-only snapshot).
	GetSourceMeta() meta.Meta

	// MergeMeta merges the given metadata into the agent token's current metadata and returns the updated model.
	MergeMeta(meta.Meta) Model
}
