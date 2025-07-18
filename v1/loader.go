package plugin

import (
	"fmt"
	"path/filepath"
	pl "plugin"
	"reflect"
)

func load[T any](path string) (T, error) {
	var zero T

	p, err := pl.Open(path)
	if err != nil {
		return zero, fmt.Errorf("opening plugin: %w", err)
	}

	sym, err := p.Lookup("Plugin")
	if err != nil {
		return zero, fmt.Errorf("looking up symbol 'Plugin': %w", err)
	}

	// Use reflect to get the correct type information
	expectedType := reflect.TypeOf((*T)(nil)).Elem()
	symValue := reflect.ValueOf(sym)

	//fmt.Printf("Symbol type: %T\n", sym)
	//fmt.Printf("Expected type: %s\n", expectedType)
	//fmt.Printf("Symbol implements expected? %v\n", symValue.Type().Implements(expectedType))

	if !symValue.Type().Implements(expectedType) {
		return zero, fmt.Errorf("symbol 'Plugin' does not implement expected interface %s", expectedType)
	}

	return sym.(T), nil
}

// Load dynamically loads a Go plugin file and asserts the exported symbol.
func Load(path string, logger Logger) {
	if pluginsList, err := filepath.Glob(path); err == nil {
		for _, cursor := range pluginsList {
			if loadedPlugin, loadErr := load[ProviderPlugin](cursor); loadErr == nil {
				_ = RegisterPlugin(RegistryList.providers, &RegistryList.mu, loadedPlugin, func(p ProviderPlugin) string {
					logger.Info(fmt.Sprintf("Registering plugin: %s", p.Name()))
					return p.PackageName()
				})
			} else {
				logger.Error(loadErr, fmt.Sprintf("Failed to load plugin path=%s", cursor))
			}

			if loadedPlugin, loadErr := load[DeployerPlugin](cursor); loadErr == nil {
				_ = RegisterPlugin(RegistryList.deployers, &RegistryList.mu, loadedPlugin, func(p DeployerPlugin) string {
					logger.Info(fmt.Sprintf("Registering plugin: %s", p.Name()))
					return p.PackageName()
				})
			} else {
				logger.Error(loadErr, fmt.Sprintf("Failed to load plugin path=%s", cursor))
			}
		}
	}
}
