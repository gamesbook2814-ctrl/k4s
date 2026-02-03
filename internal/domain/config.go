package domain

// KubeConfig represents a kubeconfig entry
type KubeConfig struct {
	Name    string `yaml:"name" mapstructure:"name"`
	Path    string `yaml:"path" mapstructure:"path"`
	Default bool   `yaml:"default" mapstructure:"default"`
}

// SSHHost represents an SSH host configuration
type SSHHost struct {
	Name    string `yaml:"name" mapstructure:"name"`
	Host    string `yaml:"host" mapstructure:"host"`
	User    string `yaml:"user" mapstructure:"user"`
	KeyPath string `yaml:"key_path" mapstructure:"key_path"`
	Port    int    `yaml:"port" mapstructure:"port"`
}

// Config represents the application configuration
type Config struct {
	KubeConfigs []KubeConfig `yaml:"kubeconfigs" mapstructure:"kubeconfigs"`
	SSHHosts    []SSHHost    `yaml:"ssh_hosts" mapstructure:"ssh_hosts"`
}

// DefaultKubeConfig returns the default kubeconfig or the first one
func (c *Config) DefaultKubeConfig() *KubeConfig {
	for i := range c.KubeConfigs {
		if c.KubeConfigs[i].Default {
			return &c.KubeConfigs[i]
		}
	}
	if len(c.KubeConfigs) > 0 {
		return &c.KubeConfigs[0]
	}
	return nil
}

// FindKubeConfig finds a kubeconfig by name
func (c *Config) FindKubeConfig(name string) *KubeConfig {
	for i := range c.KubeConfigs {
		if c.KubeConfigs[i].Name == name {
			return &c.KubeConfigs[i]
		}
	}
	return nil
}

// NodeInfo represents information about a node
type NodeInfo struct {
	Hostname string
	OS       string
	Kernel   string
	Uptime   string
	Memory   string
	Disk     string
	LoadAvg  string
}
