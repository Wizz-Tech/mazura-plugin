package plugin

import (
	"errors"
	"sync"
)

var (
	ErrPluginExists = errors.New("plugin already registered")
)

type PluginStore[T any] map[string]T

func GetPlugin[T any](registryList *Registry, storeName string, packageName string) (*T, error) {
	switch storeName {
	case "deployers":
		if p, ok := registryList.deployers[packageName]; ok {
			casted := any(p).(T)
			return &casted, nil
		}
	case "providers":
		if p, ok := registryList.providers[packageName]; ok {
			casted := any(p).(T)
			return &casted, nil
		}
	}

	return nil, errors.New("plugin not found")
}

// Registry holds loaded plugins.
type Registry struct {
	mu        sync.RWMutex
	plugins   PluginStore[Plugin]
	deployers PluginStore[DeployerPlugin]
	providers PluginStore[ProviderPlugin]
}

var RegistryList *Registry

// Init creates a new plugin registry.
func InitRegistry() {
	RegistryList = &Registry{
		plugins:   make(map[string]Plugin),
		deployers: make(map[string]DeployerPlugin),
		providers: make(map[string]ProviderPlugin),
	}
}

// Register adds a plugin to the registry.
func RegisterPlugin[T any](store map[string]T, mu *sync.RWMutex, p T, getRegistrationIdentifier func(T) string) error {
	mu.Lock()
	defer mu.Unlock()

	name := getRegistrationIdentifier(p)
	if _, exists := store[name]; exists {
		return ErrPluginExists
	}
	store[name] = p
	return nil
}

func GetPluginSore[T any](registryList *Registry, storeName string) []T {
	var result []T

	switch storeName {
	case "deployers":
		for _, plugin := range registryList.deployers {
			if casted, ok := any(plugin).(T); ok {
				result = append(result, casted)
			}
		}
	case "providers":
		for _, plugin := range registryList.providers {
			if casted, ok := any(plugin).(T); ok {
				result = append(result, casted)
			}
		}
	}

	return result
}

//func (r *Registry) GetPlugin(packageName string) (Plugin, bool) {
//	r.mu.RLock()
//	defer r.mu.RUnlock()
//	p, ok := r.plugins[packageName]
//	return p, ok
//}

func (r *Registry) GetDeployerPlugin(packageName string) (DeployerPlugin, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.deployers[packageName]
	return p, ok
}
