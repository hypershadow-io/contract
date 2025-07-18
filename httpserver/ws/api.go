package ws

import (
	"context"

	"github.com/hypershadow-io/contract/httpserver"
)

type (
	// Builder constructs a WebSocket handler with the given handler logic.
	// Typically used to create a fully configured httpserver.Handler that upgrades HTTP requests to WebSocket connections.
	Builder interface {
		// Handler wraps the given WebSocket logic into an HTTP-compatible handler.
		Handler(Handle) httpserver.Handler
	}

	// Handle defines the core WebSocket handler function.
	// It receives a Ctx representing the active WebSocket connection.
	Handle = func(ctx Ctx) error

	// Ctx represents an active WebSocket connection context.
	// It embeds context.Context for cancellation and deadline support,
	// and provides methods for message exchange, metadata access, and lifecycle control.
	Ctx interface {
		context.Context

		// WriteJSON sends the given value to the client as a JSON-encoded message.
		WriteJSON(value any) error

		// ReadJSON reads the next JSON message from the client and decodes it.
		ReadJSON(value any) error

		// GetParam returns the value of a path parameter by key (e.g. from the URL route).
		GetParam(key string) string

		// GetQuery returns the value of a query parameter by key.
		GetQuery(key string) string

		// GetHeader returns the value of a request header by key.
		GetHeader(key string) string

		// Ping sends a ping frame to the client to keep the connection alive.
		Ping() error

		// Close cleanly closes the WebSocket connection.
		Close() error

		// GetIP returns the remote IP address of the connected client.
		GetIP() string

		// IsWSCloseError checks whether the given error corresponds to a normal WebSocket close event.
		IsWSCloseError(err error) bool
	}
)
