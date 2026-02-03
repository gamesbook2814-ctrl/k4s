package domain

// Pod represents a Kubernetes pod
type Pod struct {
	Name        string
	Namespace   string
	Ready       string // e.g., "1/1", "0/2"
	Status      string // Running, Pending, Failed, etc.
	Restarts    int32
	Age         string
	Node        string
	IP          string
	Labels      map[string]string
	Annotations map[string]string
	Containers  []Container
	Conditions  []PodCondition
}

// Container represents a container within a pod
type Container struct {
	Name         string
	Image        string
	Ready        bool
	RestartCount int32
	State        string
	StateReason  string
	Started      string
	Resources    ContainerResources
}

// ContainerResources represents resource requests and limits
type ContainerResources struct {
	CPURequest    string
	CPULimit      string
	MemoryRequest string
	MemoryLimit   string
}

// PodCondition represents a pod condition
type PodCondition struct {
	Type    string
	Status  string
	Reason  string
	Message string
	LastTransition string
}

// PodEvent represents an event related to a pod
type PodEvent struct {
	Type      string
	Reason    string
	Message   string
	Count     int32
	FirstSeen string
	LastSeen  string
}

// PodStatus constants
const (
	PodStatusRunning   = "Running"
	PodStatusPending   = "Pending"
	PodStatusSucceeded = "Succeeded"
	PodStatusFailed    = "Failed"
	PodStatusUnknown   = "Unknown"
)
