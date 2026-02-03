package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

const (
	configDir  = ".k4s"
	configFile = "config.yaml"
)

// Loader handles configuration loading and management
type Loader struct {
	configPath string
	viper      *viper.Viper
}

// NewLoader creates a new config loader
func NewLoader() *Loader {
	return &Loader{
		viper: viper.New(),
	}
}

// Load reads the configuration from ~/.k4s/config.yaml
func (l *Loader) Load() (*domain.Config, error) {
	configPath, err := l.ensureConfigDir()
	if err != nil {
		return nil, fmt.Errorf("ensure config directory: %w", err)
	}
	l.configPath = configPath

	l.viper.SetConfigFile(filepath.Join(configPath, configFile))
	l.viper.SetConfigType("yaml")

	if err := l.viper.ReadInConfig(); err != nil {
		if os.IsNotExist(err) {
			return l.createDefault()
		}
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg domain.Config
	if err := l.viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}

// ConfigPath returns the configuration directory path
func (l *Loader) ConfigPath() string {
	return l.configPath
}

func (l *Loader) ensureConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home directory: %w", err)
	}

	configPath := filepath.Join(home, configDir)
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return "", fmt.Errorf("create config directory: %w", err)
	}

	return configPath, nil
}

func (l *Loader) createDefault() (*domain.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("get home directory: %w", err)
	}

	defaultKubePath := filepath.Join(home, ".kube", "config")

	cfg := &domain.Config{
		KubeConfigs: []domain.KubeConfig{
			{
				Name:    "default",
				Path:    defaultKubePath,
				Default: true,
			},
		},
		SSHHosts: []domain.SSHHost{},
	}

	if err := l.Save(cfg); err != nil {
		return nil, fmt.Errorf("save default config: %w", err)
	}

	return cfg, nil
}

// Save writes the configuration to disk
func (l *Loader) Save(cfg *domain.Config) error {
	l.viper.Set("kubeconfigs", cfg.KubeConfigs)
	l.viper.Set("ssh_hosts", cfg.SSHHosts)

	configFilePath := filepath.Join(l.configPath, configFile)
	if err := l.viper.WriteConfigAs(configFilePath); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}
