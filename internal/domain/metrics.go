package domain

// PodMetrics represents resource usage metrics for a pod
type PodMetrics struct {
	Name       string
	Namespace  string
	Containers []ContainerMetrics
	// Aggregated values for display
	CPUUsage    string // e.g., "100m" or "1.5"
	MemoryUsage string // e.g., "128Mi" or "1.2Gi"
}

// ContainerMetrics represents resource usage for a single container
type ContainerMetrics struct {
	Name        string
	CPUUsage    string
	MemoryUsage string
}
