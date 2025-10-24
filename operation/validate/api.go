package validate

import (
	"github.com/hypershadow-io/contract/fielderror"
	"github.com/hypershadow-io/contract/operation/schema"
)

// Client defines an interface for validating input parameters against a given operation schema.
// It returns a list of field-level validation errors, or a general error if the validation process fails.
type Client interface {
	// Validate checks the provided parameters against the specified schema.
	// Returns a slice of field errors or a general error if validation could not be completed.
	Validate(prop schema.Property, params any) ([]fielderror.Error, error)
}
