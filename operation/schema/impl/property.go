package impl

import (
	"github.com/hypershadow-io/contract/choice"
	"github.com/hypershadow-io/contract/operation/schema"
	"github.com/hypershadow-io/contract/utilslice"
)

// MakeComponentFrom converts a map of schema.Property interfaces to a map of concrete Property implementations.
// Each property in the input map is individually converted via MakePropertyFrom.
func MakeComponentFrom(in map[string]schema.Property) map[string]Property {
	result := make(map[string]Property, len(in))
	for path, value := range in {
		result[path] = MakePropertyFrom(value)
	}
	return result
}

// MakePropertyFrom converts a schema.Property interface to base Property implementation.
// If the provided schema.Property is already of type Property, it returns it directly.
func MakePropertyFrom(in schema.Property) Property {
	if res, ok := in.(Property); ok {
		return res
	}
	return Property{
		Ref:         in.GetRef(),
		Auth:        MakeAuthFrom(in.GetAuth()),
		Name:        in.GetName(),
		Description: choice.MakeFrom(in.GetDescription()),
		Section:     in.GetSection(),
		Type:        in.GetType(),
		AllOf:       utilslice.SliceToSlice(in.GetAllOf(), MakePropertyFrom),
		OneOf:       utilslice.SliceToSlice(in.GetOneOf(), MakePropertyFrom),
		AnyOf:       utilslice.SliceToSlice(in.GetAnyOf(), MakePropertyFrom),
		Properties:  utilslice.SliceToSlice(in.GetProperties(), MakePropertyFrom),
		Required:    in.IsRequired(),
		Limit:       NewLimitFrom(in.GetLimit()),
	}
}

// Property is a base implementation of the schema.Property interface.
type Property struct {
	Ref         string                 `json:"ref,omitempty"`
	Auth        Auth                   `json:"auth,omitzero"`
	Name        string                 `json:"name,omitempty"`
	Description choice.Model[string]   `json:"description,omitzero"`
	Section     schema.PropertySection `json:"section,omitempty"`
	Type        []schema.PropertyType  `json:"type,omitempty"`
	AllOf       []Property             `json:"allOf,omitempty"`
	OneOf       []Property             `json:"oneOf,omitempty"`
	AnyOf       []Property             `json:"anyOf,omitempty"`
	Properties  []Property             `json:"properties,omitempty"`
	Required    bool                   `json:"required,omitempty"`
	Limit       *Limit                 `json:"limit,omitzero"`
	schema      *Schema
}

func (a Property) IsValid() bool  { return a.Name != "" }
func (a Property) GetRef() string { return a.Ref }

func (a Property) GetAuth(resolveRef bool) schema.Auth {
	if a.Auth.IsValid() {
		return a.Auth
	}
	if resolveRef {
		return a.resolve().GetAuth()
	}
	return Auth{}
}

func (a Property) GetName(resolveRef bool) string {
	if a.Name != "" {
		return a.Name
	}
	if resolveRef {
		return a.resolve().GetName()
	}
	return ""
}

func (a Property) GetDescription(resolveRef bool) choice.Selector[string] {
	if len(a.Description.Options) > 0 {
		return a.Description
	}
	if resolveRef {
		return a.resolve().GetDescription()
	}
	return choice.Model[string]{}
}

func (a Property) GetSection(resolveRef bool) schema.PropertySection {
	if a.Section != "" {
		return a.Section
	}
	if resolveRef {
		return a.resolve().GetSection()
	}
	return ""
}

func (a Property) GetType(resolveRef bool) []schema.PropertyType {
	if len(a.Type) > 0 {
		return a.Type
	}
	if resolveRef {
		return a.resolve().GetType()
	}
	return nil
}

func (a Property) GetAllOf(resolveRef bool) []schema.Property {
	if len(a.AllOf) > 0 {
		return utilslice.SliceToSlice(a.AllOf, func(item Property) schema.Property {
			item.schema = a.schema
			return item
		})
	}
	if resolveRef {
		return a.resolve().GetAllOf()
	}
	return nil
}

func (a Property) GetOneOf(resolveRef bool) []schema.Property {
	if len(a.OneOf) > 0 {
		return utilslice.SliceToSlice(a.OneOf, func(item Property) schema.Property {
			item.schema = a.schema
			return item
		})
	}
	if resolveRef {
		return a.resolve().GetOneOf()
	}
	return nil
}

func (a Property) GetAnyOf(resolveRef bool) []schema.Property {
	if len(a.AnyOf) > 0 {
		return utilslice.SliceToSlice(a.AnyOf, func(item Property) schema.Property {
			item.schema = a.schema
			return item
		})
	}
	if resolveRef {
		return a.resolve().GetAnyOf()
	}
	return nil
}

func (a Property) GetProperties(resolveRef bool) []schema.Property {
	if len(a.Properties) > 0 {
		return utilslice.SliceToSlice(a.Properties, func(item Property) schema.Property {
			item.schema = a.schema
			return item
		})
	}
	if resolveRef {
		return a.resolve().GetProperties()
	}
	return nil
}

func (a Property) IsRequired(resolveRef bool) bool {
	if a.Required {
		return true
	}
	if resolveRef {
		return a.resolve().IsRequired()
	}
	return false
}

func (a Property) GetLimit(resolveRef bool) schema.Limit {
	if a.Limit != nil {
		return a.Limit
	}
	if resolveRef {
		return a.resolve().GetLimit()
	}
	return limitNil{}
}

// EnsureLimit ensures that the Limit field is initialized.
// If the Limit is nil, it creates a new instance and assigns it.
// Returns the (possibly newly created) Limit pointer.
func (a *Property) EnsureLimit() *Limit {
	if a.Limit == nil {
		a.Limit = &Limit{}
	}
	return a.Limit
}

func (a Property) resolve() schema.Property {
	if a.schema == nil {
		return propertyNil{}
	}
	return a.schema.Resolve(a.Ref)
}
