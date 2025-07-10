package runner

import "context"

// MakeCommand creates a new Command with an internal context and cancel function.
// The command runs the provided callback and can be gracefully disposed.
// When Dispose is called, the internal context is cancelled.
func MakeCommand(
	cb func(context.Context) error,
) Command {
	c, cancel := context.WithCancel(context.Background())
	return &unnamedCommand{
		c:       c,
		cb:      cb,
		dispose: cancel,
	}
}

// MakeCommandWithDispose creates a new Command using the provided callback and a custom dispose function.
// This allows external control over the command's shutdown behavior.
func MakeCommandWithDispose(
	cb func(context.Context) error,
	dispose context.CancelFunc,
) Command {
	return &unnamedCommand{
		c:       context.Background(),
		cb:      cb,
		dispose: dispose,
	}
}

// unnamedCommand is a default implementation of the Command interface.
// It wraps a callback function with a context and a dispose function,
// allowing it to be executed and gracefully shut down by the runner.
type unnamedCommand struct {
	c       context.Context             // internal context for the command execution
	cb      func(context.Context) error // execution logic
	dispose context.CancelFunc          // cleanup function called on Dispose
}

func (a *unnamedCommand) Run() error { return a.cb(a.c) }
func (a *unnamedCommand) Dispose()   { a.dispose() }
