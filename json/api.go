package json

import (
	"io"

	"github.com/hypershadow-io/contract/codec"
)

// Client is an alias for codec.Client, representing a pluggable JSON serialization engine.
type Client codec.Client

// Marshal serializes the given value into JSON using the default client.
func Marshal(value any) ([]byte, error) { return defaultClient.Marshal(value) }

// Unmarshal parses the JSON-encoded data into the provided target value using the default client.
func Unmarshal(data []byte, value any) error { return defaultClient.Unmarshal(data, value) }

// NewEncoder returns a streaming JSON encoder that writes to the given writer.
func NewEncoder(w io.Writer) codec.Encoder { return defaultClient.NewEncoder(w) }

// NewDecoder returns a streaming JSON decoder that reads from the given reader.
func NewDecoder(r io.Reader) codec.Decoder { return defaultClient.NewDecoder(r) }

// Init sets the default JSON codec client.
// This function must be called once during application initialization.
// It is not thread-safe and should not be called concurrently.
// Subsequent calls after the first one are ignored.
func Init(client Client) {
	if defined {
		return
	}
	defined = true
	defaultClient = client
}

var (
	defined       bool
	defaultClient Client // default JSON client
)
