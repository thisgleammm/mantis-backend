package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env if it exists. Ignore error if it doesn't (production might use real env vars).
	_ = godotenv.Load()
}

func GetString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func RequiredString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s environment variable is required", key)
	}
	return val
}