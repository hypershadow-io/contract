package meta

// Make creates a new Meta map with optional pre-allocated capacity.
func Make(counts ...int) Meta {
	var count int
	if len(counts) > 0 {
		count = counts[0]
	}
	return make(Meta, count)
}

// Meta is a generic key-value store for passing metadata between components/plugins.
// Keys are strings, and values can be of any type.
//
// It is commonly used to attach structured context to models, requests, events, etc.
type Meta map[string]any

// IsValid returns true if the Meta map is non-nil.
func (a Meta) IsValid() bool { return a != nil }

// IsZero returns true if the Meta map is empty or uninitialized.
func (a Meta) IsZero() bool { return len(a) == 0 }

// Set assigns the given key and value to the Meta map,
// returning itself or a new instance if Meta was not initialized.
func (a Meta) Set(key string, value string) Meta {
	result := a
	if !result.IsValid() {
		result = Make(1)
	}
	result[key] = value
	return result
}

// Merge combines the current Meta map with another one.
// If the original map is not initialized, a new one is created.
// Values from the other map will overwrite any existing keys.
func (a Meta) Merge(m Meta) Meta {
	result := a
	if !result.IsValid() {
		result = Make(len(m))
	}
	for k, v := range m {
		result[k] = v
	}
	return result
}
