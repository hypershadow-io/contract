package schema

import "github.com/hypershadow-io/contract/choice"

type (
	// List represents a slice of schema definitions.
	List []Schema

	// Schema defines the interface for a structured input/output schema.
	// Used for describing operations, parameters, and nested components.
	Schema interface {
		// GetAction returns the operation type (e.g., GET, POST, PUT for HTTP; SELECT, UPDATE for SQL).
		GetAction() string

		// GetExternalID returns the external identifier of the schema from the customer's system.
		GetExternalID() string

		// GetDescription returns a user-friendly description of the schema with one selected value and multiple possible options.
		GetDescription() choice.Selector[string]

		// GetAttributes returns the root property definition for this schema.
		GetAttributes() Property

		// GetComponents returns named reusable properties used in the schema.
		GetComponents() map[string]Property

		// Resolve returns a property by its reference key from components.
		Resolve(ref string) Property
	}

	// Property defines the interface for describing a single property or parameter within a schema.
	Property interface {
		// IsValid returns true if the property is correctly defined.
		IsValid() bool

		// GetRef returns the reference key if the property is a component reference.
		GetRef() string

		// GetAuth returns the authorization configuration for this property.
		GetAuth() Auth

		// GetName returns the name of the property.
		GetName() string

		// GetDescription returns a user-friendly description of the property with one selected value and multiple possible options.
		GetDescription() choice.Selector[string]

		// GetSection returns the location (query, header, body, etc.) of this property.
		GetSection() PropertySection

		// GetType returns the list of data types this property can accept.
		GetType() []PropertyType

		// GetAllOf returns a list of properties that must all apply (AND).
		GetAllOf() []Property

		// GetOneOf returns a list where exactly one property must apply (XOR).
		GetOneOf() []Property

		// GetAnyOf returns a list where any number of properties can apply (OR).
		GetAnyOf() []Property

		// GetProperties returns nested properties (used for objects and array).
		GetProperties() []Property

		// IsRequired indicates whether this property must be provided.
		IsRequired() bool

		// GetLimit returns validation constraints on the property's value.
		GetLimit() Limit
	}

	// Auth defines the authorization method required for the associated schema or property.
	Auth interface {
		// IsValid returns true if the auth configuration is valid.
		IsValid() bool

		// GetName returns the name of the auth method or key.
		GetName() string

		// GetType returns the authorization type (basic, bearer, apiKey, etc.).
		GetType() AuthType
	}

	// Limit defines the constraints applicable to a schema property.
	Limit interface {
		// IsValid returns true if constraints are defined correctly.
		IsValid() bool

		// GetEnum returns a list of allowed string values (enumeration).
		GetEnum() []string

		// GetMin returns the minimum allowed value (inclusive/exclusive based on IsExclusiveMin).
		GetMin() *float64

		// IsExclusiveMin returns true if the minimum bound must be strictly greater than the specified value.
		IsExclusiveMin() bool

		// GetMax returns the maximum allowed value (inclusive/exclusive based on IsExclusiveMax).
		GetMax() *float64

		// IsExclusiveMax returns true if the maximum bound must be strictly less than the specified value.
		IsExclusiveMax() bool

		// GetMultiple returns the required multiple constraint (e.g., must be divisible by X).
		GetMultiple() *float64

		// IsUnique indicates whether all values in an array must be unique.
		IsUnique() bool

		// GetPattern returns a regex pattern that the value must match.
		GetPattern() string
	}

	// PropertySection defines the location of the parameter in the request.
	PropertySection string

	// PropertyType defines the supported data types of a property.
	PropertyType string

	// AuthType defines the supported authorization schemes.
	AuthType string
)

const (
	PropertySectionPath   PropertySection = "path"
	PropertySectionQuery  PropertySection = "query"
	PropertySectionHeader PropertySection = "header"
	PropertySectionBody   PropertySection = "body"
	PropertySectionCookie PropertySection = "cookie"

	PropertyTypeString  PropertyType = "string"
	PropertyTypeNumber  PropertyType = "number"
	PropertyTypeInteger PropertyType = "integer"
	PropertyTypeBoolean PropertyType = "boolean"
	PropertyTypeObject  PropertyType = "object"
	PropertyTypeArray   PropertyType = "array"

	AuthTypeBasic  AuthType = "basic"
	AuthTypeBearer AuthType = "bearer"
	AuthTypeApiKey AuthType = "apiKey"
)

// FindByExternalID searches the schema list for a schema with the specified external ID.
func (a List) FindByExternalID(externalID string) (res_ Schema, found_ bool) {
	for _, item := range a {
		if item.GetExternalID() == externalID {
			return item, true
		}
	}
	return res_, false
}

func (a PropertySection) String() string { return string(a) }
func (a PropertyType) String() string    { return string(a) }
func (a AuthType) String() string        { return string(a) }
