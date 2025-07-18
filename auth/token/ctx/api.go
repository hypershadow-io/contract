package ctx

import (
	"context"

	"github.com/hypershadow-io/contract/auth/token"
)

// Client defines the interface for storing/retrieving Auth tokens from context.
type Client interface {
	// TokenFromContext extracts a Token from the given context.
	TokenFromContext(c context.Context) token.Token

	// TokenToContext creates a new context with the given Token embedded in it.
	TokenToContext(
		c context.Context,
		token token.Token,
	) context.Context
}
