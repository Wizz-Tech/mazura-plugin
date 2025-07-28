package plugin

type DeployerPluginCommands interface {
	Deploy(service string, image string) error
	GetRunningServices() ([]any, error)
}

type DeployerPlugin interface {
	Plugin
	DeployerPluginCommands
}
