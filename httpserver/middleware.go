package httpserver

import (
	"context"
)

// MiddlewareParams creates a middleware handler that parses path parameters into the input struct.
// It invokes the callback with the parsed data.
func MiddlewareParams[In any](
	client Client,
	cb func(c context.Context, in In) error,
) Handler {
	return makeMiddleware[In](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseParams(in) })
}

// MiddlewareQuery creates a middleware handler that parses URL query parameters into the input struct.
// It invokes the callback with the parsed data.
func MiddlewareQuery[In any](
	client Client,
	cb func(c context.Context, in In) error,
) Handler {
	return makeMiddleware[In](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseQuery(in) })
}

// MiddlewareBody creates a middleware handler that parses the request body into the input struct.
// It invokes the callback with the parsed data.
func MiddlewareBody[In any](
	client Client,
	cb func(c context.Context, in In) error,
) Handler {
	return makeMiddleware[In](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseBody(in) })
}

// MiddlewareAny creates a middleware handler that attempts to parse input from any source
// (path parameters, body and query string) into the input struct.
// Useful when input can come from multiple places.
// It invokes the callback with the parsed data.
func MiddlewareAny[In any](
	client Client,
	cb func(c context.Context, in In) error,
) Handler {
	return makeMiddleware[In](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseAny(in) })
}

// makeMiddleware is an internal helper that builds a generic request handler.
// It parses the input using the provided parser, then invokes the middleware callback
// with the parsed input. If parsing fails, the request is terminated with an error.
//
// Parameters:
//   - client: the context-aware HTTP client abstraction used to extract Ctx.
//   - cb: the core handler logic to execute after input parsing.
//   - parser: a function that parses input data into the expected input struct.
func makeMiddleware[In any](
	client Client,
	cb func(c context.Context, in In) error,
	parser func(ctx Ctx, in *In) error,
) Handler {
	return func(c context.Context) error {
		var in In
		if err := parser(client.CtxFromContext(c), &in); err != nil {
			return err
		}
		return cb(c, in)
	}
}
