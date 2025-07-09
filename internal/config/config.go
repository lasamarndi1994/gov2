package config

import (
	"fmt"
	"log"
	"os"
	"strconv" // For parsing numbers/booleans from string env vars

	"github.com/joho/godotenv"
)

// Config holds all application configuration settings
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
	AppPort    string
	// Add any other configuration variables here
}

// LoadConfig loads environment variables from .env and populates the Config struct
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		AppPort:    os.Getenv("APP_PORT"),
	}

	// Basic validation for critical config (you can add more comprehensive checks)
	if cfg.DBUser == "" || cfg.DBHost == "" ||
		cfg.DBPort == "" || cfg.DBName == "" || cfg.JWTSecret == "" ||
		cfg.AppPort == "" {
		log.Fatal("One or more critical environment variables are missing. Please check your .env file.")
	}

	// Example of parsing a boolean or integer if you had one in .env
	// debugModeStr := os.Getenv("DEBUG_MODE")
	// cfg.DebugMode, _ = strconv.ParseBool(debugModeStr)

	return cfg
}

// GetDBPortAsInt returns the DBPort as an integer, handling errors
func (c *Config) GetDBPortAsInt() (int, error) {
	port, err := strconv.Atoi(c.DBPort)
	if err != nil {
		return 0, fmt.Errorf("invalid DB_PORT in config: %v", err)
	}
	return port, nil
}
