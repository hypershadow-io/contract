package id

type (
	// Client defines an interface for generating and handling unique identifiers,
	// used across the system for consistent ID creation in plugins and core services.
	Client interface {
		// NewIDInt64 - generate a unique numeric ID
		NewIDInt64() int64

		// NewIDUint64 - generate a unique unsigned numeric ID
		NewIDUint64() uint64

		// NewIDString - generate a unique string ID
		NewIDString() string

		// CreateID - create a unique bytes ID based on the input data
		CreateID(data []byte) []byte

		// CreateIDString - create a unique string ID based on the input data
		CreateIDString(data []byte) string

		// CreateIDStringFromStrings - create a unique string ID based on the input data
		CreateIDStringFromStrings(data ...string) string

		// BinToString - converting bytes ID to string
		BinToString(id []byte) string

		// StringToBin - converting string ID to bytes
		StringToBin(id string) []byte

		// CreateString - creates string of specified length with random chars from charset treated as byte slice
		// Note: If charset is empty returns empty string
		CreateString(charset Charset, length int) string
	}

	// Charset задает пользовательский набор символов, используемый для генерации случайных строк.
	Charset string
)

const (
	// AlphaLowercaseCharset lowercase letters
	AlphaLowercaseCharset Charset = "abcdefghijklmnopqrstuvwxyz"
	// AlphaUppercaseCharset uppercase letters
	AlphaUppercaseCharset Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// AlphaCharset lowercase and uppercase letters
	AlphaCharset = AlphaLowercaseCharset + AlphaUppercaseCharset
	// NumericCharset digits
	NumericCharset Charset = "0123456789"
	// AlphaNumCharset lowercase, uppercase letters and digits
	AlphaNumCharset = AlphaCharset + NumericCharset
)
