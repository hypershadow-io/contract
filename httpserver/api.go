package httpserver

import (
	"context"
)

type (
	// Client defines the interface for working with the HTTP server API.
	Client interface {
		// CtxFromContext extracts a request-specific Ctx from the given context.Context.
		// This allows plugins and middleware to interact with the current HTTP request.
		CtxFromContext(c context.Context) Ctx
	}

	// Ctx defines a unified abstraction over an HTTP request/response lifecycle.
	// It provides methods for parsing input, sending responses, and accessing request metadata.
	Ctx interface {
		// IsValid returns true if the Ctx was successfully extracted from the base context
		// and is ready to be used for request/response handling.
		IsValid() bool

		// Next continues processing the request through the remaining middleware or handlers.
		// Should be called once to proceed with the request pipeline.
		Next(c context.Context) error

		// Send sends an HTTP response to the client.
		//
		// The `out` parameter determines the response body format:
		//   - string: sent as a plain UTF-8 text response.
		//   - []byte: sent as a raw binary response.
		//   - io.Reader: streamed directly to the client.
		//   - any other type: marshaled to JSON before sending.
		Send(out any) error

		// ParseParams populates the given struct with values from path parameters.
		// Returns an error if binding fails.
		ParseParams(in any) error

		// ParseQuery populates the given struct with values from URL query parameters.
		// Returns an error if binding fails.
		ParseQuery(in any) error

		// ParseBody populates the given struct with values from the request body.
		// The expected body format (JSON, form, etc.) depends on Content-Type.
		ParseBody(in any) error

		// ParseAny attempts to populate the struct from (path -> body -> query) data.
		// Useful for compact handlers where inputs may come from multiple sources.
		ParseAny(in any) error

		// GetBody returns the raw request body as a byte slice.
		GetBody() []byte

		// GetHeader retrieves the value of the specified request header key.
		GetHeader(key string) string

		// SetHeader sets the response header for the given key and value.
		SetHeader(key string, value string)

		// GetURI returns the full request URI, including path and query string.
		GetURI() string

		// GetIP returns the client's IP address as seen by the server.
		GetIP() string
	}
)
