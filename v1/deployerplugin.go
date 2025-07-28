package plugin

type DeployerPluginCommands interface {
	Deploy(service string, image string) (string, error)
	GetRunningServices() ([]any, error)
}

type DeployerPlugin interface {
	Plugin
	DeployerPluginCommands
}
