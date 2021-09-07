package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	Name       string      `mapstructure:"SERVICE"`
	GrpcPort   string      `mapstructure:"GRPC_PORT"`
	HealthPort string      `mapstructure:"HEALTH_PORT"`
	Debug      bool        `mapstructure:"DEBUG"`
	DB         DBConfig    `mapstructure:",squash"`
	Kafka      KafkaConfig `mapstructure:",squash"`
}

type DBConfig struct {
	Host      string       `mapstructure:"DB_HOST"`
	Port      string       `mapstructure:"DB_PORT"`
	Database  string       `mapstructure:"DB_DATABASE"`
	User      string       `mapstructure:"DB_USERNAME"`
	Password  string       `mapstructure:"DB_PASSWORD"`
	BatchSize int          `mapstructure:"DB_BATCH_SIZE"`
	Pool      DBPoolConfig `mapstructure:",squash"`
}

type DBPoolConfig struct {
	MaxOpenConns    int `mapstructure:"DB_MAX_OPEN_CONNS"`
	MaxIdleConn     int `mapstructure:"DB_MAX_IDLE_CONNS"`
	ConnMaxLifetime int `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

// DSN return the data source name of DB.
func (c *DBConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
}

type KafkaConfig struct {
	Addr  string `mapstructure:"KAFKA_ADDR"`
	Topic string `mapstructure:"KAFKA_TOPIC"`
}

// NewConfig creates a new Config.
// Precedence order:
// - explicit call to Set
// - flag
// - env
// - config
// - key/value store
// - default
func NewConfig(path string) (*Config, error) {
	conf := Config{}

	setDefault()

	if err := readConfigFile(path); err != nil {
		return nil, err
	}

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func setDefault() {
	viper.SetDefault("SERVICE", "Service")
	viper.SetDefault("GRPC_PORT", "8080")
	viper.SetDefault("HEALTH_PORT", "8181")
	viper.SetDefault("DEBUG", false)
	viper.SetDefault("DB_MAX_OPEN_CONNS", 10)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 2)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", 60)
}

func readConfigFile(path string) error {
	viper.SetConfigName(path)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read the config file: %w", err)
	}

	return nil
}

// Update updates the configuration from the file at the specified path.
// Deprecated. It is used only to satisfy the third task.
func (c *Config) Update(path string) error {
	conf := struct {
		Name     string `json:"name"`
		GrpcPort string `json:"grpc_port"`
	}{}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &conf)
}
