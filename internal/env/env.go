package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type EnvFileConfig struct {
	EnvFile      string
	EnvFiles     []string
	RequiredVars []string
}

func Load(cfg *EnvFileConfig) error {
	if cfg.EnvFile == "" {
		cfg.EnvFile = ".env"
	}

	filesToLoad := cfg.EnvFiles
	if len(filesToLoad) == 0 {
		filesToLoad = []string{cfg.EnvFile}
	}

	for _, file := range filesToLoad {
		if err := godotenv.Overload(file); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvAsInt(key string, fallback int) int {
	valueStr := GetEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

func GetEnvAsBool(key string, fallback bool) bool {
	valueStr := GetEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return fallback
}

func GetEnvAsDuration(key string, fallback time.Duration) time.Duration {
	valueStr := GetEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return fallback
}

func EnsureExists(keys ...string) error {
	var missing []string
	for _, key := range keys {
		if _, exists := os.LookupEnv(key); !exists {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}
	return nil
}
