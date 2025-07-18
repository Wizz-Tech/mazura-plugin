package plugin

type DeployerPluginCommands interface {
	Deploy(ref string) error
	GetRunningServices() ([]any, error)
}

type DeployerPlugin interface {
	Plugin
	DeployerPluginCommands
}
