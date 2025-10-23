package plugin

import (
	"github.com/Wizz-Tech/mazura-plugin/v1/responses"
)

type DeployContext struct {
	Service     string
	Image       string
	SetStatus   func(status DeploymentStatus)
	RevisionURL *string
}

type DeploymentStatus string

const (
	DeploymentStatusPending    DeploymentStatus = "pending"
	DeploymentStatusSuccess    DeploymentStatus = "success"
	DeploymentStatusInProgress DeploymentStatus = "progress"
	DeploymentStatusFailed     DeploymentStatus = "failed"
)

type DeployerPluginCommands interface {
	Deploy(ctx *DeployContext) error
	GetRunningServices() ([]any, error)
	GetAffectedRevisions(name string) (result []responses.GetAffectedDeploymentsResponse, err error)
}

type DeployerPlugin interface {
	Plugin
	DeployerPluginCommands
}
