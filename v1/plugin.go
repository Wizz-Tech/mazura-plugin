package plugin

// PluginConfigField represents a configuration field for a plugin.
type PluginConfigField struct {
	// Key is the unique identifier for the configuration field.
	Key string `json:"key"` // config key, e.g. "apiKey"

	// Label is the user-friendly name displayed for the configuration field.
	Label string `json:"label"` // user-friendly label, e.g. "API Key"

	// Type defines the input type, such as "text", "password", or "select".
	Type string `json:"type"` // "text", "password", "select", "checkbox", etc.

	// Required indicates whether the field is mandatory for configuration.
	Required bool `json:"required"` // whether the field is required

	// Placeholder provides optional placeholder text for the input field.
	Placeholder string `json:"placeholder"` // optional input placeholder

	// Options lists the available values for "select" type fields.
	Options []string `json:"options,omitempty"` // for "select" fields

	// HelpText provides an optional helper message for the field.
	HelpText string `json:"helpText,omitempty"` // small helper text

	// Default specifies the default value of the field, if any.
	Default any `json:"default,omitempty"` // optional default value

	// Group categorizes or groups the field under a specific section.
	Group string `json:"group,omitempty"` // optional grouping/category

	// Extra contains additional metadata for UI or other purposes.
	Extra map[string]string `json:"extra,omitempty"` // any extra UI metadata
}

// Plugin is the interface that all plugins must implement.
type Plugin interface {
	// PackageName returns the package name of the plugin for identification purposes.
	PackageName() string

	// Name Unique plugin name (used for identification)
	Name() string

	// Init initializes the plugin with the provided configuration and router, returning an error if initialization fails.
	Init(config map[string]string, r Router) error

	// Shutdown gracefully stops the plugin, releasing any allocated resources or performing cleanup tasks. Returns an error if the process fails.
	Shutdown() error

	// Version returns the current version of the plugin as a string.
	Version() string

	// Description returns a brief explanation of the plugin's functionality or purpose as a string.
	Description() string

	// GetConfig retrieves a list of configuration fields required by the plugin, describing their properties and constraints.
	GetConfig() []PluginConfigField
}
