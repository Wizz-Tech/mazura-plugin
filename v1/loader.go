package plugin

import (
	"fmt"
	"log"
	"os"
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
func Load(pluginDirectoryPath string, logger Logger) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err.(error), "Failed to load plugins")
		}
	}()

	directories, err := os.ReadDir(pluginDirectoryPath)
	if err != nil {
		log.Panic(err)
	}

	for _, v := range directories {
		if v.IsDir() {
			continue
		}
		if filepath.Ext(v.Name()) != ".so" {
			continue
		}

		foundPlugin := filepath.Join(pluginDirectoryPath, v.Name())
		if loadedPlugin, loadErr := load[Plugin](foundPlugin); loadErr == nil {
			_ = RegisterPlugin(RegistryList.plugins, &RegistryList.mu, loadedPlugin, func(p Plugin) string {
				logger.Info(fmt.Sprintf("Registering plugin: %s", p.Name()))
				return p.PackageName()
			})
		} else {
			logger.Error(loadErr, fmt.Sprintf("Failed to load plugin path=%s", foundPlugin))
		}
	}
}
