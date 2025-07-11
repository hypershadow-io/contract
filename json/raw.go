package json

import "errors"

type (
	// RawMessage - equivalent of json.RawMessage
	RawMessage []byte
	// Raw - equivalent of json.RawMessage
	Raw = RawMessage
)

func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
