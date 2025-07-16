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

// Identifiable is a base implementation of the Identification interface.
type Identifiable struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a Identifiable) GetKey() string         { return a.Key }
func (a Identifiable) GetName() string        { return a.Name }
func (a Identifiable) GetDescription() string { return a.Description }
