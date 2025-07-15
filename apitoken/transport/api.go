package transport

import (
	"github.com/hypershadow-io/contract/apitoken/model"
)

type (
	// Client defines the interface for encoding and decoding API token models
	// to/from a network-safe transport representation (e.g., for passing via HTTPS).
	Client interface {
		// Encode encodes an API token model into a transportable string,
		// typically for sending over a network.
		Encode(
			organizationID int64,
			model model.Model,
		) (token_ string, err_ error)

		// Decode decodes a previously encoded token string back into a Transport object.
		// Returns validity flag and error, if any.
		Decode(token string) (res_ Model, valid_ bool, err_ error)
	}

	// Model defines the minimal, transport-safe representation of an API token model,
	// suitable for use in network protocols.
	Model interface {
		// GetTokenID returns the unique identifier of the API token.
		GetTokenID() int64

		// GetOrganizationID returns the ID of the organization to which the token belongs.
		GetOrganizationID() int64
	}
)
