package domain

import "errors"

// Sentinel errors for configuration
var (
	ErrConfigNotFound    = errors.New("configuration file not found")
	ErrNoKubeConfigs     = errors.New("no kubeconfigs configured")
	ErrKubeConfigInvalid = errors.New("kubeconfig is invalid")
)
