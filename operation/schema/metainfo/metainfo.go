package metainfo

import (
	"github.com/hypershadow-io/contract/identity"
	"github.com/hypershadow-io/contract/metainfo/field"
	"github.com/hypershadow-io/contract/operation/schema"
)

// Fields defines the core set of fields used to describe an Operation schema
// for registration in the metadata registry.
// In essence, this corresponds to the metadata representation of the schema.Schema interface.
var Fields = []field.Field{
	field.Model{
		Identification: identity.Model{
			Key:         "action",
			Name:        "Action",
			Description: "Specifies the type of operation to be performed when executing the request. The action defines the semantic intent of the operation and may vary depending on the underlying protocol or target system.",
		},
		Type: schema.PropertyTypeString.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "externalId",
			Name:        "External ID",
			Description: "Represents an external or predefined identifier associated with the operation. It can be used to reference the operation in external systems, documentation, or tooling.",
		},
		Type: schema.PropertyTypeString.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "description",
			Name:        "Description",
			Description: "Provides a machine-oriented description of the operation, intended to help AI systems understand the purpose and behavior of the operation.",
		},
		Type:   propertyTypeTextarea.String(),
		Select: true,
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes",
			Name:        "Attributes",
			Description: "Describes the structure and metadata of a single field within the data schema, including its type, validation rules, access requirements, and optional nested properties.",
		},
		Type: schema.PropertyTypeObject.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "components",
			Name:        "Components",
			Description: "Defines external reusable components that can be referenced using 'Ref'. It is represented as an array of objects, where each object describes a component available for referencing.",
		},
		Type: schema.PropertyTypeArray.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "components.path",
			Name:        "Path",
			Description: "Specifies the reference path to the component, used for linking via 'Ref'.",
		},
		Type: schema.PropertyTypeString.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "components.value",
			Name:        "Value",
			Description: "Defines the actual content of the component that is referenced by the given 'Path'.",
		},
		Type: schema.PropertyTypeObject.String(),
		Ref:  "attributes",
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.auth",
			Name:        "Authorization",
			Description: "Marks the field as authorization-related, indicating that its value will be automatically populated by the system based on the current authentication context.",
		},
		Type: schema.PropertyTypeObject.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.auth.name",
			Name:        "Alternative name",
			Description: "Specifies an alternative name to be used for the authorization field when mapping its value into the target request.",
		},
		Type: schema.PropertyTypeString.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.auth.type",
			Name:        "Type",
			Description: "Defines the type of authorization to be applied to the field. The value must be selected from a list of supported authorization types.",
		},
		Type: schema.PropertyTypeString.String(),
		Enum: []string{
			schema.AuthTypeBasic.String(),
			schema.AuthTypeBearer.String(),
			schema.AuthTypeApiKey.String(),
		},
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.name",
			Name:        "Name",
			Description: "Specifies the name of the field as it should appear in the target request or schema.",
		},
		Type: schema.PropertyTypeString.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.description",
			Name:        "Description",
			Description: "Provides a detailed explanation of the field’s purpose and semantics within the schema. Intended to help AI systems and developers understand how the field should be used.",
		},
		Type:   propertyTypeTextarea.String(),
		Select: true,
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.section",
			Name:        "Section",
			Description: "Specifies the location within the request where the field is applied.",
		},
		Type: schema.PropertyTypeString.String(),
		Enum: []string{
			schema.PropertySectionPath.String(),
			schema.PropertySectionQuery.String(),
			schema.PropertySectionHeader.String(),
			schema.PropertySectionBody.String(),
			schema.PropertySectionCookie.String(),
		},
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.type",
			Name:        "Type",
			Description: "Defines the data type of the field according to the schema specification.",
		},
		Type:   schema.PropertyTypeString.String(),
		Select: true,
		Enum: []string{
			schema.PropertyTypeString.String(),
			schema.PropertyTypeNumber.String(),
			schema.PropertyTypeInteger.String(),
			schema.PropertyTypeBoolean.String(),
			schema.PropertyTypeObject.String(),
			schema.PropertyTypeArray.String(),
		},
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.required",
			Name:        "Required",
			Description: "Indicates whether the field must be provided in order to execute the operation.",
		},
		Type: schema.PropertyTypeBoolean.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit",
			Name:        "Conditions and restrictions",
			Description: "Defines constraints and validation rules that apply to the field’s value.",
		},
		Type: schema.PropertyTypeObject.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.properties",
			Name:        "Properties",
			Description: "Describes the structure of nested fields if the field is an object, or the structure of array elements if the field is an array.",
		},
		Type: schema.PropertyTypeArray.String(),
		Ref:  "attributes",
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.ref",
			Name:        "Ref",
			Description: "Specifies a reference to a separately defined type that describes the structure of the field. The referenced component is defined in the 'Components' section.",
		},
		Type: schema.PropertyTypeString.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit.enum",
			Name:        "Allowed values",
			Description: "Specifies the set of allowed values that the field can accept.",
		},
		Type: schema.PropertyTypeString.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit.min",
			Name:        "Minimum",
			Description: "Specifies the minimum allowed value for numeric fields or the minimum length for string fields.",
		},
		Type: schema.PropertyTypeNumber.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit.exclusiveMin",
			Name:        "Exclusive minimum",
			Description: "Indicates whether the minimum value is exclusive, meaning the field's value must be strictly greater than the defined minimum.",
		},
		Type: schema.PropertyTypeBoolean.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit.max",
			Name:        "Maximum",
			Description: "Specifies the maximum allowed value for numeric fields or the maximum length for string fields.",
		},
		Type: schema.PropertyTypeNumber.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit.exclusiveMax",
			Name:        "Exclusive maximum",
			Description: "Indicates whether the maximum value is exclusive, meaning the field's value must be strictly less than the defined maximum.",
		},
		Type: schema.PropertyTypeBoolean.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit.multiple",
			Name:        "Multiple",
			Description: "Specifies that the field’s value must be a multiple of the defined number.",
		},
		Type: schema.PropertyTypeNumber.String(),
	},
	field.Model{
		Identification: identity.Model{
			Key:         "attributes.limit.unique",
			Name:        "Unique",
			Description: "Indicates that if the field is an array, all its elements must be unique.",
		},
		Type: schema.PropertyTypeBoolean.String(),
	},
}

const (
	// propertyTypeTextarea is a custom property type representing a multi-line string input.
	// Functionally similar to schema.PropertyTypeString, but intended for longer or formatted text.
	propertyTypeTextarea schema.PropertyType = "textarea"
)
