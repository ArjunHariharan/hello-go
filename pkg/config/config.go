package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config exported
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig exported
type ServerConfig struct {
	Port int
}

// DatabaseConfig exported
type DatabaseConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
}

// New creates an instance of configuration
func New() (Config, error) {
	viper.AddConfigPath("/")
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	c := Config{}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return c, err
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return c, err
	}

	return c, nil
}
