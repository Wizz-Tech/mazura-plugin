package responses

type GetAffectedDeploymentsResponse struct {
	Service       string `json:"service"`
	CreatedAtTime string `json:"createdAt"`
}
