package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConfig *DatabaseConfig
	Port           string
}

func New() *Config {
	return &Config{
		DatabaseConfig: NewDatabaseConfig(),
		Port:           getEnv("PORT"),
	}
}

type DatabaseConfig struct {
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	DbHost     string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DbUser:     getEnv("DB_USER"),
		DbPassword: getEnv("DB_PASSWORD"),
		DbName:     getEnv("DB_NAME"),
		DbPort:     getEnv("DB_PORT"),
		DbHost:     getEnv("DB_HOST"),
	}
}

func getEnv(key string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(".env doesnt exists")
}
