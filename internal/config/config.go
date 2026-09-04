package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DatabaseUrl string
}

func load() (*config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}
	dbURL := os.Getenv("DATABAE_URL")

	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	return &config{DatabaseUrl: dbURL}, nil
}
