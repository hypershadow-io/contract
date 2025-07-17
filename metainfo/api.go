package metainfo

import (
	"github.com/hypershadow-io/contract/field"
	"github.com/hypershadow-io/contract/identity"
)

type (
	// Client defines the interface for registering metadata information about entities.
	Client interface {
		// Registry registers an entity's metadata for the given plugin.
		Registry(pluginID string, entity Entity)
	}

	// Entity describes metadata information about a business entity.
	Entity interface {
		// Identification provides key, name, and description of the entity.
		identity.Identification

		// GetGroup returns the identification of the group this entity belongs to.
		GetGroup() identity.Identification

		// GetIntegratesKeys returns the list of integration definition keys related to this entity.
		// This defines the context or category in which the entity is applicable or allowed to operate.
		GetIntegratesKeys() []string

		// GetFields returns the list of fields associated with the entity.
		GetFields() []field.Field
	}
)
