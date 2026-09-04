package main

import (
	"context"
	"log"
	"subscription-tracker/internal/config"
	"subscription-tracker/internal/db"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	pool, err := db.Connect(context.Background(), cfg.DatabaseUrl)

	if err != nil {
		log.Fatalf("error connect to database: %v", err)
	}

	defer pool.Close()

	var version string
	if err := pool.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	log.Println("Connected to:", version)

}
