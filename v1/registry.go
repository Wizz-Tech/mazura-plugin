package plugin

import (
	"errors"
	"sync"
)

// ErrPluginExists indicates that a plugin with the specified identifier is already registered in the registry.
var ErrPluginExists = errors.New("plugin already registered")

// Store is a generic map-based type for storing keyed values, parameterized by any data type provided.
type Store[T any] map[string]T

// GetPlugin retrieves and return the according plugin if found.
func GetPlugin[T any](registryList *Registry, packageName string) (*T, error) {
	if p, ok := registryList.plugins[packageName]; ok {
		casted := any(p).(T)

		return &casted, nil
	}

	return nil, errors.New("plugin not found")
}

// Registry holds loaded plugins.
type Registry struct {
	mu      sync.RWMutex
	plugins Store[Plugin]
}

// RegistryList is a global pointer to the plugin registry holding loaded plugins and manages thread-safe access.
var RegistryList *Registry

// InitRegistry creates a new plugin registry.
func InitRegistry() {
	RegistryList = &Registry{
		plugins: make(map[string]Plugin),
	}
}

// RegisterPlugin adds a plugin to the registry.
func RegisterPlugin[T any](
	store map[string]T,
	mu *sync.RWMutex,
	pluginToRegister T,
	getRegistrationIdentifier func(T) string,
) error {
	mu.Lock()
	defer mu.Unlock()

	name := getRegistrationIdentifier(pluginToRegister)
	if _, exists := store[name]; exists {
		return ErrPluginExists
	}

	store[name] = pluginToRegister

	return nil
}

// GetPluginStore retrieves all plugins of a specified type T from the given
// registry. It iterates through the registry's plugin store and casts each
// plugin to the desired type, adding valid ones to the result.
func GetPluginStore[T any](registryList *Registry) []T {
	var result []T

	for _, plugin := range registryList.plugins {
		if casted, ok := any(plugin).(T); ok {
			result = append(result, casted)
		}
	}

	return result
}
