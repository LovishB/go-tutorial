package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	AuthKey string
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

	return &Config{
		Port:    port,
		AuthKey: authKey,
	}
}
