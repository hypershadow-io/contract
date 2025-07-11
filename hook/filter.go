package hook

import "context"

// AndFilters combines multiple filters using logical AND.
// Returns true only if all filters return true.
func AndFilters[T any](list ...Filter[T]) Filter[T] {
	return func(ctx context.Context, kinds Kinds, model T) bool {
		for _, filter := range list {
			if !filter(ctx, kinds, model) {
				return false
			}
		}
		return true
	}
}

// OrFilters combines multiple filters using logical OR.
// Returns true if at least one filter returns true.
func OrFilters[T any](list ...Filter[T]) Filter[T] {
	return func(ctx context.Context, kinds Kinds, model T) bool {
		for _, filter := range list {
			if filter(ctx, kinds, model) {
				return true
			}
		}
		return false
	}
}

// MatchAny returns a filter that always evaluates to true.
func MatchAny[T any]() Filter[T] {
	return func(_ context.Context, kinds Kinds, _ T) bool { return true }
}

// MatchKind returns a filter that matches only if the given kind is present.
func MatchKind[T any](kind Kind) Filter[T] {
	return func(_ context.Context, kinds Kinds, _ T) bool { return kinds.Has(kind) }
}

// ExcludeKind returns a filter that matches only if the given kind is NOT present.
func ExcludeKind[T any](kind Kind) Filter[T] {
	return func(_ context.Context, kinds Kinds, _ T) bool { return kinds.Not(kind) }
}

// MatchAllKinds returns a filter that matches only if all specified kinds are present.
func MatchAllKinds[T any](list ...Kind) Filter[T] {
	return func(_ context.Context, kinds Kinds, _ T) bool { return kinds.HasAll(list...) }
}

// MatchAnyKinds returns a filter that matches if at least one of the specified kinds is present.
func MatchAnyKinds[T any](list ...Kind) Filter[T] {
	return func(_ context.Context, kinds Kinds, _ T) bool { return kinds.HasAny(list...) }
}

// MatchOnlyKinds returns a filter that matches only if the kinds exactly match the provided list.
func MatchOnlyKinds[T any](list ...Kind) Filter[T] {
	return func(_ context.Context, kinds Kinds, _ T) bool { return kinds.HasOnly(list...) }
}

// Filter defines a predicate used to determine whether a hook should apply
// based on the current context, hook kinds, and target value.
type Filter[T any] func(c context.Context, kinds Kinds, value T) bool
