package plugin

import (
	"github.com/Wizz-Tech/mazura-plugin/v1/router"
)

type PluginConfigField struct {
	Key         string            `json:"key"`                // config key, e.g. "apiKey"
	Label       string            `json:"label"`              // user-friendly label, e.g. "API Key"
	Type        string            `json:"type"`               // "text", "password", "select", "checkbox", etc.
	Required    bool              `json:"required"`           // whether the field is required
	Placeholder string            `json:"placeholder"`        // optional input placeholder
	Options     []string          `json:"options,omitempty"`  // for "select" fields
	HelpText    string            `json:"helpText,omitempty"` // small helper text
	Default     any               `json:"default,omitempty"`  // optional default value
	Group       string            `json:"group,omitempty"`    // optional grouping/category
	Extra       map[string]string `json:"extra,omitempty"`    // any extra UI metadata
}

// Plugin is the interface that all plugins must implement.
type Plugin interface {
	PackageName() string
	Name() string // Unique plugin name (used for identification)
	Init(config map[string]string, r router.Router) error
	Shutdown() error
	Version() string
	Description() string
	GetConfig() []PluginConfigField
}
