package codec

import "io"

type (
	// Client defines a generic serialization interface for encoding and decoding data.
	Client interface {
		// Marshal serializes the given value into a byte slice.
		Marshal(value any) ([]byte, error)

		// Unmarshal deserializes the byte slice into the provided destination value.
		Unmarshal(data []byte, value any) error

		// NewEncoder returns a streaming encoder that writes encoded data to the given writer.
		NewEncoder(w io.Writer) Encoder

		// NewDecoder returns a streaming decoder that reads and decodes data from the given reader.
		NewDecoder(r io.Reader) Decoder
	}

	// Encoder defines a streaming encoder that writes serialized data to an underlying writer.
	Encoder interface {
		// Encode serializes the given value and writes it to the output stream.
		Encode(value any) error
	}

	// Decoder defines a streaming decoder that reads from an input stream and deserializes data.
	Decoder interface {
		// Decode reads the next value from the input stream and populates the provided target.
		Decode(value any) error
	}
)
