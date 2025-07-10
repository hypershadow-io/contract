package httpserver

import (
	"context"
)

// HandleParams creates a request handler that parses path parameters into the input struct.
// It invokes the callback with the parsed data and sends the output as a response.
func HandleParams[In, Out any](
	client Client,
	cb func(c context.Context, in In) (Out, error),
) Handler {
	return makeHandler[In, Out](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseParams(in) })
}

// HandleQuery creates a request handler that parses query parameters into the input struct.
// It invokes the callback with the parsed data and sends the output as a response.
func HandleQuery[In, Out any](
	client Client,
	cb func(c context.Context, in In) (Out, error),
) Handler {
	return makeHandler[In, Out](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseQuery(in) })
}

// HandleBody creates a request handler that parses the request body into the input struct.
// It invokes the callback with the parsed data and sends the output as a response.
func HandleBody[In, Out any](
	client Client,
	cb func(c context.Context, in In) (Out, error),
) Handler {
	return makeHandler[In, Out](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseBody(in) })
}

// HandleAny creates a request handler that attempts to parse input from any source
// (path parameters, body and query string) into the input struct.
// Useful when input can come from multiple places.
// It invokes the callback with the parsed data and sends the output as a response.
func HandleAny[In, Out any](
	client Client,
	cb func(c context.Context, in In) (Out, error),
) Handler {
	return makeHandler[In, Out](client, cb, func(ctx Ctx, in *In) error { return ctx.ParseAny(in) })
}

// Handle creates a request handler that does not require any parsed input.
// It invokes the callback with the base context, and sends the output as a response.
func Handle[Out any](
	client Client,
	cb func(c context.Context) (Out, error),
) Handler {
	return func(c context.Context) error {
		out, err := cb(c)
		if err != nil {
			return err
		}
		return client.CtxFromContext(c).Send(out)
	}
}

// makeHandler is an internal helper that builds a generic request handler.
// It parses the input using the provided parser, then invokes the callback, and sends the response.
//
// Parameters:
//   - client: the context-aware HTTP client abstraction used to extract Ctx.
//   - cb: the core handler logic to execute after input parsing.
//   - parser: a function that parses input data into the expected input struct.
func makeHandler[In, Out any](
	client Client,
	cb func(c context.Context, in In) (Out, error),
	parser func(r Ctx, in *In) error,
) Handler {
	return func(c context.Context) error {
		ctx := client.CtxFromContext(c)
		var in In
		if err := parser(ctx, &in); err != nil {
			return err
		}
		out, err := cb(c, in)
		if err != nil {
			return err
		}
		return ctx.Send(out)
	}
}
