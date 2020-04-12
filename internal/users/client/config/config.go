package config

import (
	"os"
	"strconv"
)

type WebServerConfig struct {
	TCPPort int
}

type UserServiceConfig struct {
	URL string
}

type Config struct {
	WebConfig   WebServerConfig
	UserService UserServiceConfig
}

func New() *Config {
	port := getEnvAsInt("USERS_WEB_PORT", 9090)
	userServiceURL := getEnv("USERS_SERVER_URL", "localhost:7777")

	return &Config{
		WebConfig: WebServerConfig{
			TCPPort: port,
		},
		UserService: UserServiceConfig{
			URL: userServiceURL,
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
