package responses

type GetReleaseResponse struct {
	Name         string `json:"name"`
	CreatedAt    string `json:"createdAt"`
	Draft        bool   `json:"draft"`
	Repository   string `json:"repository"`
	URL          string `json:"url"`
	TargetBranch string `json:"targetBranch,omitempty"`
	Notes        string `json:"notes,omitempty"`
}
