package plugin

// ProviderPlugin defines methods for Git providers like GitHub, GitLab, etc.
type ProviderPlugin interface {
	Plugin
	IsOauthCapable() bool
	GenerateOauthConnectionURL() string
	ListReleases() ([]any, error)
	GetReleaseByTag(tag string) (any, error)
	CreateDeployment(ref string) (any, error)
	UpdateDeploymentStatus(id int64, state string, envURL string) (any, error)
	GetPullRequest(number int) (any, error)
}
