package httpauth

import (
	"context"

	"github.com/hypershadow-io/contract/auth"
	"github.com/hypershadow-io/contract/entity"
	"github.com/hypershadow-io/contract/httpserver"
)

// ScopeTypeParams returns a auth.ScopeMaker that builds a static entity type
// from path parameters.
func ScopeTypeParams[In any](
	serverClient httpserver.Client,
	entityType entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeTypeBuilder(
		serverClient,
		entityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseParams(in) },
	)
}

// ScopeParams returns a auth.ScopeMaker that dynamically resolves entity type and ID
// from path parameters.
func ScopeParams[In any](
	serverClient httpserver.Client,
	getEntityType func(in In) entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeBuilder(
		serverClient,
		getEntityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseParams(in) },
	)
}

// ScopeTypeQuery returns a auth.ScopeMaker that uses a static entity type
// and reads the entity ID from query parameters.
func ScopeTypeQuery[In any](
	serverClient httpserver.Client,
	entityType entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeTypeBuilder(
		serverClient,
		entityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseQuery(in) },
	)
}

// ScopeQuery returns a auth.ScopeMaker that dynamically resolves entity type and ID
// from query parameters.
func ScopeQuery[In any](
	serverClient httpserver.Client,
	getEntityType func(in In) entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeBuilder(
		serverClient,
		getEntityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseQuery(in) },
	)
}

// ScopeTypeBody returns a auth.ScopeMaker that uses a static entity type
// and reads the entity ID from request body.
func ScopeTypeBody[In any](
	serverClient httpserver.Client,
	entityType entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeTypeBuilder(
		serverClient,
		entityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseBody(in) },
	)
}

// ScopeBody returns a auth.ScopeMaker that dynamically resolves entity type and ID
// from request body.
func ScopeBody[In any](
	serverClient httpserver.Client,
	getEntityType func(in In) entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeBuilder(
		serverClient,
		getEntityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseBody(in) },
	)
}

// ScopeTypeAny returns a auth.ScopeMaker that uses a static entity type
// and reads the entity ID from (path -> body -> query) data.
func ScopeTypeAny[In any](
	serverClient httpserver.Client,
	entityType entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeTypeBuilder(
		serverClient,
		entityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseAny(in) },
	)
}

// ScopeAny returns a auth.ScopeMaker that dynamically resolves entity type and ID
// from (path -> body -> query) data.
func ScopeAny[In any](
	serverClient httpserver.Client,
	getEntityType func(in In) entity.Type,
	getEntityID func(in In) entity.ID,
) auth.ScopeMaker {
	return makeScopeBuilder(
		serverClient,
		getEntityType,
		getEntityID,
		func(ctx httpserver.Ctx, in *In) error { return ctx.ParseAny(in) },
	)
}

// ScopeContext returns a auth.ScopeMaker that uses a static entity type
// and receiving the entity ID from callback by context.
func ScopeContext(
	entityType entity.Type,
	getEntityID func(c context.Context) (entity.ID, error),
) auth.ScopeMaker {
	return scopeBuilder{
		builder: func(c context.Context) (auth.Scope, error) {
			entityID, err := getEntityID(c)
			if err != nil {
				return nil, err
			}
			return scope{
				entityType: entityType,
				entityID:   entityID,
			}, nil
		},
	}
}

// makeScopeTypeBuilder is a generic helper that builds scopes with a static entity type.
func makeScopeTypeBuilder[In any](
	serverClient httpserver.Client,
	entityType entity.Type,
	getEntityID func(in In) entity.ID,
	parser func(r httpserver.Ctx, in *In) error,
) auth.ScopeMaker {
	return scopeBuilder{
		builder: func(c context.Context) (auth.Scope, error) {
			var in In
			if err := parser(serverClient.CtxFromContext(c), &in); err != nil {
				return nil, err
			}
			return scope{
				entityType: entityType,
				entityID:   getEntityID(in),
			}, nil
		},
	}
}

// makeScopeBuilder is a generic helper that builds scopes with dynamic entity type and ID.
func makeScopeBuilder[In any](
	serverClient httpserver.Client,
	getEntityType func(in In) entity.Type,
	getEntityID func(in In) entity.ID,
	parser func(r httpserver.Ctx, in *In) error,
) auth.ScopeMaker {
	return scopeBuilder{
		builder: func(c context.Context) (auth.Scope, error) {
			var in In
			if err := parser(serverClient.CtxFromContext(c), &in); err != nil {
				return nil, err
			}
			return scope{
				entityType: getEntityType(in),
				entityID:   getEntityID(in),
			}, nil
		},
	}
}

type (
	// scopeBuilder is an implementation of auth.ScopeMaker using a context-based builder function.
	scopeBuilder struct {
		builder func(c context.Context) (auth.Scope, error)
	}

	// scope is a simple implementation of auth.Scope.
	scope struct {
		entityType entity.Type
		entityID   entity.ID
	}
)

func (a scopeBuilder) Scope(c context.Context) (auth.Scope, error) { return a.builder(c) }

func (a scope) EntityType() entity.Type { return a.entityType }
func (a scope) GetEntityID() entity.ID  { return a.entityID }
