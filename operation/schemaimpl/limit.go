package schemaimpl

import (
	"github.com/hypershadow-io/contract/operation/schema"
)

// NewLimitFrom converts a schema.Limit interface to base *Limit implementation.
// If the provided schema.Limit is already of type *Limit, it returns it directly.
func NewLimitFrom(in schema.Limit) *Limit {
	if res, ok := in.(*Limit); ok {
		return res
	}
	return &Limit{
		Enum:         in.GetEnum(),
		Min:          in.GetMin(),
		ExclusiveMin: in.IsExclusiveMin(),
		Max:          in.GetMax(),
		ExclusiveMax: in.IsExclusiveMax(),
		Multiple:     in.GetMultiple(),
		Unique:       in.IsUnique(),
	}
}

// Limit is a base implementation of the schema.Limit interface.
type Limit struct {
	Enum         []string `json:"enum,omitempty"`
	Min          *float64 `json:"min,omitzero"`
	ExclusiveMin bool     `json:"exclusiveMin,omitempty"`
	Max          *float64 `json:"max,omitzero"`
	ExclusiveMax bool     `json:"exclusiveMax,omitempty"`
	Multiple     *float64 `json:"multiple,omitzero"`
	Unique       bool     `json:"unique,omitempty"`
	Pattern      string   `json:"pattern,omitempty"`
}

func (a Limit) IsValid() bool         { return true }
func (a Limit) GetEnum() []string     { return a.Enum }
func (a Limit) GetMin() *float64      { return a.Min }
func (a Limit) SetMin(v float64)      { a.Min = &v }
func (a Limit) IsExclusiveMin() bool  { return a.ExclusiveMin }
func (a Limit) GetMax() *float64      { return a.Max }
func (a Limit) SetMax(v float64)      { a.Max = &v }
func (a Limit) IsExclusiveMax() bool  { return a.ExclusiveMax }
func (a Limit) GetMultiple() *float64 { return a.Multiple }
func (a Limit) SetMultiple(v float64) { a.Multiple = &v }
func (a Limit) IsUnique() bool        { return a.Unique }
func (a Limit) GetPattern() string    { return a.Pattern }
