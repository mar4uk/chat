package configs

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// MongoConfig is used to describe database configuration
type MongoConfig struct {
	Protocol string `yaml:"protocol" envconfig:"MONGO_PROTOCOL"`
	Host     string `yaml:"host" envconfig:"MONGO_HOST"`
	Port     string `yaml:"port" envconfig:"MONGO_PORT"`
	Username string `yaml:"username" envconfig:"MONGO_USER"`
	Password string `envconfig:"MONGO_PASSWORD"`
	Name     string `yaml:"db-name" envconfig:"MONGO_NAME"`
}

// ServerConfig describes host and port where app is running
type ServerConfig struct {
	Host string `yaml:"host" envconfig:"HOST"`
	Port uint16 `yaml:"port" envconfig:"PORT"`
}

// Config is configuration of app
type Config struct {
	Server ServerConfig `yaml:"server"`
	Mongo  MongoConfig  `yaml:"mongo"`
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
