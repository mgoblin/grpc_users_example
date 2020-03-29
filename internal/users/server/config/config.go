package config

import (
	"os"
	"strconv"
)

type UsersServerConfig struct {
	TCPPort int
}

type Config struct {
	UsersServer UsersServerConfig
}

func New() *Config {
	return &Config{
		UsersServer: UsersServerConfig{
			TCPPort: getEnvAsInt("USERS_SERVER_PORT", 7777),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
