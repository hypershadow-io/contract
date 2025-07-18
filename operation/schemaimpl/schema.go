package schemaimpl

import (
	"github.com/hypershadow-io/contract/choice"
	"github.com/hypershadow-io/contract/operation/schema"
)

// MakeSchemaFrom converts a schema.Schema interface to base Schema implementation.
// If the provided schema.Schema is already of type Schema, it returns it directly.
func MakeSchemaFrom(in schema.Schema) Schema {
	if res, ok := in.(Schema); ok {
		return res
	}
	return Schema{
		Description: choice.MakeFrom(in.GetDescription()),
		Action:      in.GetAction(),
		ExternalID:  in.GetExternalID(),
		Attributes:  MakePropertyFrom(in.GetAttributes()),
		Components:  MakeComponentFrom(in.GetComponents()),
	}
}

// Schema is a base implementation of the schema.Schema interface.
type Schema struct {
	Description choice.Model[string] `json:"description,omitempty"`
	Action      string               `json:"action,omitempty"`
	ExternalID  string               `json:"externalId,omitempty"`
	Attributes  Property             `json:"attributes,omitempty"`
	Components  map[string]Property  `json:"components,omitempty"`
	// TODO rathil add tags!!!
	// TODO rathil add response!!!
}

func (a Schema) GetAction() string                       { return a.Action }
func (a Schema) GetExternalID() string                   { return a.ExternalID }
func (a Schema) GetDescription() choice.Selector[string] { return a.Description }

func (a Schema) GetAttributes() schema.Property {
	a.Attributes.schema = &a
	return a.Attributes
}

func (a Schema) GetComponents() map[string]schema.Property {
	result := make(map[string]schema.Property, len(a.Components))
	for path, value := range a.Components {
		value.schema = &a
		result[path] = value
	}
	return result
}

func (a Schema) Resolve(ref string) schema.Property {
	if ref == "" {
		return propertyNil{}
	}
	if result, ok := a.Components[ref]; ok {
		result.schema = &a
		return result
	}
	return propertyNil{}
}
