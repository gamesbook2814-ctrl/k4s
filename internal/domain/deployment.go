package domain

// Deployment represents a Kubernetes Deployment
type Deployment struct {
	Name          string
	Namespace     string
	Ready         string // e.g., "3/3" (ready/desired replicas)
	UpToDate      int32
	Available     int32
	Age           string
	Replicas      int32
	ReadyReplicas int32
	Strategy      string
	Labels        map[string]string
	Selector      map[string]string
	Conditions    []DeploymentCondition
	Images        []string
}

// DeploymentCondition represents a condition of a deployment
type DeploymentCondition struct {
	Type           string
	Status         string
	Reason         string
	Message        string
	LastTransition string
}

// DeploymentStatus constants
const (
	DeploymentStatusProgressing = "Progressing"
	DeploymentStatusAvailable   = "Available"
	DeploymentStatusFailed      = "Failed"
)
