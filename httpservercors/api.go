package httpservercors

import (
	"github.com/hypershadow-io/contract/httpserver"
)

type (
	// Builder defines an interface for constructing CORS HTTP handlers based on provided configuration.
	Builder interface {
		// Handler creates a new CORS handler based on the provided configuration.
		Handler(Config) httpserver.Handler
	}

	// Config defines the configuration for the CORS middleware.
	Config struct {
		// AllowOriginsFunc defines a function that will set the 'Access-Control-Allow-Origin'
		// response header to the 'origin' request header when returned true. This allows for
		// dynamic evaluation of allowed origins. Note if AllowCredentials is true, wildcard origins
		// will be not have the 'Access-Control-Allow-Credentials' header set to 'true'.
		//
		// Optional. Default: nil
		AllowOriginsFunc func(origin string) bool

		// AllowOrigin defines a comma separated list of origins that may access the resource.
		//
		// Optional. Default value "*"
		AllowOrigins string

		// AllowMethods defines a list methods allowed when accessing the resource.
		// This is used in response to a preflight request.
		//
		// Optional. Default value "GET,POST,HEAD,PUT,DELETE,PATCH"
		AllowMethods string

		// AllowHeaders defines a list of request headers that can be used when
		// making the actual request. This is in response to a preflight request.
		//
		// Optional. Default value "".
		AllowHeaders string

		// AllowCredentials indicates whether or not the response to the request
		// can be exposed when the credentials flag is true. When used as part of
		// a response to a preflight request, this indicates whether or not the
		// actual request can be made using credentials. Note: If true, AllowOrigins
		// cannot be set to a wildcard ("*") to prevent security vulnerabilities.
		//
		// Optional. Default value false.
		AllowCredentials bool

		// ExposeHeaders defines a whitelist headers that clients are allowed to
		// access.
		//
		// Optional. Default value "".
		ExposeHeaders string

		// MaxAge indicates how long (in seconds) the results of a preflight request
		// can be cached.
		// If you pass MaxAge 0, Access-Control-Max-Age header will not be added and
		// browser will use 5 seconds by default.
		// To disable caching completely, pass MaxAge value negative. It will set the Access-Control-Max-Age header 0.
		//
		// Optional. Default value 0.
		MaxAge int
	}
)
