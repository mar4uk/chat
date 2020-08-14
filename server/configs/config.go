package configs

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// DatabaseConfig is used to describe database configuration
type DatabaseConfig struct {
	Protocol string `yaml:"protocol" envconfig:"DB_PROTOCOL"`
	Host     string `yaml:"host" envconfig:"DB_HOST"`
	Port     string `yaml:"port" envconfig:"DB_PORT"`
	Username string `yaml:"username" envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `yaml:"db-name" envconfig:"DB_NAME"`
}

// ServerConfig describes host and port where app is running
type ServerConfig struct {
	Host string `yaml:"host" envconfig:"HOST"`
	Port string `yaml:"port" envconfig:"PORT"`
}

// Config is configuration of app
type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
}

// ParseConfig reads config file and returns Config
func ParseConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	err = envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
