package impl

import (
	"github.com/hypershadow-io/contract/operation/schema"
)

// MakeAuthFrom converts a schema.Auth interface to base Auth implementation.
// If the provided schema.Auth is already of type Auth, it returns it directly.
func MakeAuthFrom(in schema.Auth) Auth {
	if res, ok := in.(Auth); ok {
		return res
	}
	return Auth{
		Name: in.GetName(),
		Type: in.GetType(),
	}
}

// Auth is a base implementation of the schema.Auth interface.
type Auth struct {
	Name string          `json:"name,omitempty"`
	Type schema.AuthType `json:"type,omitempty"`
}

func (a Auth) IsValid() bool            { return a.Name != "" || a.Type != "" }
func (a Auth) GetName() string          { return a.Name }
func (a Auth) GetType() schema.AuthType { return a.Type }
