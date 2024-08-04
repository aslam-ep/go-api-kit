package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	DBTimeout    int
	JWTSecret    string
	APIRateLimit int
}

var AppConfig *Config

// LoadConfig loads enviorment variables and store them in AppConfig
func LoadConfig() {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	AppConfig = &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnvAsInt("DB_PORT", 5432),
		DBUser:       getEnv("DB_USER", "root"),
		DBPassword:   getEnv("DB_PASSWORD", "password"),
		DBName:       getEnv("DB_NAME", "e-commerce"),
		DBTimeout:    getEnvAsInt("DB_TIMEOUT", 2),
		JWTSecret:    getEnv("JWT_SECRET", "someSecretKey"),
		APIRateLimit: getEnvAsInt("API_RATE_LIMIT", 100),
	}
}

// getEnv reads environment variable and return default value if not found
func getEnv(key, defalutValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defalutValue
}

// getEnvAsInt reads enviorement variable as integer and return default value if not found
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
