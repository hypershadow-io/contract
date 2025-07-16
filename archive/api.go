package archive

import (
	"context"
	"iter"

	"github.com/hypershadow-io/contract/db"
	"github.com/hypershadow-io/contract/entity"
)

type (
	// Builder defines an interface for constructing archive instances
	// based on the entity type and table/column metadata.
	Builder interface {
		// NewInstance creates a new archive instance for the given entity.
		NewInstance(
			model any,
			entityType entity.Type,
			entityTableName string,
			entityIDColumnName string,
		) Instance
	}

	// Instance defines the interface for interacting with archive storage.
	Instance interface {
		// AddMany inserts archived records into the archive storage based on the query.
		// Returns a sequence of inserted entities.
		AddMany(
			c context.Context,
			errBuilder func() error,
			customerID int64,
			query db.Query,
		) iter.Seq2[any, error]

		// GetMany retrieves archived entities by a list of IDs.
		GetMany(
			c context.Context,
			errBuilder func() error,
			entityIDs []int64,
		) iter.Seq2[any, error]

		// AddOne inserts a single archived record into the archive storage.
		// Returns the inserted entity, a found flag, and an error if occurred.
		AddOne(
			c context.Context,
			errBuilder func() error,
			customerID int64,
			query db.Query,
		) (res_ any, found_ bool, err_ error)

		// GetOne retrieves a single archived entity by its ID.
		// Returns the entity, a found flag, and an error if occurred.
		GetOne(
			c context.Context,
			errBuilder func() error,
			entityID int64,
		) (res_ any, found_ bool, err_ error)
	}
)
