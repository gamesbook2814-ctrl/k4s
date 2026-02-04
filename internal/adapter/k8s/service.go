package k8s

import (
	"context"
	"fmt"
	"maps"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// GetServices returns all services in the specified namespace
func (c *Client) GetServices(ctx context.Context, namespace string) ([]domain.Service, error) {
	if namespace == "" {
		namespace = c.namespace
	}

	svcList, err := c.clientset.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list services: %w", err)
	}

	services := make([]domain.Service, 0, len(svcList.Items))
	for _, s := range svcList.Items {
		services = append(services, convertService(&s))
	}
	return services, nil
}

// GetService returns a single service with full details
func (c *Client) GetService(ctx context.Context, namespace, name string) (*domain.Service, error) {
	if namespace == "" {
		namespace = c.namespace
	}

	s, err := c.clientset.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get service %s: %w", name, err)
	}

	service := convertServiceDetailed(s)
	return &service, nil
}

func convertService(s *corev1.Service) domain.Service {
	return domain.Service{
		Name:       s.Name,
		Namespace:  s.Namespace,
		Type:       string(s.Spec.Type),
		ClusterIP:  s.Spec.ClusterIP,
		ExternalIP: formatExternalIP(s),
		Ports:      formatPorts(s.Spec.Ports),
		Age:        formatAge(s.CreationTimestamp.Time),
	}
}

func convertServiceDetailed(s *corev1.Service) domain.Service {
	// Copy selector
	selector := make(map[string]string, len(s.Spec.Selector))
	maps.Copy(selector, s.Spec.Selector)

	// Copy labels
	labels := make(map[string]string, len(s.Labels))
	maps.Copy(labels, s.Labels)

	// Convert port details
	portDetails := make([]domain.ServicePort, 0, len(s.Spec.Ports))
	for _, p := range s.Spec.Ports {
		portDetails = append(portDetails, domain.ServicePort{
			Name:       p.Name,
			Port:       p.Port,
			TargetPort: p.TargetPort.String(),
			NodePort:   p.NodePort,
			Protocol:   string(p.Protocol),
		})
	}

	return domain.Service{
		Name:        s.Name,
		Namespace:   s.Namespace,
		Type:        string(s.Spec.Type),
		ClusterIP:   s.Spec.ClusterIP,
		ExternalIP:  formatExternalIP(s),
		Ports:       formatPorts(s.Spec.Ports),
		Age:         formatAge(s.CreationTimestamp.Time),
		Selector:    selector,
		Labels:      labels,
		PortDetails: portDetails,
	}
}

// formatPorts formats service ports for display
func formatPorts(ports []corev1.ServicePort) string {
	if len(ports) == 0 {
		return "<none>"
	}

	var parts []string
	for _, p := range ports {
		portStr := fmt.Sprintf("%d/%s", p.Port, p.Protocol)
		if p.NodePort > 0 {
			portStr = fmt.Sprintf("%d:%d/%s", p.Port, p.NodePort, p.Protocol)
		}
		parts = append(parts, portStr)
	}
	return strings.Join(parts, ",")
}

// formatExternalIP returns external IP or status
func formatExternalIP(s *corev1.Service) string {
	switch s.Spec.Type {
	case corev1.ServiceTypeLoadBalancer:
		// Check ingress IPs
		if len(s.Status.LoadBalancer.Ingress) > 0 {
			var ips []string
			for _, ing := range s.Status.LoadBalancer.Ingress {
				if ing.IP != "" {
					ips = append(ips, ing.IP)
				} else if ing.Hostname != "" {
					ips = append(ips, ing.Hostname)
				}
			}
			if len(ips) > 0 {
				return strings.Join(ips, ",")
			}
		}
		return "<pending>"

	case corev1.ServiceTypeNodePort:
		if len(s.Spec.ExternalIPs) > 0 {
			return strings.Join(s.Spec.ExternalIPs, ",")
		}
		return "<none>"

	case corev1.ServiceTypeExternalName:
		return s.Spec.ExternalName

	default:
		if len(s.Spec.ExternalIPs) > 0 {
			return strings.Join(s.Spec.ExternalIPs, ",")
		}
		return "<none>"
	}
}
