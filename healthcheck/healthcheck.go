package healthcheck

// HealthCheck - status
type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}
