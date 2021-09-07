package health

type Status string

const (
	Ok Status = "OK"
	Error Status = "ERROR"
)

type Resource struct {
	Name    string `json:"name"`
	Status  Status `json:"status"`
	Message string `json:"message"`
}

// Check checks the status of resources.
type Check func() []Resource
