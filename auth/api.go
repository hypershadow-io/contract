package auth

import (
	"context"
	"time"

	"github.com/hypershadow-io/contract/entity"
	"github.com/hypershadow-io/contract/meta"
)

type (
	// Client defines the interface for working with authentication tokens.
	Client interface {
		// Encode generates a token string for the given customer,
		// including optional metadata and expiration time.
		Encode(
			customerID int64,
			m meta.Meta,
			expiredAt time.Time,
		) (token_ string, err_ error)

		// Decode parses and validates a token string, returning the decoded token if valid.
		Decode(token string) (res_ Token, valid_ bool, err_ error)

		// TokenFromContext extracts a Token from the given context.
		TokenFromContext(c context.Context) Token

		// TokenToContext creates a new context with the given Token embedded in it.
		TokenToContext(
			c context.Context,
			token Token,
		) context.Context
	}

	// Token represents a decoded authentication token.
	Token interface {
		// IsValid returns true if the token is valid and not expired.
		IsValid() bool

		// GetCustomerID returns the customer ID associated with the token.
		GetCustomerID() int64

		// GetMeta returns metadata attached to the token.
		GetMeta() meta.Meta

		// GetExpiredAt returns the token's expiration time.
		GetExpiredAt() time.Time
	}

	// ScopeMaker is responsible for constructing an access scope from the request context.
	ScopeMaker interface {
		// Scope extracts or builds a Scope from the given context.
		Scope(c context.Context) (Scope, error)
	}

	// Scope represents a security scope tied to a specific entity (e.g., "agent", "dispatcher").
	Scope interface {
		// EntityType returns the type of the scoped entity (e.g. "agent", "dispatcher").
		EntityType() entity.Type

		// GetEntityID returns the ID of the scoped entity.
		GetEntityID() entity.ID
	}
)
