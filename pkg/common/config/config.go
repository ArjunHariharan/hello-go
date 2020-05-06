package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config exported
type Config struct {
	WebServer WebServerConfig
	Database  DatabaseConfig
}

// WebServerConfig exported
type WebServerConfig struct {
	Port int
}

// DatabaseConfig exported
type DatabaseConfig struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

// New creates an instance of configuration
func New() (*Config, error) {
	viper.AddConfigPath("/")
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil, err
	}

	c := Config{}
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
