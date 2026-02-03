package k8s

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// Client wraps the Kubernetes clientset
type Client struct {
	clientset  *kubernetes.Clientset
	config     *clientcmd.ClientConfig
	rawConfig  clientcmd.ClientConfig
	kubeconfig string
	context    string
	namespace  string
}

// NewClient creates a new Kubernetes client from a kubeconfig path
func NewClient(kubeconfigPath string) (*Client, error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{
		ExplicitPath: kubeconfigPath,
	}

	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides,
	)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("build client config: %w", err)
	}

	// Set reasonable timeouts
	config.Timeout = 10 * time.Second

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create clientset: %w", err)
	}

	// Get current context and namespace
	rawConfig, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, fmt.Errorf("get raw config: %w", err)
	}

	currentContext := rawConfig.CurrentContext
	namespace := "default"
	if ctx, ok := rawConfig.Contexts[currentContext]; ok && ctx.Namespace != "" {
		namespace = ctx.Namespace
	}

	return &Client{
		clientset:  clientset,
		kubeconfig: kubeconfigPath,
		context:    currentContext,
		namespace:  namespace,
	}, nil
}

// CheckConnection verifies the connection to the cluster
func (c *Client) CheckConnection(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := c.clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{Limit: 1})
	if err != nil {
		return fmt.Errorf("connection check failed: %w", err)
	}
	return nil
}

// GetClusterInfo returns information about the connected cluster
func (c *Client) GetClusterInfo(ctx context.Context) (*domain.ClusterInfo, error) {
	info := &domain.ClusterInfo{
		Context:   c.context,
		Namespace: c.namespace,
		Connected: false,
	}

	// Check connection and get server version for cluster name
	version, err := c.clientset.Discovery().ServerVersion()
	if err != nil {
		return info, fmt.Errorf("get server version: %w", err)
	}

	info.Connected = true
	info.Name = fmt.Sprintf("Kubernetes %s", version.GitVersion)

	return info, nil
}

// CurrentNamespace returns the current namespace
func (c *Client) CurrentNamespace() string {
	return c.namespace
}

// SetNamespace sets the current namespace
func (c *Client) SetNamespace(ns string) {
	c.namespace = ns
}

// CurrentContext returns the current context name
func (c *Client) CurrentContext() string {
	return c.context
}

// Clientset returns the underlying Kubernetes clientset
func (c *Client) Clientset() *kubernetes.Clientset {
	return c.clientset
}

// ListNamespaces returns all namespaces in the cluster
func (c *Client) ListNamespaces(ctx context.Context) ([]string, error) {
	nsList, err := c.clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list namespaces: %w", err)
	}

	namespaces := make([]string, 0, len(nsList.Items))
	for _, ns := range nsList.Items {
		namespaces = append(namespaces, ns.Name)
	}
	return namespaces, nil
}
