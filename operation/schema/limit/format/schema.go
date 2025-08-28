package format

// Limit extends the schema.Limit interface with additional constraints.
type Limit interface {
	// GetFormat returns a hint about the expected format of the property's value.
	GetFormat() string
}
