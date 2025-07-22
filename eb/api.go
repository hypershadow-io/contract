package eb

import (
	"context"

	"github.com/hypershadow-io/contract/meta"
)

type (
	// Builder defines a pluggable and extensible error builder interface.
	// All plugins should use this interface for consistent error handling.
	Builder interface {
		// Error returns the final error message as a string (implements error).
		Error() string

		// GetMessage returns the error message.
		// If not explicitly set on the builder, attempts to retrieve it from wrapped errors.
		GetMessage() string

		// SetMessagef sets a formatted error message.
		SetMessagef(format string, args ...any) Builder

		// GetKey returns the machine-readable error key.
		// This key is included in the client response and can be used to correlate the error
		// with extended logs or internal diagnostics.
		// If not explicitly set on the builder, attempts to retrieve it from wrapped errors.
		GetKey() string

		// SetKey sets the error key that will be returned to the client.
		// It should uniquely identify the error for further log lookup or debugging.
		SetKey(key string) Builder

		// GetValidation returns field-level validation errors.
		// If not explicitly set on the builder, attempts to retrieve it from wrapped errors.
		GetValidation() map[string]string

		// SetValidation sets field-level validation errors.
		SetValidation(err map[string]string) Builder

		// GetCode returns the optional numeric code (e.g., HTTP status).
		// If not explicitly set on the builder, attempts to retrieve it from wrapped errors.
		GetCode() int

		// SetCode sets the numeric code.
		SetCode(code int) Builder

		// AddWrap adds another error to the wrap stack (for traceability).
		AddWrap(err error) Builder

		// GetLogMessage returns the log-specific message.
		// If not explicitly set on the builder, attempts to retrieve it from wrapped errors.
		GetLogMessage() string

		// SetLogMessagef sets a formatted log message for internal logging purposes.
		SetLogMessagef(format string, args ...any) Builder

		// GetMeta returns associated metadata.
		// If not explicitly set on the builder, attempts to retrieve it from wrapped errors.
		GetMeta() meta.Meta

		// SetMeta replaces the metadata.
		SetMeta(m meta.Meta) Builder

		// MergeMeta merges additional metadata into the existing set.
		MergeMeta(m meta.Meta) Builder

		// DrainMeta collects all metadata from the builder and any wrapped errors.
		// It removes the metadata from each source as it is collected, ensuring no duplication.
		DrainMeta() meta.Meta

		// GetLogger returns the associated logging function (if any).
		// If not explicitly set on the builder, attempts to retrieve it from wrapped errors.
		GetLogger() LogFunc

		// SetLogger sets a logging function.
		SetLogger(logger LogFunc) Builder

		// SetNoLogger disables logging for this error.
		SetNoLogger() Builder
	}

	// LazyBuilder is a deferred factory function that produces a new Builder instance.
	LazyBuilder = func() Builder

	// LogFunc defines the signature for logging error-related messages with context.
	LogFunc = func(c context.Context, msg string, args ...any)
)
