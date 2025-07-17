package choice

// Selector represents a generic selection interface that allows querying selected value and options.
type Selector[T comparable] interface {
	// IsSelected returns true if a value is selected (i.e., not the zero value).
	IsSelected() bool

	// GetSelected returns the currently selected value.
	GetSelected() T

	// GetOptions returns all available selectable options.
	GetOptions() []T
}

// Make creates a new Model with the given value selected and added as the only available option.
func Make[T comparable](value T) Model[T] {
	return Model[T]{
		Selected: value,
		Options:  []T{value},
	}
}

// MakeFrom creates a new Model based on the state of the given Selector.
func MakeFrom[T comparable](in Selector[T]) Model[T] {
	return Model[T]{
		Selected: in.GetSelected(),
		Options:  in.GetOptions(),
	}
}

// Model is a base implementation of the Selector interface.
type Model[T comparable] struct {
	Selected T   `json:"selected,omitempty"` // currently selected value
	Options  []T `json:"options,omitempty"`  // list of available options
}

func (a Model[T]) IsSelected() bool {
	var zero T
	return a.Selected != zero
}
func (a Model[T]) GetSelected() T  { return a.Selected }
func (a Model[T]) GetOptions() []T { return a.Options }

// Add adds a new option to the Options list, and sets it as selected
// if no value has been selected yet. Zero values are ignored.
// Duplicates are not added.
func (a *Model[T]) Add(value T) {
	var zero T
	if value == zero {
		return
	}
	if a.Selected == zero {
		a.Selected = value
	}
	for _, opt := range a.Options {
		if opt == value {
			return
		}
	}
	a.Options = append(a.Options, value)
}
