package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config represents the application configuration.
type Config struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	GrpcPort string `json:"grpc_port"`
	Version  string `json:"version"`
}

// NewConfig creates a new Config.
func NewConfig() *Config {
	return &Config{
		Name:    "ova-account-api",
		Address: ":8080",
		Version: "v1",
	}
}

// ParseConfig parses the configuration from the file at the specified path.
func ParseConfig(path string) (*Config, error) {
	conf := NewConfig()
	err := conf.Update(path)

	return conf, err
}

// Update updates the configuration from the file at the specified path.
func (c *Config) Update(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}
