package hook

import "strings"

// NewKinds creates a Kinds set from a list of Kind values.
func NewKinds(list ...Kind) Kinds {
	a := make(Kinds, len(list))
	for _, kind := range list {
		a[kind] = true
	}
	return a
}

type (
	// Kind represents a type or phase of a hookable event.
	Kind string

	// Kinds is a set of Kind values used to define which events a hook applies to.
	Kinds map[Kind]bool
)

// Base hook kinds commonly used across modules.
const (
	KindBefore Kind = "Before"
	KindAfter  Kind = "After"

	KindFind   Kind = "Find"
	KindCreate Kind = "Create"
	KindUpdate Kind = "Update"
	KindDelete Kind = "Delete"

	KindLock Kind = "Lock"
	KindOne  Kind = "One"
	KindMany Kind = "Many"
	KindByID Kind = "ByID"

	KindUI     Kind = "UI"
	KindSystem Kind = "System"
	KindSilent Kind = "Silent"
)

// With returns a new Kinds set containing the current kinds plus the provided ones.
func (a Kinds) With(list ...Kind) Kinds {
	result := make(Kinds, len(a)+len(list))
	for item := range a {
		result[item] = true
	}
	for _, item := range list {
		result[item] = true
	}
	return result
}

// Pick returns a new Kinds set containing only the kinds from the input list that are present in the original set.
func (a Kinds) Pick(list ...Kind) Kinds {
	result := make(Kinds, len(list))
	for _, item := range list {
		if a[item] {
			result[item] = true
		}
	}
	return result
}

// Has checks whether the given kind is present in the set.
func (a Kinds) Has(item Kind) bool {
	return a[item]
}

// Not returns true if the given kind is not present in the set.
func (a Kinds) Not(item Kind) bool {
	return !a[item]
}

// HasAll checks if all provided kinds are present in the set.
func (a Kinds) HasAll(list ...Kind) bool {
	for _, item := range list {
		if !a[item] {
			return false
		}
	}
	return true
}

// HasAny checks if at least one of the provided kinds is present in the set.
func (a Kinds) HasAny(list ...Kind) bool {
	for _, item := range list {
		if a[item] {
			return true
		}
	}
	return false
}

// HasOnly checks if the set contains only the specified kinds (no more, no less).
func (a Kinds) HasOnly(list ...Kind) bool {
	if len(a) != len(list) {
		return false
	}
	for _, item := range list {
		if !a[item] {
			return false
		}
	}
	return true
}

// String returns the Kind as a string.
func (a Kind) String() string {
	return string(a)
}

// String returns a comma-separated list of kind names in the set.
func (a Kinds) String() string {
	var builder strings.Builder
	for item := range a {
		builder.WriteRune(',')
		builder.WriteString(item.String())
	}
	if builder.Len() > 0 {
		return builder.String()[1:]
	}
	return ""
}
