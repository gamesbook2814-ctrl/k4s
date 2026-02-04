package domain

// Service represents a Kubernetes Service
type Service struct {
	Name       string
	Namespace  string
	Type       string // ClusterIP, NodePort, LoadBalancer, ExternalName
	ClusterIP  string
	ExternalIP string // "<none>" or actual IP/hostname
	Ports      string // formatted: "80/TCP,443/TCP"
	Age        string
	Selector   map[string]string
	Labels     map[string]string
	PortDetails []ServicePort
}

// ServicePort represents a port exposed by a service
type ServicePort struct {
	Name       string
	Port       int32
	TargetPort string
	NodePort   int32
	Protocol   string
}

// ServiceType constants
const (
	ServiceTypeClusterIP    = "ClusterIP"
	ServiceTypeNodePort     = "NodePort"
	ServiceTypeLoadBalancer = "LoadBalancer"
	ServiceTypeExternalName = "ExternalName"
)
