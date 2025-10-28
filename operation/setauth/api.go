package setauth

import (
	"github.com/hypershadow-io/contract/operation/schema"
)

// Client defines an interface for injecting authorization values into the parameter set
// according to the property definition.
type Client interface {
	// SetAuth populates the provided parameters with authorization values
	// based on the rules defined in the property.
	SetAuth(prop schema.Property, params any, auth map[string]string) any
}
