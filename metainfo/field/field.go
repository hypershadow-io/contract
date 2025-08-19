package field

import "github.com/hypershadow-io/contract/identity"

// Field defines the interface for describing a single field.
type Field interface {
	identity.Identification

	// GetType returns the type of the field (e.g., "string", "int", "bool").
	GetType() string

	// IsSelect indicates whether the field supports a predefined set of selectable values.
	IsSelect() bool

	// GetEnum returns the list of allowed enum values for the field, if defined.
	GetEnum() []string

	// GetRef returns the external reference key for the field, if it is linked to another definition.
	GetRef() string
}

// Model provides a base implementation of the Field interface.
type Model struct {
	identity.Identification
	Type   string   `json:"type"`   // Field type (e.g. string, int, bool, etc.)
	Select bool     `json:"select"` // Indicates whether the field supports a predefined set of selectable values
	Enum   []string `json:"enum"`   // Optional list of allowed enum values
	Ref    string   `json:"ref"`    // Optional reference key to external definition
}

func (a Model) GetType() string   { return a.Type }
func (a Model) IsSelect() bool    { return a.Select }
func (a Model) GetEnum() []string { return a.Enum }
func (a Model) GetRef() string    { return a.Ref }
