package impl

import (
	"github.com/hypershadow-io/contract/choice"
	"github.com/hypershadow-io/contract/operation/schema"
)

// propertyNil provides a no-op implementation of the schema.Property interface.
// It represents the absence of any property constraints and always returns zero values.
type propertyNil struct{}

func (a propertyNil) IsValid() bool                               { return false }
func (a propertyNil) GetRef() string                              { return "" }
func (a propertyNil) GetAuth(bool) schema.Auth                    { return Auth{} }
func (a propertyNil) GetName(bool) string                         { return "" }
func (a propertyNil) GetDescription(bool) choice.Selector[string] { return choice.Model[string]{} }
func (a propertyNil) GetSection(bool) schema.PropertySection      { return "" }
func (a propertyNil) IsRequired(bool) bool                        { return false }
func (a propertyNil) GetType(bool) []schema.PropertyType          { return nil }
func (a propertyNil) GetAllOf(bool) []schema.Property             { return nil }
func (a propertyNil) GetOneOf(bool) []schema.Property             { return nil }
func (a propertyNil) GetAnyOf(bool) []schema.Property             { return nil }
func (a propertyNil) GetProperties(bool) []schema.Property        { return nil }
func (a propertyNil) GetLimit(bool) schema.Limit                  { return limitNil{} }
