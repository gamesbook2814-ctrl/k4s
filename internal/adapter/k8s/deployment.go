package k8s

import (
	"context"
	"fmt"
	"maps"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// GetDeployments returns all deployments in the specified namespace
func (c *Client) GetDeployments(ctx context.Context, namespace string) ([]domain.Deployment, error) {
	if namespace == "" {
		namespace = c.namespace
	}

	depList, err := c.clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list deployments: %w", err)
	}

	deployments := make([]domain.Deployment, 0, len(depList.Items))
	for _, d := range depList.Items {
		deployments = append(deployments, convertDeployment(&d))
	}
	return deployments, nil
}

// GetDeployment returns a single deployment with full details
func (c *Client) GetDeployment(ctx context.Context, namespace, name string) (*domain.Deployment, error) {
	if namespace == "" {
		namespace = c.namespace
	}

	d, err := c.clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get deployment %s: %w", name, err)
	}

	deployment := convertDeploymentDetailed(d)
	return &deployment, nil
}

// ScaleDeployment scales a deployment to the specified number of replicas
func (c *Client) ScaleDeployment(ctx context.Context, namespace, name string, replicas int32) error {
	if namespace == "" {
		namespace = c.namespace
	}

	scale, err := c.clientset.AppsV1().Deployments(namespace).GetScale(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("get deployment scale %s: %w", name, err)
	}

	scale.Spec.Replicas = replicas
	_, err = c.clientset.AppsV1().Deployments(namespace).UpdateScale(ctx, name, scale, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("scale deployment %s to %d: %w", name, replicas, err)
	}

	return nil
}

// RestartDeployment triggers a rolling restart by patching the deployment's annotations
func (c *Client) RestartDeployment(ctx context.Context, namespace, name string) error {
	if namespace == "" {
		namespace = c.namespace
	}

	// Use strategic merge patch to update the restartedAt annotation
	patch := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"kubectl.kubernetes.io/restartedAt":"%s"}}}}}`,
		time.Now().Format(time.RFC3339))

	_, err := c.clientset.AppsV1().Deployments(namespace).Patch(
		ctx,
		name,
		types.StrategicMergePatchType,
		[]byte(patch),
		metav1.PatchOptions{},
	)
	if err != nil {
		return fmt.Errorf("restart deployment %s: %w", name, err)
	}

	return nil
}

// DeleteDeployment deletes a deployment by name
func (c *Client) DeleteDeployment(ctx context.Context, namespace, name string) error {
	if namespace == "" {
		namespace = c.namespace
	}

	err := c.clientset.AppsV1().Deployments(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("delete deployment %s: %w", name, err)
	}

	return nil
}

func convertDeployment(d *appsv1.Deployment) domain.Deployment {
	replicas := int32(1)
	if d.Spec.Replicas != nil {
		replicas = *d.Spec.Replicas
	}

	// Collect container images
	images := make([]string, 0, len(d.Spec.Template.Spec.Containers))
	for _, c := range d.Spec.Template.Spec.Containers {
		images = append(images, c.Image)
	}

	return domain.Deployment{
		Name:          d.Name,
		Namespace:     d.Namespace,
		Ready:         fmt.Sprintf("%d/%d", d.Status.ReadyReplicas, replicas),
		UpToDate:      d.Status.UpdatedReplicas,
		Available:     d.Status.AvailableReplicas,
		Age:           formatAge(d.CreationTimestamp.Time),
		Replicas:      replicas,
		ReadyReplicas: d.Status.ReadyReplicas,
		Images:        images,
	}
}

func convertDeploymentDetailed(d *appsv1.Deployment) domain.Deployment {
	replicas := int32(1)
	if d.Spec.Replicas != nil {
		replicas = *d.Spec.Replicas
	}

	// Copy labels
	labels := make(map[string]string, len(d.Labels))
	maps.Copy(labels, d.Labels)

	// Copy selector
	selector := make(map[string]string)
	if d.Spec.Selector != nil {
		maps.Copy(selector, d.Spec.Selector.MatchLabels)
	}

	// Convert conditions
	conditions := make([]domain.DeploymentCondition, 0, len(d.Status.Conditions))
	for _, c := range d.Status.Conditions {
		conditions = append(conditions, domain.DeploymentCondition{
			Type:           string(c.Type),
			Status:         string(c.Status),
			Reason:         c.Reason,
			Message:        c.Message,
			LastTransition: formatAge(c.LastTransitionTime.Time),
		})
	}

	// Collect container images
	images := make([]string, 0, len(d.Spec.Template.Spec.Containers))
	for _, c := range d.Spec.Template.Spec.Containers {
		images = append(images, c.Image)
	}

	// Get strategy type
	strategy := string(d.Spec.Strategy.Type)

	return domain.Deployment{
		Name:          d.Name,
		Namespace:     d.Namespace,
		Ready:         fmt.Sprintf("%d/%d", d.Status.ReadyReplicas, replicas),
		UpToDate:      d.Status.UpdatedReplicas,
		Available:     d.Status.AvailableReplicas,
		Age:           formatAge(d.CreationTimestamp.Time),
		Replicas:      replicas,
		ReadyReplicas: d.Status.ReadyReplicas,
		Strategy:      strategy,
		Labels:        labels,
		Selector:      selector,
		Conditions:    conditions,
		Images:        images,
	}
}
