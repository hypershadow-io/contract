package crypt

import (
	"github.com/hypershadow-io/contract/codec"
)

type (
	// Builder constructs a cryptographic Client with a specific serialization strategy.
	Builder interface {
		// NewClientWithCodec returns a new Client that uses the provided codec for (de)serialization
		// of input/output data during encryption and decryption.
		NewClientWithCodec(codecClient codec.Client) Client
	}
	// Client defines the interface for encryption and decryption operations.
	// It supports context-specific encryption using organization-scoped keys and optional salts.
	Client interface {
		// Encrypt serializes and encrypts the given data using the provided organization ID and module key.
		// An optional salt can be supplied to strengthen key uniqueness.
		//
		// Returns an encrypted, base64-encoded string.
		Encrypt(
			data any,
			organizationID int64,
			moduleKey string,
			salt ...string,
		) (string, error)

		// Decrypt decodes and decrypts the given encrypted string, and deserializes it into result.
		// Requires the same organization ID, module key, and optional salt used during encryption.
		Decrypt(
			encoded string,
			result any,
			organizationID int64,
			moduleKey string,
			salt ...string,
		) error
	}

	// NoopClient is a drop-in replacement for Client that performs no encryption or decryption.
	// Useful for testing or plugins where cryptography is disabled.
	NoopClient Client
)
