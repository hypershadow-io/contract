package httpserver

import (
	"context"
)

type (
	// RouterBuilder is responsible for creating isolated Router instances per plugin.
	// Each plugin receives its own Router, which can be configured independently.
	RouterBuilder interface {
		// Router returns a new Router instance associated for the given plugin ID.
		Router(pluginID string) Router
	}

	// Router defines the interface for registering HTTP route handlers and middleware.
	// It supports route grouping, middleware chaining, and handling of standard HTTP methods.
	Router interface {
		// Builder creates a new RouterBuilder instance for this router,
		// allowing other plugins to register nested routes under this router's path.
		Builder() RouterBuilder

		// Use registers one or more middleware handlers to be executed for all routes in this router.
		Use(handlers ...Handler) Router

		// Group creates a sub-router with the specified path prefix and optional middleware.
		// Useful for namespacing routes under a common base path.
		Group(prefix string, handlers ...Handler) Router

		// Get registers a handler for HTTP GET requests on the given path.
		Get(path string, handlers ...Handler)

		// Head registers a handler for HTTP HEAD requests on the given path.
		Head(path string, handlers ...Handler)

		// Post registers a handler for HTTP POST requests on the given path.
		Post(path string, handlers ...Handler)

		// Put registers a handler for HTTP PUT requests on the given path.
		Put(path string, handlers ...Handler)

		// Delete registers a handler for HTTP DELETE requests on the given path.
		Delete(path string, handlers ...Handler)

		// Connect registers a handler for HTTP CONNECT requests on the given path.
		Connect(path string, handlers ...Handler)

		// Options registers a handler for HTTP OPTIONS requests on the given path.
		Options(path string, handlers ...Handler)

		// Trace registers a handler for HTTP TRACE requests on the given path.
		Trace(path string, handlers ...Handler)

		// Patch registers a handler for HTTP PATCH requests on the given path.
		Patch(path string, handlers ...Handler)

		// Any registers a handler for all HTTP methods on the given path.
		Any(path string, handlers ...Handler)
	}

	// Handler defines a request handler or middleware function.
	// It receives a context.Context and returns an error if the request fails.
	Handler = func(c context.Context) error
)
