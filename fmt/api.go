package fmt

import "io"

// Client defines a pluggable formatting and printing interface,
// compatible with the standard fmt package.
// It allows redirecting all formatting and output logic through a custom implementation.
type Client interface {
	Print(a ...any) (n int, err error)
	Println(a ...any) (n int, err error)
	Printf(f string, a ...any) (n int, err error)

	Sprint(a ...any) string
	Sprintln(a ...any) string
	Sprintf(f string, a ...any) string

	Fprint(w io.Writer, a ...any) (n int, err error)
	Fprintln(w io.Writer, a ...any) (n int, err error)
	Fprintf(w io.Writer, f string, a ...any) (n int, err error)

	Errorf(f string, a ...any) error
}

// Print delegates to the default Client.Print.
func Print(a ...any) (n int, err error) {
	return defaultClient.Print(a...)
}

// Println delegates to the default Client.Println.
func Println(a ...any) (n int, err error) {
	return defaultClient.Println(a...)
}

// Printf delegates to the default Client.Printf.
func Printf(f string, a ...any) (n int, err error) {
	return defaultClient.Printf(f, a...)
}

// Sprint delegates to the default Client.Sprint.
func Sprint(a ...any) string {
	return defaultClient.Sprint(a...)
}

// Sprintln delegates to the default Client.Sprintln.
func Sprintln(a ...any) string {
	return defaultClient.Sprintln(a...)
}

// Sprintf delegates to the default Client.Sprintf.
func Sprintf(f string, a ...any) string {
	return defaultClient.Sprintf(f, a...)
}

// Fprint delegates to the default Client.Fprint.
func Fprint(w io.Writer, a ...any) (n int, err error) {
	return defaultClient.Fprint(w, a...)
}

// Fprintln delegates to the default Client.Fprintln.
func Fprintln(w io.Writer, a ...any) (n int, err error) {
	return defaultClient.Fprintln(w, a...)
}

// Fprintf delegates to the default Client.Fprintf.
func Fprintf(w io.Writer, f string, a ...any) (n int, err error) {
	return defaultClient.Fprintf(w, f, a...)
}

// Errorf delegates to the default Client.Errorf.
func Errorf(f string, a ...any) error {
	return defaultClient.Errorf(f, a...)
}

// Init sets the default FMT client.
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
	defaultClient Client // default FMT client
)
