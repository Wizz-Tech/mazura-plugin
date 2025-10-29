package plugin

// GetAffectedDeploymentsResponse represents a deployment affected by a revision.
type GetAffectedDeploymentsResponse struct {
	// Service represents the identifier of the affected service in the deployment response.
	Service string `json:"service"`
	// CreatedAtTime indicates the timestamp when the affected deployment was created, formatted as a string.
	CreatedAtTime string `json:"createdAt"`
}

// DeployContext defines the context for a deployment.
type DeployContext struct {
	// Name of the deployment.
	Service string
	// Image of the deployment.
	Image string
	// SetStatus sets the status of the deployment using the provided DeploymentStatus value.
	SetStatus func(status DeploymentStatus)
	// SetRevisionURL sets the revision URL of the deployment.
	RevisionURL *string
}

// DeploymentStatus defines the status of a deployment.
type DeploymentStatus string

const (
	// DeploymentStatusPending indicates that the deployment is awaiting execution or processing.
	DeploymentStatusPending DeploymentStatus = "pending"
	// DeploymentStatusSuccess indicates that the deployment was successful.
	DeploymentStatusSuccess DeploymentStatus = "success"
	// DeploymentStatusInProgress indicates that the deployment is in progress.
	DeploymentStatusInProgress DeploymentStatus = "progress"
	// DeploymentStatusFailed indicates that the deployment failed.
	DeploymentStatusFailed DeploymentStatus = "failed"
)

// DeployerPluginCommands defines methods for deployer plugins.
type DeployerPluginCommands interface {
	// Deploy performs the deployment of a service based on the provided DeployContext and returns an error if it fails.
	Deploy(ctx *DeployContext) error
	// GetRunningServices retrieves a list of currently running services and returns an error if the operation fails.
	GetRunningServices() ([]any, error)
	// GetAffectedRevisions retrieves a list of affected deployments for the given service name and returns an error if it fails.
	GetAffectedRevisions(name string) (result []GetAffectedDeploymentsResponse, err error)
}

// DeployerPlugin definition for deployer plugins.
type DeployerPlugin interface {
	Plugin
	DeployerPluginCommands
}
