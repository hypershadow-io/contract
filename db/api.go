package db

import (
	"context"
	"io/fs"
	"iter"

	"github.com/hypershadow-io/contract/eb"
)

type (
	// Builder defines constructors and accessors for platform and organization-specific DB instances.
	// Also provides factory methods for basic and pooled connection setups.
	Builder interface {
		// NewSimpleInstance returns a direct, non-pooled database instance based on the given URI.
		NewSimpleInstance(c context.Context, uri string) (Instance, error)

		// NewPoolInstance returns a connection-pooled database instance.
		NewPoolInstance(c context.Context, uri string) (Instance, error)

		// GetPlatform retrieves the platform DB instance from context.
		GetPlatform(c context.Context) Instance

		// SetPlatform stores the platform DB instance into context.
		SetPlatform(c context.Context, instance Instance) context.Context

		// GetOrganization retrieves the organization DB instance from context.
		GetOrganization(c context.Context) Instance

		// SetOrganization stores the organization DB instance into context.
		SetOrganization(c context.Context, instance Instance) context.Context
	}

	// MigrationClient allows plugins to register and retrieve migration sources (FS).
	MigrationClient interface {
		// Add registers a migration source (filesystem) for the given plugin.
		Add(pluginID string, fs FS)

		// GetSystemOnly returns migrations registered by system-level plugins only.
		GetSystemOnly() []FS

		// Get returns all registered migrations from all plugins, including tenant-specific ones.
		Get(c context.Context) []FS
	}

	// Instance abstracts a database connection or transaction.
	// Supports basic operations, transaction lifecycle, and generic querying.
	Instance interface {
		// IsValid returns true if the instance is initialized and usable.
		IsValid() bool

		// MigrateUp applies the given migration sources to the database.
		MigrateUp(c context.Context, fs ...FS) error

		// Begin starts a new transaction and returns a new context containing it.
		Begin(c context.Context) (context.Context, error)

		// Rollback rolls back the current transaction in context.
		Rollback(c context.Context) error

		// Commit commits the current transaction in context.
		Commit(c context.Context) error

		// Detach returns a new context with the transaction removed (e.g., for sub-contexts).
		Detach(c context.Context) context.Context

		// Close shuts down the instance and releases resources.
		Close(c context.Context) error

		// Exec executes a generic query and returns the result with affected rows.
		Exec(
			c context.Context,
			query Query,
		) (ExecResult, error)

		// FindOne executes a SELECT query and decodes the first row into proto.
		// Returns found = false if no rows match.
		FindOne(
			c context.Context,
			errBuilder eb.LazyBuilder,
			proto any,
			query Query,
		) (res_ any, found_ bool, err_ error)

		// FindIterator executes a SELECT query and returns an iterator of results.
		FindIterator(
			c context.Context,
			errBuilder eb.LazyBuilder,
			proto any,
			query Query,
		) iter.Seq2[any, error]
	}

	// ExecResult represents the result of a write operation (INSERT/UPDATE/DELETE).
	ExecResult interface {
		// RowsAffected returns the number of rows modified by the operation.
		RowsAffected() int64
	}

	// Query defines an interface for objects that can be translated into SQL + arguments.
	Query interface {
		// ToSql generates the final SQL query string and argument list.
		ToSql() (sql_ string, args_ []any, err_ error)
	}

	// FS represents a virtual filesystem used for loading SQL migration files.
	// Implements both ReadDirFS and ReadFileFS.
	FS interface {
		fs.ReadDirFS
		fs.ReadFileFS
	}
)
