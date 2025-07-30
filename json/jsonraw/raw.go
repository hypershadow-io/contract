package jsonraw

import "errors"

// Message - equivalent of json.RawMessage
type Message []byte

func (m Message) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *Message) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("jsonraw.Message: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
