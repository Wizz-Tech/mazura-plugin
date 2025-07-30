package plugin

type DeployerPluginCommands interface {
	Deploy(service string, image string, statusCh chan<- string)
	GetRunningServices() ([]any, error)
}

type DeployerPlugin interface {
	Plugin
	DeployerPluginCommands
}
