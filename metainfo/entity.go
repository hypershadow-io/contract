package metainfo

import (
	"github.com/hypershadow-io/contract/field"
	"github.com/hypershadow-io/contract/identity"
)

// Model provides the default implementation of the Entity interface.
type Model struct {
	identity.Identification
	Group          identity.Identification
	IntegratesKeys []string
	Fields         []field.Field
}

func (a Model) GetGroup() identity.Identification { return a.Group }
func (a Model) GetIntegratesKeys() []string       { return a.IntegratesKeys }
func (a Model) GetFields() []field.Field          { return a.Fields }
