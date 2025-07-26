package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT       string
	DBHOST     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func getEnv(key, defaultKey string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultKey
	}
	return val
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	cfg := &Config{
		PORT:       getEnv("PORT", "8080"),
		DBHOST:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBName:     getEnv("DB_NAME", "notes_db"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecret"),
	}
	return cfg
}
