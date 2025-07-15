package metajson

import (
	"github.com/hypershadow-io/contract/json"
	"github.com/hypershadow-io/contract/meta"
)

// MetaJson is a network-safe representation of meta.Meta,
// designed for receiving structured metadata over the network (e.g., via HTTP)
// before transforming it into the internal Meta type.
type MetaJson map[string]json.RawMessage

// Meta converts MetaJson into the internal meta.Meta type.
func (a MetaJson) Meta() meta.Meta {
	m := meta.Make(len(a))
	for k, v := range a {
		m[k] = []byte(v)
	}
	return m
}
