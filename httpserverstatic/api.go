package httpserverstatic

import (
	"io/fs"

	"github.com/hypershadow-io/contract/httpserver"
)

type (
	// Builder constructs an HTTP handler for serving static files.
	// Typically used to create plugin-specific or route-specific static file handlers.
	Builder interface {
		// Handler creates a new handler for serving static content
		// based on the provided configuration.
		Handler(Config) httpserver.Handler
	}

	// Config defines the configuration for the HTTP static file handler.
	Config struct {
		// Root is a FileSystem that provides access
		// to a collection of files and directories.
		Root fs.FS

		// PathPrefix defines a prefix to be added to a filepath when
		// reading a file from the Root.
		PathPrefix string

		// Index file for serving a directory.
		// Default: "index.html"
		Index string

		// The value for the Cache-Control HTTP-header
		// that is set on the file response. MaxAge is defined in seconds.
		MaxAgeSecond int

		// Additional headers for all static in Root.
		Headers map[string]string
	}
)
