package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config represents the application configuration.
type Config struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	GrpcPort string   `json:"grpc_port"`
	Version  string   `json:"version"`
	DB       DBConfig `json:"db"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// DSN return the data source name of DB.
func (c *DBConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
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
