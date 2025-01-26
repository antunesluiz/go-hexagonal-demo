package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

func LoadAppConfig() *AppConfig {
	_ = godotenv.Load()

	return &AppConfig{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
