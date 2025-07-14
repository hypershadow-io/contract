package agentmodel

import (
	"github.com/hypershadow-io/contract/meta"
)

type (
	// Model defines the base interface for Agent entities.
	Model interface {
		// GetID returns the unique identifier of the agent.
		GetID() int64

		// GetTitle returns the display title of the agent.
		GetTitle() string

		// GetStatus returns the current status of the agent.
		GetStatus() Status

		// GetMeta returns the metadata used during runtime (editable).
		GetMeta() meta.Meta

		// GetSourceMeta returns the metadata loaded from storage (read-only snapshot).
		GetSourceMeta() meta.Meta

		// MergeMeta merges the given metadata into the agent's current metadata and returns the updated model.
		MergeMeta(meta.Meta) Model
	}

	// Status defines the agent's current lifecycle state.
	Status int
)

const (
	StatusDraft  Status = 0 // Draft — the agent is not yet published or finalized.
	StatusActive Status = 1 // Active — the agent is ready and operational.
)
