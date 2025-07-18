package codec

import (
	"time"

	"github.com/hypershadow-io/contract/auth/token"
	"github.com/hypershadow-io/contract/meta"
)

// Client defines the interface for encoding/decoding Auth tokens.
type Client interface {
	// TokenEncode generates a token string for the given customer,
	// including optional metadata and expiration time.
	TokenEncode(
		customerID int64,
		m meta.Meta,
		expiredAt time.Time,
	) (token_ string, err_ error)

	// TokenDecode parses and validates a token string, returning the decoded token if valid.
	TokenDecode(token string) (res_ token.Token, valid_ bool, err_ error)
}
