package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env if it exists. Check current and parent directory.
	_ = godotenv.Load(".env", "../.env")
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
