package token

import (
	"time"

	"github.com/hypershadow-io/contract/meta"
)

// Token represents a decoded authentication token.
type Token interface {
	// IsValid returns true if the token is valid and not expired.
	IsValid() bool

	// GetCustomerID returns the customer ID associated with the token.
	GetCustomerID() int64

	// GetMeta returns metadata attached to the token.
	GetMeta() meta.Meta

	// GetExpiredAt returns the token's expiration time.
	GetExpiredAt() time.Time
}
