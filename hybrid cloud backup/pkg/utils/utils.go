package utils

import (
	"log"
	"os"
	"strconv"
	"time"
)

// CheckError logs a fatal error if one is encountered
func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

// FileExists checks if a file exists at the given path
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ParseEnvInt parses an integer environment variable or returns a fallback value
func ParseEnvInt(key string, fallback int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

func ParseEnvDuration(key string, fallback time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return fallback
}

func ParseEnvBool(key string, fallback bool) bool {
	valueStr := os.Getenv(key)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return fallback
}

func Logger(message string) {
	log.Printf("%s: %s", time.Now().Format(time.RFC3339), message)
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
