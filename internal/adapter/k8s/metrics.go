package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsclient "k8s.io/metrics/pkg/client/clientset/versioned"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// MetricsClient wraps the Kubernetes metrics API client
type MetricsClient struct {
	client *metricsclient.Clientset
}

// NewMetricsClient creates a new metrics client from the kubeconfig path
func NewMetricsClient(kubeconfigPath string) (*MetricsClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("build config for metrics: %w", err)
	}

	client, err := metricsclient.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create metrics client: %w", err)
	}

	return &MetricsClient{client: client}, nil
}

// CheckMetricsAvailable tests if the metrics server is available
func (m *MetricsClient) CheckMetricsAvailable(ctx context.Context) bool {
	if m == nil || m.client == nil {
		return false
	}

	// Try to list node metrics as a quick availability check
	_, err := m.client.MetricsV1beta1().NodeMetricses().List(ctx, metav1.ListOptions{Limit: 1})
	return err == nil
}

// GetPodMetrics returns metrics for all pods in the namespace
func (m *MetricsClient) GetPodMetrics(ctx context.Context, namespace string) (map[string]domain.PodMetrics, error) {
	if m == nil || m.client == nil {
		return nil, fmt.Errorf("metrics client not available")
	}

	podMetricsList, err := m.client.MetricsV1beta1().PodMetricses(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list pod metrics: %w", err)
	}

	result := make(map[string]domain.PodMetrics, len(podMetricsList.Items))
	for _, pm := range podMetricsList.Items {
		result[pm.Name] = convertPodMetrics(&pm)
	}
	return result, nil
}

// GetPodMetric returns metrics for a single pod
func (m *MetricsClient) GetPodMetric(ctx context.Context, namespace, name string) (*domain.PodMetrics, error) {
	if m == nil || m.client == nil {
		return nil, fmt.Errorf("metrics client not available")
	}

	pm, err := m.client.MetricsV1beta1().PodMetricses(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get pod metrics %s: %w", name, err)
	}

	metrics := convertPodMetrics(pm)
	return &metrics, nil
}

func convertPodMetrics(pm *metricsv1beta1.PodMetrics) domain.PodMetrics {
	containers := make([]domain.ContainerMetrics, 0, len(pm.Containers))

	var totalCPU int64
	var totalMemory int64

	for _, c := range pm.Containers {
		cpuQuantity := c.Usage.Cpu()
		memQuantity := c.Usage.Memory()

		totalCPU += cpuQuantity.MilliValue()
		totalMemory += memQuantity.Value()

		containers = append(containers, domain.ContainerMetrics{
			Name:        c.Name,
			CPUUsage:    formatCPU(cpuQuantity.MilliValue()),
			MemoryUsage: formatMemory(memQuantity.Value()),
		})
	}

	return domain.PodMetrics{
		Name:        pm.Name,
		Namespace:   pm.Namespace,
		Containers:  containers,
		CPUUsage:    formatCPU(totalCPU),
		MemoryUsage: formatMemory(totalMemory),
	}
}

// formatCPU formats CPU millicores for display
func formatCPU(milliCores int64) string {
	if milliCores < 1000 {
		return fmt.Sprintf("%dm", milliCores)
	}
	return fmt.Sprintf("%.1f", float64(milliCores)/1000)
}

// formatMemory formats memory bytes for display
func formatMemory(bytes int64) string {
	const (
		Ki = 1024
		Mi = Ki * 1024
		Gi = Mi * 1024
	)

	switch {
	case bytes >= Gi:
		return fmt.Sprintf("%.1fGi", float64(bytes)/float64(Gi))
	case bytes >= Mi:
		return fmt.Sprintf("%.0fMi", float64(bytes)/float64(Mi))
	case bytes >= Ki:
		return fmt.Sprintf("%.0fKi", float64(bytes)/float64(Ki))
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}
