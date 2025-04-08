package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	AuthKey string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

/*
LoadConfig loads the configuration from environment variables
and returns a Config struct.
*/
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT environment variable not set")
	}

	authKey := os.Getenv("AUTH_KEY")
	if authKey == "" {
		log.Fatalf("AUTH_KEY environment variable not set")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatalf("DB_HOST environment variable not set")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatalf("DB_PORT environment variable not set")
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatalf("DB_USER environment variable not set")
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		log.Fatalf("DB_PASS environment variable not set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatalf("DB_NAME environment variable not set")
	}

	return &Config{
		Port:    port,
		AuthKey: authKey,
		DBHost:  dbHost,
		DBPort:  dbPort,
		DBUser:  dbUser,
		DBPass:  dbPass,
		DBName:  dbName,
	}
}
