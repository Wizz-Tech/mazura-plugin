# Plugin SDK for Mazura (Go)

## Overview
This repository provides a small Go SDK for building and loading dynamically linked plugins (.so) for a host application. It defines:

- Core plugin interfaces (Plugin, ProviderPlugin, DeployerPlugin)
- A loader to discover and load compiled plugins from a directory
- A registry to keep track of loaded plugins
- Minimal Router and HTTPClient interfaces the host can provide to plugins
- Lightweight response types for common use-cases

The SDK is intended to be embedded in a host application. Third-party plugins implement these interfaces, are compiled as Go plugins with `-buildmode=plugin`, and are dynamically loaded by the host at runtime.

## Stack
- Language: Go
- Package manager/build tool: Go modules / `go` toolchain
- Module path: `github.com/Wizz-Tech/mazura-plugin`
- Notable stdlib feature: `plugin` package for dynamic loading
- CI: GitHub Actions workflow for auto-tagging releases (`.github/workflows/auto-tag.yml`)

## Requirements
- Go toolchain as declared in go.mod: `go 1.24.4`
  - The CI workflow currently sets up Go `1.22`. These differ. See TODOs to reconcile the supported Go version.
- Go plugins have platform and version constraints and are not supported on all platforms (notably Windows). Ensure your host and plugins are built with compatible Go versions and settings.

## Project Structure
```
/ (module: github.com/Wizz-Tech/mazura-plugin)
├─ go.mod
├─ v1/
│  ├─ plugin.go                 # Core Plugin interface & config schema
│  ├─ logger.go                 # Logger interface used by loader/host
│  ├─ loader.go                 # Loader that discovers and loads *.so plugins
│  ├─ registry.go               # In-memory registry for loaded plugins
│  ├─ deployerplugin.go         # Deployer-specific plugin interface
│  ├─ providerplugin.go         # Provider-specific plugin interface
│  ├─ responses/
│  │  ├─ deployment.go          # Response types for deployments
│  │  └─ release.go             # Response types for releases
│  └─ router/
│     ├─ router.go              # Minimal HTTP router abstraction
│     └─ httpclient.go          # Minimal HTTP client abstraction
└─ .github/workflows/auto-tag.yml
```

## Public APIs (v1)
- Plugin interface (v1/plugin.go):
  - PackageName() string
  - Name() string
  - Init(config map[string]string, r router.Router) error
  - Shutdown() error
  - Version() string
  - Description() string
  - GetConfig() []PluginConfigField
- Specialized plugin types:
  - ProviderPlugin: adds provider-oriented methods (OAuth capability, releases, deployments, PRs)
  - DeployerPlugin: adds deploy/inspect methods
- Loader (v1/loader.go):
  - Load(pluginDirectoryPath string, logger Logger)
    - Scans a directory for `*.so`, opens each with `plugin.Open`, expects an exported symbol named `Plugin` implementing the `Plugin` interface, and registers it.
- Registry (v1/registry.go):
  - InitRegistry(), RegisterPlugin(...), GetPlugin(...), GetPluginStore(...)
- Router/HTTP client abstractions (v1/router):
  - Router, Context, HandlerFunc and HTTPClient interfaces for host-provided HTTP handling.

## Installation (as a dependency)
Use the module path from go.mod:

- Add to your host application's go.mod by importing packages under `github.com/Wizz-Tech/mazura-plugin/v1/...`, then run:

```
go get github.com/Wizz-Tech/mazura-plugin/v1
```

If the canonical module path differs (e.g., repository renamed), update imports accordingly. See TODOs.

## Building a Plugin (.so)
Implement one of the interfaces in this module and export a package-level symbol named `Plugin` with the concrete implementation. Example skeleton:

```go
// yourplugin/main.go
package main

import (
    pluginv1 "github.com/Wizz-Tech/mazura-plugin/v1"
    "github.com/Wizz-Tech/mazura-plugin/v1/router"
)

type MyPlugin struct{}

func (p *MyPlugin) PackageName() string                 { return "your.plugin.package" }
func (p *MyPlugin) Name() string                        { return "YourPlugin" }
func (p *MyPlugin) Version() string                     { return "0.1.0" }
func (p *MyPlugin) Description() string                 { return "Example plugin" }
func (p *MyPlugin) GetConfig() []pluginv1.PluginConfigField { return nil }
func (p *MyPlugin) Init(cfg map[string]string, r router.Router) error { return nil }
func (p *MyPlugin) Shutdown() error                     { return nil }

// Exported symbol expected by loader: must implement pluginv1.Plugin
var Plugin pluginv1.Plugin = &MyPlugin{}
```

Build it as a Go plugin:

```
go build -buildmode=plugin -o yourplugin.so ./yourplugin
```

Important notes:
- The plugin must be built with a Go version and dependency graph compatible with the host application to avoid symbol/type mismatches.
- OS support for Go plugins is limited; ensure your target OS and architecture are supported by the Go `plugin` package.

## Loading Plugins in the Host
Example host usage:

```go
package main

import (
    "fmt"
    pluginv1 "github.com/Wizz-Tech/mazura-plugin/v1"
)

type stdLogger struct{}
func (l stdLogger) Debug(msg string)            { fmt.Println("DEBUG:", msg) }
func (l stdLogger) Info(msg string)             { fmt.Println("INFO:", msg) }
func (l stdLogger) Warn(msg string)             { fmt.Println("WARN:", msg) }
func (l stdLogger) Error(err error, msg string) { fmt.Println("ERROR:", msg, "=>", err) }

func main() {
    pluginv1.InitRegistry()
    pluginv1.Load("./plugins", stdLogger{})

    // Retrieve a plugin by package name
    if p, err := pluginv1.GetPlugin[pluginv1.Plugin](pluginv1.RegistryList, "your.plugin.package"); err == nil {
        fmt.Println("Loaded:", (*p).Name(), (*p).Version())
    }
}
```

The loader will scan the provided directory for `*.so` files, verify each exports a `Plugin` symbol that implements the `Plugin` interface, and register them by their `PackageName()`.

## Scripts
No repository-local scripts or Makefiles were found.
- CI: `.github/workflows/auto-tag.yml` automatically computes and pushes the next semantic version tag using `svu` when changes land on `main`.

## Environment Variables
No environment variables are defined in this repository. The host application may define its own variables that are passed to plugins through their Init config; document them in the host project.
- TODO: Document any expected environment variables (if any) used by the host when integrating this SDK.

## Tests
No test files were found in this repository.
- TODO: Add unit tests for loader, registry, and interfaces (mock-based).

## License
A LICENSE file is not present in this repository.
- TODO: Add a LICENSE (e.g., Apache-2.0, MIT, etc.) and include the appropriate notice here.

## Versioning and Releases
- Auto-tagging workflow uses `github.com/caarlos0/svu` to compute the next tag based on commit history.
- TODO: Confirm how releases are published/consumed and whether tags map to SDK API stability guarantees.

## Known Limitations
- Go plugins require compatible Go versions and are typically not supported on Windows.
- Plugins must export a `Plugin` symbol implementing the `v1.Plugin` interface; this is enforced by the loader via reflection.
- The go.mod specifies `go 1.24.4` but CI currently uses Go `1.22`.

## TODOs / Open Questions
- Reconcile Go version: go.mod (1.24.4) vs CI (1.22). Define the minimum supported Go version and update both.
- Add LICENSE file and update this README accordingly.
- Document any environment variables expected by the host application.
- Provide an end-to-end example repository: minimal host + sample plugin.
- Add tests, particularly for loader error paths and registry behavior.
