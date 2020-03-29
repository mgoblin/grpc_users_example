package config

import (
	"os"
	"strconv"
)

type UsersServerConfig struct {
	TCPPort int
}

type IdGenConfig struct {
	URL string
}

type Config struct {
	UsersServer UsersServerConfig
	IdGen       IdGenConfig
}

func New() *Config {
	return &Config{
		UsersServer: UsersServerConfig{
			TCPPort: getEnvAsInt("USERS_SERVER_PORT", 7777),
		},
		IdGen: IdGenConfig{
			URL: getEnv("USERS_SERVER_IDGEN_URL", "localhost:8080"),
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
