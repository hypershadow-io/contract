package plugin

import "context"

// Client defines the interface for interacting with plugins within the system.
type Client interface {
	// IsSystem returns true if the plugin with the given ID is a system plugin.
	// System plugins are globally available, not tied to any specific organization,
	// and cannot be disabled at the organization level.
	IsSystem(pluginID string) bool

	// IsActive returns true if the plugin with the given ID is active for the organization
	// associated with the provided context.
	IsActive(c context.Context, pluginID string) bool
	// TODO rathil add Add, Load, etc.
}
