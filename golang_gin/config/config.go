package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadENV untuk membaca file .env
func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}

// GetENV untuk membaca nilai dari environment variable di .env
func GetENV(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return value
}
