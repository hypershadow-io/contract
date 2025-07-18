package schemaimpl

import (
	"github.com/hypershadow-io/contract/choice"
	"github.com/hypershadow-io/contract/operation/schema"
)

// propertyNil provides a no-op implementation of the schema.Property interface.
// It represents the absence of any property constraints and always returns zero values.
type propertyNil struct{}

func (a propertyNil) IsValid() bool                           { return false }
func (a propertyNil) GetRef() string                          { return "" }
func (a propertyNil) GetAuth() schema.Auth                    { return Auth{} }
func (a propertyNil) GetName() string                         { return "" }
func (a propertyNil) GetDescription() choice.Selector[string] { return choice.Model[string]{} }
func (a propertyNil) GetSection() schema.PropertySection      { return "" }
func (a propertyNil) IsRequired() bool                        { return false }
func (a propertyNil) GetType() []schema.PropertyType          { return nil }
func (a propertyNil) GetAllOf() []schema.Property             { return nil }
func (a propertyNil) GetOneOf() []schema.Property             { return nil }
func (a propertyNil) GetAnyOf() []schema.Property             { return nil }
func (a propertyNil) GetProperties() []schema.Property        { return nil }
func (a propertyNil) GetLimit() schema.Limit                  { return limitNil{} }
