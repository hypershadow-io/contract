package json

import (
	"github.com/hypershadow-io/contract/json/raw"
	"github.com/hypershadow-io/contract/meta"
)

// Meta is a network-safe representation of meta.Meta,
// designed for receiving structured metadata over the network (e.g., via HTTP)
// before transforming it into the internal Meta type.
type Meta map[string]raw.Message

// Meta converts json.Meta into the internal meta.Meta type.
func (a *Meta) Meta() meta.Meta {
	if a == nil {
		return nil
	}
	m := meta.Make(len(*a))
	for k, v := range *a {
		m[k] = []byte(v)
	}
	return m
}
