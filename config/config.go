package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig holds the server configuration
type ServerConfig struct {
	Port string
}

// DatabaseConfig holds the database configuration
type DatabaseConfig struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// LoadConfig loads the configuration from a .env file
func LoadConfig(env string) (*Config, error) {
	configPath := filepath.Join(basepath, env+".env")
	if err := godotenv.Load(configPath); err != nil {
		return nil, err
	}

	config := Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		// Placeholder: here you need to change the database default configs
		Database: DatabaseConfig{
			Type:     getEnv("DB_TYPE", "mysql"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "root"),
			Name:     getEnv("DB_NAME", "rest-template"),
		},
	}

	return &config, nil
}

// getEnv gets the value of the environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
