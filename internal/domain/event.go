package domain

// Event represents a Kubernetes Event
type Event struct {
	Name            string
	Namespace       string
	Type            string // Normal, Warning
	Reason          string
	Message         string
	Object          string // formatted: "Pod/my-pod", "Deployment/my-deploy"
	ObjectKind      string // Pod, Deployment, Service, etc.
	ObjectName      string
	Count           int32
	FirstSeen       string
	LastSeen        string
	Age             string
	SourceComponent string
}

// EventType constants
const (
	EventTypeNormal  = "Normal"
	EventTypeWarning = "Warning"
)
