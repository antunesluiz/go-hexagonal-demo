package config

import (
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadDatabaseConfig() *DatabaseConfig {
	_ = godotenv.Load()

	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     port,
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "root"),
		DBName:   getEnv("DB_NAME", "hexagonal_user"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}
