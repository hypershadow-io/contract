package plugin

import (
	"context"

	"github.com/hypershadow-io/contract/di"
)

type (
	// FunctionInstallSignature defines an optional one-time installation step that is invoked only once,
	// when the plugin is first installed into the system. It allows the plugin to perform initial setup tasks,
	// such as prompting the administrator for configuration input, registering with external services,
	// or preparing persistent resources required for future runs.
	//
	// This method is optional and may be omitted by plugins that do not require an installation step.
	// If not implemented, the host will simply skip the Install phase.
	FunctionInstallSignature = func(di.DI) error

	// FunctionInitSignature is called first for every plugin.
	// It allows the plugin to register its own dependencies into the DI container.
	// Returns either the same or a new DI container (e.g. with overridden scope).
	// This step should not rely on any external dependencies or configuration.
	//
	// This function is optional and may be omitted by plugins that do not require initialize.
	FunctionInitSignature = func(di.DI) (di.DI, error)

	// FunctionConfigureSignature is called after Init to provide raw configuration data to the plugin.
	// This step allows the plugin to receive its own isolated configuration in a raw []byte form.
	// The plugin is responsible for decoding, validating, or decrypting the config using its own logic,
	// typically during the subsequent Prepare step when all required dependencies are available.
	//
	// This function is optional and may be omitted by plugins that do not require configuration.
	// In such cases, the host will simply skip the Configure step if the function is not implemented by the plugin.
	FunctionConfigureSignature = func(data string) error

	// FunctionPrepareSignature is called after configuration is set.
	// At this point, the plugin can retrieve any required dependencies from the DI container
	// and initialize its internal state.
	//
	// This function is optional and may be omitted by plugins that do not require preparing.
	FunctionPrepareSignature = func(di.DI) error

	// FunctionRunSignature is the final step, where the plugin starts its execution logic.
	// All dependencies and configuration should already be initialized by this point.
	//
	// This function is optional and may be omitted by plugins that do not require running.
	FunctionRunSignature = func(c context.Context) error
)

const (
	// FunctionInstall is the name of the Install function expected in the plugin.
	// This function is optional and may be omitted by plugins that do not require installation.
	// If implemented, it will be called once during the plugin's installation into the system.
	FunctionInstall = "Install"

	// FunctionInit is the name of the Init function expected in the plugin.
	// This function is optional and may be omitted by plugins that do not require initialize.
	FunctionInit = "Init"

	// FunctionConfigure is the name of the Configure function expected in the plugin.
	// This function is optional and may be omitted by plugins that do not require configuration.
	FunctionConfigure = "Configure"

	// FunctionPrepare is the name of the Prepare function expected in the plugin.
	// This function is optional and may be omitted by plugins that do not require preparing.
	FunctionPrepare = "Prepare"

	// FunctionRun is the name of the Run function expected in the plugin.
	// This function is optional and may be omitted by plugins that do not require running.
	FunctionRun = "Run"
)
