package jsonint

import (
	"strconv"
	"strings"
)

// Int64 is a wrapper over int64 that marshals/unmarshals JSON as a string.
// Useful for ensuring compatibility with systems that require numeric values to be represented as strings (e.g., JavaScript or certain APIs).
type Int64 int64

// MarshalJSON implements the json.Marshaler interface.
// It serializes the int64 value as a JSON string.
func (a Int64) MarshalJSON() ([]byte, error) {
	if a == 0 {
		return []byte(`"0"`), nil
	}
	buf := make([]byte, 0, 22) // max uint64: len(`"18446744073709551615"`) = 22
	buf = append(buf, '"')
	buf = append(buf[:], strconv.FormatInt(int64(a), 10)...)
	buf = append(buf, '"')
	return buf, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It deserializes a JSON string (or number-as-string) into an int64.
func (a *Int64) UnmarshalJSON(raw []byte) error {
	x, err := strconv.ParseInt(strings.Trim(strings.TrimSpace(string(raw)), `"`), 10, 64)
	if err == nil {
		*a = Int64(x)
	}
	return err
}

// Int64 returns the native int64 value.
func (a Int64) Int64() int64 { return int64(a) }
