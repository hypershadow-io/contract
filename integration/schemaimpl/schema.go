package schemaimpl

import (
	"github.com/hypershadow-io/contract/integration/schema"
	operationschema "github.com/hypershadow-io/contract/operation/schema"
)

// MakeSchemaFrom converts a schema.Schema interface to base Schema implementation.
// If the provided schema.Schema is already of type Schema, it returns it directly.
func MakeSchemaFrom(in schema.Schema) Schema {
	if res, ok := in.(Schema); ok {
		return res
	}
	return Schema{
		Title:      in.GetTitle(),
		Version:    in.GetVersion(),
		Operations: in.GetOperations(),
	}
}

// Schema is a base implementation of the schema.Schema interface.
type Schema struct {
	Title      string               `json:"title,omitempty"`
	Version    string               `json:"version,omitempty"`
	Operations operationschema.List `json:"operations,omitempty"`
}

func (a Schema) GetTitle() string                    { return a.Title }
func (a Schema) GetVersion() string                  { return a.Version }
func (a Schema) GetOperations() operationschema.List { return a.Operations }
