package runner

type (
	// Client defines an interface for registering and controlling long-lived or background commands.
	// This plugin provides infrastructure to "daemonize" components and manage their lifecycle
	// based on context and application state.
	Client interface {
		// Add registers a new command to be managed by the runner.
		// The command will be executed and controlled as part of the system's lifecycle.
		Add(cmd Command)
	}

	// Command defines the interface for a unit of execution managed by the runner.
	// It can represent a long-running background task, a persistent service, or other plugin process.
	Command interface {
		// Run starts the command's execution. It should block until the command completes or is canceled.
		//
		// Note: As soon as any command returns (successfully or with an error), the entire runner,
		// all other commands, and the application itself will be stopped.
		Run() error

		// Dispose is called to release any resources or stop background processes.
		// It is guaranteed to be called even if Run returns early or fails.
		Dispose()
	}
)
