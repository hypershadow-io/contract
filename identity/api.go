package identity

// Identification defines a common interface for identifiable entities.
type Identification interface {
	// GetKey returns the unique key of the entity.
	GetKey() string

	// GetName returns the display name of the entity.
	GetName() string

	// GetDescription returns the description of the entity.
	GetDescription() string
}

// Model is a base implementation of the Identification interface.
type Model struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a Model) GetKey() string         { return a.Key }
func (a Model) GetName() string        { return a.Name }
func (a Model) GetDescription() string { return a.Description }
