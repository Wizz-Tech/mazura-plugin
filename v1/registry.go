package plugin

import (
	"errors"
	"sync"
)

var ErrPluginExists = errors.New("plugin already registered")

type PluginStore[T any] map[string]T

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
	plugins PluginStore[Plugin]
}

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

func GetPluginStore[T any](registryList *Registry) []T {
	var result []T

	for _, plugin := range registryList.plugins {
		if casted, ok := any(plugin).(T); ok {
			result = append(result, casted)
		}
	}

	return result
}
