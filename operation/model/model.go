package model

import (
	"time"

	"github.com/hypershadow-io/contract/meta"
)

// Model defines the base interface for operation entity.
type Model interface {
	// GetID returns the unique identifier of the operation.
	GetID() int64

	// GetIntegrationID returns the ID of the integration this operation belongs to.
	GetIntegrationID() int64

	// GetTitle returns the display name of the operation.
	GetTitle() string

	// GetAction returns the operation type (e.g., GET, POST, PUT for HTTP; SELECT, UPDATE for SQL).
	GetAction() string

	// GetExternalID returns the external identifier of the operation from the customer's system.
	GetExternalID() string

	// GetDispatcherKey returns the dispatcher key responsible for executing the operation request.
	GetDispatcherKey() string

	// IsLocked indicates whether the operation is protected from deletion.
	IsLocked() bool

	// GetModifiedAt returns the time of the last modification.
	GetModifiedAt() time.Time

	// GetMeta returns the metadata used during runtime (editable).
	GetMeta() meta.Meta

	// GetSourceMeta returns the metadata loaded from storage (read-only snapshot).
	GetSourceMeta() meta.Meta

	// MergeMeta merges the given metadata into the operation's current metadata and returns the updated model.
	MergeMeta(meta.Meta) Model
}
