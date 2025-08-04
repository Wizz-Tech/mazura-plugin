package plugin

type DeployContext struct {
	Service   string
	Image     string
	SetStatus func(status string)
}

type DeployerPluginCommands interface {
	Deploy(ctx *DeployContext) error
	GetRunningServices() ([]any, error)
}

type DeployerPlugin interface {
	Plugin
	DeployerPluginCommands
}
