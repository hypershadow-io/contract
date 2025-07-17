package fielderror

// Error defines the interface for field-level errors.
type Error interface {
	// GetField returns the name of the field that caused the error.
	GetField() string

	// GetMessage returns the description of the error.
	GetMessage() string
}

// Model is a basic implementation of the Error interface.
type Model struct {
	Field   string // Field name associated with the error.
	Message string // Error message.
}

func (a Model) GetField() string   { return a.Field }
func (a Model) GetMessage() string { return a.Message }
