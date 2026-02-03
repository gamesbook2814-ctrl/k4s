package k8s

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// GetPods returns all pods in the specified namespace
func (c *Client) GetPods(ctx context.Context, namespace string) ([]domain.Pod, error) {
	if namespace == "" {
		namespace = c.namespace
	}

	podList, err := c.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list pods: %w", err)
	}

	pods := make([]domain.Pod, 0, len(podList.Items))
	for _, p := range podList.Items {
		pods = append(pods, convertPod(&p))
	}
	return pods, nil
}

// GetPod returns a single pod by name with full details
func (c *Client) GetPod(ctx context.Context, namespace, name string) (*domain.Pod, error) {
	if namespace == "" {
		namespace = c.namespace
	}

	p, err := c.clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get pod %s: %w", name, err)
	}

	pod := convertPodDetailed(p)
	return &pod, nil
}

// GetPodEvents returns events for a specific pod
func (c *Client) GetPodEvents(ctx context.Context, namespace, podName string) ([]domain.PodEvent, error) {
	if namespace == "" {
		namespace = c.namespace
	}

	fieldSelector := fmt.Sprintf("involvedObject.name=%s,involvedObject.kind=Pod", podName)
	eventList, err := c.clientset.CoreV1().Events(namespace).List(ctx, metav1.ListOptions{
		FieldSelector: fieldSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("list events: %w", err)
	}

	events := make([]domain.PodEvent, 0, len(eventList.Items))
	for _, e := range eventList.Items {
		events = append(events, domain.PodEvent{
			Type:      e.Type,
			Reason:    e.Reason,
			Message:   e.Message,
			Count:     e.Count,
			FirstSeen: formatAge(e.FirstTimestamp.Time),
			LastSeen:  formatAge(e.LastTimestamp.Time),
		})
	}
	return events, nil
}

func convertPod(p *corev1.Pod) domain.Pod {
	// Calculate ready containers
	readyContainers := 0
	totalContainers := len(p.Spec.Containers)
	var totalRestarts int32

	containers := make([]domain.Container, 0, totalContainers)
	for _, c := range p.Spec.Containers {
		container := domain.Container{
			Name:  c.Name,
			Image: c.Image,
		}

		// Find container status
		for _, cs := range p.Status.ContainerStatuses {
			if cs.Name == c.Name {
				container.Ready = cs.Ready
				container.RestartCount = cs.RestartCount
				container.State = getContainerState(&cs)
				totalRestarts += cs.RestartCount
				if cs.Ready {
					readyContainers++
				}
				break
			}
		}

		containers = append(containers, container)
	}

	return domain.Pod{
		Name:       p.Name,
		Namespace:  p.Namespace,
		Ready:      fmt.Sprintf("%d/%d", readyContainers, totalContainers),
		Status:     getPodStatus(p),
		Restarts:   totalRestarts,
		Age:        formatAge(p.CreationTimestamp.Time),
		Node:       p.Spec.NodeName,
		Containers: containers,
	}
}

func convertPodDetailed(p *corev1.Pod) domain.Pod {
	// Calculate ready containers
	readyContainers := 0
	totalContainers := len(p.Spec.Containers)
	var totalRestarts int32

	containers := make([]domain.Container, 0, totalContainers)
	for _, c := range p.Spec.Containers {
		container := domain.Container{
			Name:  c.Name,
			Image: c.Image,
			Resources: domain.ContainerResources{
				CPURequest:    c.Resources.Requests.Cpu().String(),
				CPULimit:      c.Resources.Limits.Cpu().String(),
				MemoryRequest: c.Resources.Requests.Memory().String(),
				MemoryLimit:   c.Resources.Limits.Memory().String(),
			},
		}

		// Find container status
		for _, cs := range p.Status.ContainerStatuses {
			if cs.Name == c.Name {
				container.Ready = cs.Ready
				container.RestartCount = cs.RestartCount
				container.State, container.StateReason = getContainerStateDetailed(&cs)
				if cs.State.Running != nil {
					container.Started = formatAge(cs.State.Running.StartedAt.Time)
				}
				totalRestarts += cs.RestartCount
				if cs.Ready {
					readyContainers++
				}
				break
			}
		}

		containers = append(containers, container)
	}

	// Convert conditions
	conditions := make([]domain.PodCondition, 0, len(p.Status.Conditions))
	for _, c := range p.Status.Conditions {
		conditions = append(conditions, domain.PodCondition{
			Type:           string(c.Type),
			Status:         string(c.Status),
			Reason:         c.Reason,
			Message:        c.Message,
			LastTransition: formatAge(c.LastTransitionTime.Time),
		})
	}

	// Copy labels and annotations
	labels := make(map[string]string)
	for k, v := range p.Labels {
		labels[k] = v
	}

	annotations := make(map[string]string)
	for k, v := range p.Annotations {
		annotations[k] = v
	}

	return domain.Pod{
		Name:        p.Name,
		Namespace:   p.Namespace,
		Ready:       fmt.Sprintf("%d/%d", readyContainers, totalContainers),
		Status:      getPodStatus(p),
		Restarts:    totalRestarts,
		Age:         formatAge(p.CreationTimestamp.Time),
		Node:        p.Spec.NodeName,
		IP:          p.Status.PodIP,
		Labels:      labels,
		Annotations: annotations,
		Containers:  containers,
		Conditions:  conditions,
	}
}

func getPodStatus(p *corev1.Pod) string {
	// Check for deletion
	if p.DeletionTimestamp != nil {
		return "Terminating"
	}

	// Check init container statuses
	for _, cs := range p.Status.InitContainerStatuses {
		if cs.State.Waiting != nil && cs.State.Waiting.Reason != "" {
			return "Init:" + cs.State.Waiting.Reason
		}
		if cs.State.Terminated != nil && cs.State.Terminated.ExitCode != 0 {
			return "Init:Error"
		}
	}

	// Check container statuses
	for _, cs := range p.Status.ContainerStatuses {
		if cs.State.Waiting != nil && cs.State.Waiting.Reason != "" {
			return cs.State.Waiting.Reason
		}
		if cs.State.Terminated != nil {
			if cs.State.Terminated.Reason != "" {
				return cs.State.Terminated.Reason
			}
			if cs.State.Terminated.ExitCode != 0 {
				return fmt.Sprintf("Error:%d", cs.State.Terminated.ExitCode)
			}
		}
	}

	// Default to phase
	return string(p.Status.Phase)
}

func getContainerState(cs *corev1.ContainerStatus) string {
	if cs.State.Running != nil {
		return "Running"
	}
	if cs.State.Waiting != nil {
		if cs.State.Waiting.Reason != "" {
			return cs.State.Waiting.Reason
		}
		return "Waiting"
	}
	if cs.State.Terminated != nil {
		if cs.State.Terminated.Reason != "" {
			return cs.State.Terminated.Reason
		}
		return "Terminated"
	}
	return "Unknown"
}

func getContainerStateDetailed(cs *corev1.ContainerStatus) (state, reason string) {
	if cs.State.Running != nil {
		return "Running", ""
	}
	if cs.State.Waiting != nil {
		reason := cs.State.Waiting.Reason
		if reason == "" {
			reason = "Unknown"
		}
		return "Waiting", reason
	}
	if cs.State.Terminated != nil {
		reason := cs.State.Terminated.Reason
		if reason == "" {
			reason = fmt.Sprintf("ExitCode:%d", cs.State.Terminated.ExitCode)
		}
		return "Terminated", reason
	}
	return "Unknown", ""
}

// DeletePod deletes a pod by name
func (c *Client) DeletePod(ctx context.Context, namespace, name string) error {
	if namespace == "" {
		namespace = c.namespace
	}

	err := c.clientset.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("delete pod %s: %w", name, err)
	}

	return nil
}
