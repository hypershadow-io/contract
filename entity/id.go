package entity

import "strconv"

// MakeIDInt creates an ID from an int64 value.
func MakeIDInt(id int64) ID { return ID{Int: id} }

// MakeIDString creates an ID from a string value.
func MakeIDString(id string) ID { return ID{String: id} }

// ID represents a generic entity identifier that can be either numeric or string-based.
// Useful in cases where entity IDs may come from different sources or formats (e.g., DB vs external API).
type ID struct {
	Int    int64  // numeric ID
	String string // string-based ID
}

// IsValid returns true if either the numeric or string ID is set.
func (a ID) IsValid() bool {
	return a.Int > 0 || a.String != ""
}

// GetInt returns the numeric part of the ID.
func (a ID) GetInt() int64 { return a.Int }

// GetString returns the string part of the ID.
func (a ID) GetString() string { return a.String }

// ToString returns a string representation of the ID.
// If the numeric part is set, it takes priority.
func (a ID) ToString() string {
	if a.Int > 0 {
		return strconv.Itoa(int(a.Int))
	}
	return a.String
}
