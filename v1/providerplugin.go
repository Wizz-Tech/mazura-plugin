package plugin

// GetReleaseResponse represents the response containing details about a release.
type GetReleaseResponse struct {
	// Name represents the name of the release.
	Name string `json:"name"`

	// CreatedAt represents the date the resource was created.
	//
	// It is in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ).
	//
	CreatedAt string `json:"createdAt"`

	// Draft        bool   `json:"draft"`
	Draft bool `json:"draft"`

	// Repository represents a repository identifier.
	//
	// It is used to store the identifier for the release repository.
	//
	// The value is converted to lowercase during JSON serialization.
	Repository string `json:"repository"`
	//
	URL string `json:"url"`

	// TargetBranch represents the branch to which the release will be deployed.
	//
	// It is optional during JSON deserialization.
	TargetBranch string `json:"targetBranch,omitempty"`

	//
	Notes string `json:"notes,omitempty"`
}

// OauthCallbackContext represents the context data for handling OAuth callback requests.
// It provides the authorization code and HTTP client for processing the callback.
type OauthCallbackContext struct {
	//
	// Represents a unique identifier for an OAuth callback request.
	//
	Code string

	// Represents a client used to make HTTP requests.
	//
	// It provides methods for sending HTTP requests and receiving responses.
	//
	// Uses the HTTP protocol version specified in the `router.HTTPClient`.
	//
	// This client can be configured with additional headers, query parameters,
	// and other options to customize the request behavior.
	// Sets the base URL of the HTTP client.
	HTTPClient HTTPClient
}

// ProviderPluginCommands is the interface for provider plugin commands.
type ProviderPluginCommands interface {
	// IsOauthCapable describes whether the provider supports OAuth.
	// Returns true if the provider supports OAuth, false otherwise.
	IsOauthCapable() bool

	// Generates the OAuth connection URL based on the provided data.
	//
	// @param data Map of string parameters to generate the connection URL.
	//
	// @return The generated OAuth connection URL as a string.
	// @return An error if generating the connection URL fails.
	GenerateOauthConnectionURL(data map[string]string) (OauthURL string, err error)

	// Handles the OAuth callback by processing the provided data and updating the context.
	//
	// @param c The OAuth callback context containing the code from the request.
	//
	// @return A map of strings representing the result of the callback operation,
	//         or an error if the operation fails.
	HandleOauthCallback(c *OauthCallbackContext) (callbackResult map[string]string, err error)

	// Returns a list of releases, along with any associated metadata.
	ListReleases() ([]any, error)

	// Retrieves the release by tag from the repository.
	//
	GetReleaseByTag(tag string) (any, error)

	// CreateDeployment creates a new deployment.
	// @param ref the reference to the deployment to be created
	// @return any the newly created deployment or an error if creation fails.
	CreateDeployment(ref string) (any, error)

	// id: The unique identifier of the deployment to update.
	// state: The new status of the deployment.
	// envURL: The URL of the environment where the deployment is running.
	// Returns: No return value.
	// Errors: An error if the deployment could not be updated.
	UpdateDeploymentStatus(id int64, state string, envURL string) (any, error)

	//
	// GetPullRequest returns a pull request by its number.
	//
	// @param number The number of the pull request to get.
	//
	// @return The pull request, or an error if it cannot be found.
	//
	// @ Panics If the provided number is invalid.
	//
	// @ See ProviderPluginCommands
	//
	GetPullRequest(number int) (any, error)
}

// ProviderPlugin defines methods for Git providers like GitHub, GitLab, etc.
type ProviderPlugin interface {
	Plugin
	ProviderPluginCommands
}
